package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/bufbuild/connect-go"
	"github.com/jasonblanchard/di-notebook-connect/gen/proto/go/notebookapis/notebook/v1/notebookv1connect"
	"github.com/jasonblanchard/di-notebook-connect/ingress"
	"go.uber.org/zap"

	"database/sql"

	notebookstore "github.com/jasonblanchard/di-notebook-connect/gen/sqlc/notebook"
	notebookv1 "github.com/jasonblanchard/di-notebook-connect/services/notebook/v1"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

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

	rootMux := http.NewServeMux()

	notebookService := &notebookv1.Service{
		Store:  store,
		Logger: sugaredLogger,
	}
	interceptors := connect.WithInterceptors(ingress.NewAuthInterceptor())
	notebookpath, notebookhandler := notebookv1connect.NewNotebookServiceHandler(notebookService, interceptors)
	// notebookMux := http.NewServeMux()
	rootMux.Handle(notebookpath, notebookhandler)
	// rootMux.Handle("/connect/", http.StripPrefix("/connect/", notebookMux))

	port := os.Getenv("PORT")

	sugaredLogger.Infow("Starting server", "port", port)

	wrappedMux := ingress.NewWithLogger(rootMux, sugaredLogger)

	http.ListenAndServe(
		fmt.Sprintf("0.0.0.0:%s", port),
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(wrappedMux, &http2.Server{}),
	)
}
