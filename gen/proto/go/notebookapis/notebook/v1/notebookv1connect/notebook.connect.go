// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: notebookapis/notebook/v1/notebook.proto

package notebookv1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/jasonblanchard/di-notebook-connect/gen/proto/go/notebookapis/notebook/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// NotebookServiceName is the fully-qualified name of the NotebookService service.
	NotebookServiceName = "notebook.v1.NotebookService"
)

// NotebookServiceClient is a client for the notebook.v1.NotebookService service.
type NotebookServiceClient interface {
	Ping(context.Context, *connect_go.Request[v1.PingRequest]) (*connect_go.Response[v1.PingResponse], error)
	ReadAuthorEntry(context.Context, *connect_go.Request[v1.ReadAuthorEntryRequest]) (*connect_go.Response[v1.ReadAuthorEntryResponse], error)
	BeginNewEntry(context.Context, *connect_go.Request[v1.BeginNewEntryRequest]) (*connect_go.Response[v1.BeginNewEntryResponse], error)
	WriteToEntry(context.Context, *connect_go.Request[v1.WriteToEntryRequest]) (*connect_go.Response[v1.WriteToEntryResponse], error)
	ListEntries(context.Context, *connect_go.Request[v1.ListEntriesRequest]) (*connect_go.Response[v1.ListEntriesResponse], error)
	DeleteEntry(context.Context, *connect_go.Request[v1.DeleteEntryRequest]) (*connect_go.Response[v1.DeleteEntryResponse], error)
	UndeleteEntry(context.Context, *connect_go.Request[v1.UnDeleteEntryRequest]) (*connect_go.Response[v1.UnDeleteEntryResponse], error)
}

// NewNotebookServiceClient constructs a client for the notebook.v1.NotebookService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewNotebookServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) NotebookServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &notebookServiceClient{
		ping: connect_go.NewClient[v1.PingRequest, v1.PingResponse](
			httpClient,
			baseURL+"/notebook.v1.NotebookService/Ping",
			opts...,
		),
		readAuthorEntry: connect_go.NewClient[v1.ReadAuthorEntryRequest, v1.ReadAuthorEntryResponse](
			httpClient,
			baseURL+"/notebook.v1.NotebookService/ReadAuthorEntry",
			opts...,
		),
		beginNewEntry: connect_go.NewClient[v1.BeginNewEntryRequest, v1.BeginNewEntryResponse](
			httpClient,
			baseURL+"/notebook.v1.NotebookService/BeginNewEntry",
			opts...,
		),
		writeToEntry: connect_go.NewClient[v1.WriteToEntryRequest, v1.WriteToEntryResponse](
			httpClient,
			baseURL+"/notebook.v1.NotebookService/WriteToEntry",
			opts...,
		),
		listEntries: connect_go.NewClient[v1.ListEntriesRequest, v1.ListEntriesResponse](
			httpClient,
			baseURL+"/notebook.v1.NotebookService/ListEntries",
			opts...,
		),
		deleteEntry: connect_go.NewClient[v1.DeleteEntryRequest, v1.DeleteEntryResponse](
			httpClient,
			baseURL+"/notebook.v1.NotebookService/DeleteEntry",
			opts...,
		),
		undeleteEntry: connect_go.NewClient[v1.UnDeleteEntryRequest, v1.UnDeleteEntryResponse](
			httpClient,
			baseURL+"/notebook.v1.NotebookService/UndeleteEntry",
			opts...,
		),
	}
}

// notebookServiceClient implements NotebookServiceClient.
type notebookServiceClient struct {
	ping            *connect_go.Client[v1.PingRequest, v1.PingResponse]
	readAuthorEntry *connect_go.Client[v1.ReadAuthorEntryRequest, v1.ReadAuthorEntryResponse]
	beginNewEntry   *connect_go.Client[v1.BeginNewEntryRequest, v1.BeginNewEntryResponse]
	writeToEntry    *connect_go.Client[v1.WriteToEntryRequest, v1.WriteToEntryResponse]
	listEntries     *connect_go.Client[v1.ListEntriesRequest, v1.ListEntriesResponse]
	deleteEntry     *connect_go.Client[v1.DeleteEntryRequest, v1.DeleteEntryResponse]
	undeleteEntry   *connect_go.Client[v1.UnDeleteEntryRequest, v1.UnDeleteEntryResponse]
}

// Ping calls notebook.v1.NotebookService.Ping.
func (c *notebookServiceClient) Ping(ctx context.Context, req *connect_go.Request[v1.PingRequest]) (*connect_go.Response[v1.PingResponse], error) {
	return c.ping.CallUnary(ctx, req)
}

// ReadAuthorEntry calls notebook.v1.NotebookService.ReadAuthorEntry.
func (c *notebookServiceClient) ReadAuthorEntry(ctx context.Context, req *connect_go.Request[v1.ReadAuthorEntryRequest]) (*connect_go.Response[v1.ReadAuthorEntryResponse], error) {
	return c.readAuthorEntry.CallUnary(ctx, req)
}

// BeginNewEntry calls notebook.v1.NotebookService.BeginNewEntry.
func (c *notebookServiceClient) BeginNewEntry(ctx context.Context, req *connect_go.Request[v1.BeginNewEntryRequest]) (*connect_go.Response[v1.BeginNewEntryResponse], error) {
	return c.beginNewEntry.CallUnary(ctx, req)
}

// WriteToEntry calls notebook.v1.NotebookService.WriteToEntry.
func (c *notebookServiceClient) WriteToEntry(ctx context.Context, req *connect_go.Request[v1.WriteToEntryRequest]) (*connect_go.Response[v1.WriteToEntryResponse], error) {
	return c.writeToEntry.CallUnary(ctx, req)
}

// ListEntries calls notebook.v1.NotebookService.ListEntries.
func (c *notebookServiceClient) ListEntries(ctx context.Context, req *connect_go.Request[v1.ListEntriesRequest]) (*connect_go.Response[v1.ListEntriesResponse], error) {
	return c.listEntries.CallUnary(ctx, req)
}

// DeleteEntry calls notebook.v1.NotebookService.DeleteEntry.
func (c *notebookServiceClient) DeleteEntry(ctx context.Context, req *connect_go.Request[v1.DeleteEntryRequest]) (*connect_go.Response[v1.DeleteEntryResponse], error) {
	return c.deleteEntry.CallUnary(ctx, req)
}

// UndeleteEntry calls notebook.v1.NotebookService.UndeleteEntry.
func (c *notebookServiceClient) UndeleteEntry(ctx context.Context, req *connect_go.Request[v1.UnDeleteEntryRequest]) (*connect_go.Response[v1.UnDeleteEntryResponse], error) {
	return c.undeleteEntry.CallUnary(ctx, req)
}

// NotebookServiceHandler is an implementation of the notebook.v1.NotebookService service.
type NotebookServiceHandler interface {
	Ping(context.Context, *connect_go.Request[v1.PingRequest]) (*connect_go.Response[v1.PingResponse], error)
	ReadAuthorEntry(context.Context, *connect_go.Request[v1.ReadAuthorEntryRequest]) (*connect_go.Response[v1.ReadAuthorEntryResponse], error)
	BeginNewEntry(context.Context, *connect_go.Request[v1.BeginNewEntryRequest]) (*connect_go.Response[v1.BeginNewEntryResponse], error)
	WriteToEntry(context.Context, *connect_go.Request[v1.WriteToEntryRequest]) (*connect_go.Response[v1.WriteToEntryResponse], error)
	ListEntries(context.Context, *connect_go.Request[v1.ListEntriesRequest]) (*connect_go.Response[v1.ListEntriesResponse], error)
	DeleteEntry(context.Context, *connect_go.Request[v1.DeleteEntryRequest]) (*connect_go.Response[v1.DeleteEntryResponse], error)
	UndeleteEntry(context.Context, *connect_go.Request[v1.UnDeleteEntryRequest]) (*connect_go.Response[v1.UnDeleteEntryResponse], error)
}

// NewNotebookServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewNotebookServiceHandler(svc NotebookServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/notebook.v1.NotebookService/Ping", connect_go.NewUnaryHandler(
		"/notebook.v1.NotebookService/Ping",
		svc.Ping,
		opts...,
	))
	mux.Handle("/notebook.v1.NotebookService/ReadAuthorEntry", connect_go.NewUnaryHandler(
		"/notebook.v1.NotebookService/ReadAuthorEntry",
		svc.ReadAuthorEntry,
		opts...,
	))
	mux.Handle("/notebook.v1.NotebookService/BeginNewEntry", connect_go.NewUnaryHandler(
		"/notebook.v1.NotebookService/BeginNewEntry",
		svc.BeginNewEntry,
		opts...,
	))
	mux.Handle("/notebook.v1.NotebookService/WriteToEntry", connect_go.NewUnaryHandler(
		"/notebook.v1.NotebookService/WriteToEntry",
		svc.WriteToEntry,
		opts...,
	))
	mux.Handle("/notebook.v1.NotebookService/ListEntries", connect_go.NewUnaryHandler(
		"/notebook.v1.NotebookService/ListEntries",
		svc.ListEntries,
		opts...,
	))
	mux.Handle("/notebook.v1.NotebookService/DeleteEntry", connect_go.NewUnaryHandler(
		"/notebook.v1.NotebookService/DeleteEntry",
		svc.DeleteEntry,
		opts...,
	))
	mux.Handle("/notebook.v1.NotebookService/UndeleteEntry", connect_go.NewUnaryHandler(
		"/notebook.v1.NotebookService/UndeleteEntry",
		svc.UndeleteEntry,
		opts...,
	))
	return "/notebook.v1.NotebookService/", mux
}

// UnimplementedNotebookServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedNotebookServiceHandler struct{}

func (UnimplementedNotebookServiceHandler) Ping(context.Context, *connect_go.Request[v1.PingRequest]) (*connect_go.Response[v1.PingResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("notebook.v1.NotebookService.Ping is not implemented"))
}

func (UnimplementedNotebookServiceHandler) ReadAuthorEntry(context.Context, *connect_go.Request[v1.ReadAuthorEntryRequest]) (*connect_go.Response[v1.ReadAuthorEntryResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("notebook.v1.NotebookService.ReadAuthorEntry is not implemented"))
}

func (UnimplementedNotebookServiceHandler) BeginNewEntry(context.Context, *connect_go.Request[v1.BeginNewEntryRequest]) (*connect_go.Response[v1.BeginNewEntryResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("notebook.v1.NotebookService.BeginNewEntry is not implemented"))
}

func (UnimplementedNotebookServiceHandler) WriteToEntry(context.Context, *connect_go.Request[v1.WriteToEntryRequest]) (*connect_go.Response[v1.WriteToEntryResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("notebook.v1.NotebookService.WriteToEntry is not implemented"))
}

func (UnimplementedNotebookServiceHandler) ListEntries(context.Context, *connect_go.Request[v1.ListEntriesRequest]) (*connect_go.Response[v1.ListEntriesResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("notebook.v1.NotebookService.ListEntries is not implemented"))
}

func (UnimplementedNotebookServiceHandler) DeleteEntry(context.Context, *connect_go.Request[v1.DeleteEntryRequest]) (*connect_go.Response[v1.DeleteEntryResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("notebook.v1.NotebookService.DeleteEntry is not implemented"))
}

func (UnimplementedNotebookServiceHandler) UndeleteEntry(context.Context, *connect_go.Request[v1.UnDeleteEntryRequest]) (*connect_go.Response[v1.UnDeleteEntryResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("notebook.v1.NotebookService.UndeleteEntry is not implemented"))
}
