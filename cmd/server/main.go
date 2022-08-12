package main

import (
	"fmt"
	"net/http"

	"github.com/jasonblanchard/di-notebook-connect/gen/proto/go/notebookapis/notebook/v1/notebookv1connect"
	"github.com/jasonblanchard/di-notebook-connect/gen/proto/go/notebookapis/ping/v1/pingv1connect"

	"database/sql"

	notebookstore "github.com/jasonblanchard/di-notebook-connect/gen/sqlc/notebook"
	notebookv1 "github.com/jasonblanchard/di-notebook-connect/services/notebook/v1"
	pingv1 "github.com/jasonblanchard/di-notebook-connect/services/ping/v1"
	_ "github.com/lib/pq"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	// TODO: Fix this string
	db, err := sql.Open("postgres", "user=pqgotest dbname=pqgotest sslmode=verify-full")
	if err != nil {
		panic(err)
	}

	store := notebookstore.New(db)

	mux := http.NewServeMux()
	pingService := &pingv1.Service{}
	pingpath, pinghandler := pingv1connect.NewPingServiceHandler(pingService)
	mux.Handle(pingpath, pinghandler)
	notebookService := &notebookv1.Service{
		Store: store,
	}
	notebookpath, notebookhandler := notebookv1connect.NewNotebookServiceHandler(notebookService)
	mux.Handle(notebookpath, notebookhandler)
	fmt.Println("Starting server on port 8080")
	http.ListenAndServe(
		"localhost:8080",
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
	)
}
