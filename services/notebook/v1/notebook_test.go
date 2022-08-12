package notebookv1

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/bufbuild/connect-go"
	"github.com/stretchr/testify/assert"

	notebookv1 "github.com/jasonblanchard/di-notebook-connect/gen/proto/go/notebookapis/notebook/v1"
	notebookstore "github.com/jasonblanchard/di-notebook-connect/gen/sqlc/notebook"
)

type MockStore struct{}

func (s *MockStore) GetEntryByIdAndAuthor(ctx context.Context, params notebookstore.GetEntryByIdAndAuthorParams) (notebookstore.GetEntryByIdAndAuthorRow, error) {
	result := &notebookstore.GetEntryByIdAndAuthorRow{
		ID: 123,
		Text: sql.NullString{
			String: "Entry from mock store",
		},
		CreatorID: "abc123",
	}

	return *result, nil
}

func (s *MockStore) CreateEntry(ctx context.Context, params notebookstore.CreateEntryParams) (int32, error) {
	return 1, nil
}

func (s *MockStore) UpdateEntryText(ctx context.Context, params notebookstore.UpdateEntryTextParams) (notebookstore.UpdateEntryTextRow, error) {
	result := &notebookstore.UpdateEntryTextRow{
		ID: 123,
		Text: sql.NullString{
			String: "Entry from mock store",
		},
		CreatorID: "abc123",
		CreatedAt: time.Now(),
		UpdatedAt: sql.NullTime{
			Time: time.Now(),
		},
	}

	return *result, nil
}

type NullMockStore struct{}

func (s *NullMockStore) GetEntryByIdAndAuthor(ctx context.Context, params notebookstore.GetEntryByIdAndAuthorParams) (notebookstore.GetEntryByIdAndAuthorRow, error) {
	result := &notebookstore.GetEntryByIdAndAuthorRow{}

	return *result, nil
}

func (s *NullMockStore) CreateEntry(ctx context.Context, params notebookstore.CreateEntryParams) (int32, error) {
	return 1, nil
}

func (s *NullMockStore) UpdateEntryText(ctx context.Context, params notebookstore.UpdateEntryTextParams) (notebookstore.UpdateEntryTextRow, error) {
	result := &notebookstore.UpdateEntryTextRow{}

	return *result, nil
}

func TestGetAuthorEntry(t *testing.T) {
	notebookStore := &MockStore{}
	notebookService := &Service{
		Store: notebookStore,
	}
	c := context.TODO()

	req := &connect.Request[notebookv1.GetAuthorEntryRequest]{
		Msg: &notebookv1.GetAuthorEntryRequest{
			Id: 123,
		},
	}

	req.Header().Set("principalId", "abc123")

	result, err := notebookService.GetAuthorEntry(c, req)

	assert.Nil(t, err)
	assert.Equal(t, "Entry from mock store", result.Msg.Entry.Text)
}

func TestGetAuthorEntryNotFound(t *testing.T) {
	notebookStore := &NullMockStore{}
	notebookService := &Service{
		Store: notebookStore,
	}
	c := context.TODO()

	req := &connect.Request[notebookv1.GetAuthorEntryRequest]{
		Msg: &notebookv1.GetAuthorEntryRequest{
			Id: 123,
		},
	}

	req.Header().Set("principalId", "abc123")

	_, err := notebookService.GetAuthorEntry(c, req)

	assert.Equal(t, "not_found: no entry found for id 123 principalId abc123", err.Error())
	assert.Equal(t, connect.CodeNotFound, connect.CodeOf(err))
}

func TestGetEntryUnauthorized(t *testing.T) {
	notebookStore := &NullMockStore{}
	notebookService := &Service{
		Store: notebookStore,
	}
	c := context.TODO()

	req := &connect.Request[notebookv1.GetAuthorEntryRequest]{
		Msg: &notebookv1.GetAuthorEntryRequest{
			Id: 123,
		},
	}

	_, err := notebookService.GetAuthorEntry(c, req)

	assert.Equal(t, "unauthenticated: no principalId", err.Error())
	assert.Equal(t, connect.CodeUnauthenticated, connect.CodeOf(err))
}

func TestBeginNewEntry(t *testing.T) {
	notebookStore := &MockStore{}
	notebookService := &Service{
		Store: notebookStore,
	}
	c := context.TODO()

	req := &connect.Request[notebookv1.BeginNewEntryRequest]{
		Msg: &notebookv1.BeginNewEntryRequest{
			Text: "New entry",
		},
	}

	req.Header().Set("principalId", "abc123")

	result, err := notebookService.BeginNewEntry(c, req)

	assert.Nil(t, err)
	assert.Equal(t, "Entry from mock store", result.Msg.Entry.Text)
}

func TestWriteToEntry(t *testing.T) {
	notebookStore := &MockStore{}
	notebookService := &Service{
		Store: notebookStore,
	}
	c := context.TODO()

	req := &connect.Request[notebookv1.WriteToEntryRequest]{
		Msg: &notebookv1.WriteToEntryRequest{
			Text: "New entry",
		},
	}

	req.Header().Set("principalId", "abc123")

	result, err := notebookService.WriteToEntry(c, req)

	assert.Nil(t, err)
	assert.Equal(t, "Entry from mock store", result.Msg.Entry.Text)
}
