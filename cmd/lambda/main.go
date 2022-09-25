package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
	"github.com/bufbuild/connect-go"
	"github.com/jasonblanchard/di-notebook-connect/gen/proto/go/notebookapis/notebook/v1/notebookv1connect"
	"github.com/jasonblanchard/di-notebook-connect/ingress"
	"go.uber.org/zap"

	"database/sql"

	notebookstore "github.com/jasonblanchard/di-notebook-connect/gen/sqlc/notebook"
	notebookv1 "github.com/jasonblanchard/di-notebook-connect/services/notebook/v1"
	_ "github.com/lib/pq"
)

func main() {
	pghost := os.Getenv("PGHOST")
	pgport := os.Getenv("PGPORT")
	pgdatabase := os.Getenv("PGDATABASE")
	pgpassword := os.Getenv("PGPASSWORD")

	db, err := sql.Open("postgres", fmt.Sprintf("user=postgres host=%s port=%s dbname=%s password=%s sslmode=disable", pghost, pgport, pgdatabase, pgpassword))
	if err != nil {
		panic(err)
	}

	store := notebookstore.New(db)

	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	sugaredLogger := logger.Sugar()

	mux := http.NewServeMux()

	notebookService := &notebookv1.Service{
		Store:  store,
		Logger: sugaredLogger,
	}

	interceptors := connect.WithInterceptors(ingress.NewAuthInterceptor())

	notebookpath, notebookhandler := notebookv1connect.NewNotebookServiceHandler(notebookService, interceptors)
	mux.Handle(notebookpath, notebookhandler)
	wrappedMux := ingress.NewWithLogger(mux, sugaredLogger)

	sugaredLogger.Info("Starting handler")
	lambda.Start(httpadapter.NewV2(wrappedMux).ProxyWithContext)
}
