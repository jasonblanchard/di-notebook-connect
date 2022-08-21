package main

import (
	"fmt"
	"net/http"

	"github.com/bufbuild/connect-go"
	"github.com/jasonblanchard/di-notebook-connect/gen/proto/go/notebookapis/notebook/v1/notebookv1connect"
	"github.com/jasonblanchard/di-notebook-connect/gen/proto/go/notebookapis/ping/v1/pingv1connect"
	"github.com/jasonblanchard/di-notebook-connect/ingress"
	"go.uber.org/zap"

	"database/sql"

	notebookstore "github.com/jasonblanchard/di-notebook-connect/gen/sqlc/notebook"
	notebookv1 "github.com/jasonblanchard/di-notebook-connect/services/notebook/v1"
	pingv1 "github.com/jasonblanchard/di-notebook-connect/services/ping/v1"
	_ "github.com/lib/pq"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	db, err := sql.Open("postgres", "user=postgres host=localhost port=52926 dbname=postgres password=mysecretpassword sslmode=disable")
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

	pingService := &pingv1.Service{}
	pingpath, pinghandler := pingv1connect.NewPingServiceHandler(pingService)
	mux.Handle(pingpath, pinghandler)

	notebookService := &notebookv1.Service{
		Store:  store,
		Logger: sugaredLogger,
	}

	// TODO: Interceptor for translating `Authorization: Bearer aaa.bbb.ccc` JWT to x-principal-id https://connect.build/docs/go/interceptors
	interceptors := connect.WithInterceptors(ingress.NewAuthInterceptor())

	notebookpath, notebookhandler := notebookv1connect.NewNotebookServiceHandler(notebookService, interceptors)
	mux.Handle(notebookpath, notebookhandler)

	fmt.Println("Starting server on port 8080")

	wrappedMux := NewWithLogger(mux, sugaredLogger)

	http.ListenAndServe(
		"localhost:8080",
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(wrappedMux, &http2.Server{}),
	)
}
