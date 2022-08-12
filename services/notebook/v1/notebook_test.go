package notebookv1

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/bufbuild/connect-go"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"

	notebookv1 "github.com/jasonblanchard/di-notebook-connect/gen/proto/go/notebookapis/notebook/v1"
	notebookstore "github.com/jasonblanchard/di-notebook-connect/gen/sqlc/notebook"
)

type MockStore struct{}

func (s *MockStore) GetEntryByIdAndAuthor(ctx context.Context, params notebookstore.GetEntryByIdAndAuthorParams) (notebookstore.GetEntryByIdAndAuthorRow, error) {
	if params.ID == 86 {
		result := &notebookstore.GetEntryByIdAndAuthorRow{}

		return *result, sql.ErrNoRows
	}

	result := &notebookstore.GetEntryByIdAndAuthorRow{
		ID: 123,
		Text: sql.NullString{
			String: "Entry from mock store",
		},
		CreatorID: "abc123",
	}

	return *result, nil

}

func (s *MockStore) ListEntriesByAuthor(ctx context.Context, params notebookstore.ListEntriesByAuthorParams) ([]notebookstore.Entry, error) {
	result := []notebookstore.Entry{}

	return result, nil
}

func (s *MockStore) CreateEntry(ctx context.Context, params notebookstore.CreateEntryParams) (int32, error) {
	return 1, nil
}

func (s *MockStore) UpdateEntryText(ctx context.Context, params notebookstore.UpdateEntryTextParams) (notebookstore.UpdateEntryTextRow, error) {
	result := &notebookstore.UpdateEntryTextRow{
		ID: 123,
		Text: sql.NullString{
			String: params.Text.String,
		},
		CreatorID: "abc123",
		CreatedAt: time.Now(),
		UpdatedAt: sql.NullTime{
			Time: time.Now(),
		},
	}

	return *result, nil
}

func (s *MockStore) DeleteEntryByIdAndAuthor(ctx context.Context, params notebookstore.DeleteEntryByIdAndAuthorParams) (notebookstore.Entry, error) {
	result := &notebookstore.Entry{
		ID: 123,
		Text: sql.NullString{
			String: "Entry from mock store",
		},
		CreatorID: "abc123",
		CreatedAt: time.Now(),
		UpdatedAt: sql.NullTime{
			Time: time.Now(),
		},
		DeleteTime: sql.NullTime{
			Time: time.Now(),
		},
	}

	return *result, nil
}

func (s *MockStore) UnDeleteEntryByIdAndAuthor(ctx context.Context, params notebookstore.UnDeleteEntryByIdAndAuthorParams) (notebookstore.Entry, error) {
	result := &notebookstore.Entry{
		ID: 123,
		Text: sql.NullString{
			String: "Entry from mock store",
		},
		CreatorID: "abc123",
		CreatedAt: time.Now(),
		UpdatedAt: sql.NullTime{
			Time: time.Now(),
		},
		DeleteTime: sql.NullTime{
			Valid: false,
		},
	}

	return *result, nil
}

func TestReadAuthorEntry(t *testing.T) {
	notebookStore := &MockStore{}
	notebookService := &Service{
		Store: notebookStore,
	}
	c := context.TODO()

	req := &connect.Request[notebookv1.ReadAuthorEntryRequest]{
		Msg: &notebookv1.ReadAuthorEntryRequest{
			Id: 123,
		},
	}

	req.Header().Set("x-principal-id", "abc123")

	result, err := notebookService.ReadAuthorEntry(c, req)

	assert.Nil(t, err)
	assert.Equal(t, "Entry from mock store", result.Msg.Entry.Text)
	assert.Positive(t, result.Msg.Entry.Id)
	assert.Equal(t, result.Msg.Entry.CreatorId, "abc123")
	assert.True(t, result.Msg.Entry.CreatedAt.IsValid())
	assert.False(t, result.Msg.Entry.DeleteTime.IsValid())
}

func TestReadAuthorEntryNotFound(t *testing.T) {
	notebookStore := &MockStore{}
	logger, _ := zap.NewDevelopment()
	sugar := logger.Sugar()
	notebookService := &Service{
		Store:  notebookStore,
		Logger: sugar,
	}
	c := context.TODO()

	req := &connect.Request[notebookv1.ReadAuthorEntryRequest]{
		Msg: &notebookv1.ReadAuthorEntryRequest{
			Id: 86,
		},
	}

	req.Header().Set("x-principal-id", "abc123")

	_, err := notebookService.ReadAuthorEntry(c, req)

	assert.Equal(t, "not_found: not_found", err.Error())
	assert.Equal(t, connect.CodeNotFound, connect.CodeOf(err))
}

func TestReadAuthorEntryAunauthorized(t *testing.T) {
	notebookStore := &MockStore{}
	notebookService := &Service{
		Store: notebookStore,
	}
	c := context.TODO()

	req := &connect.Request[notebookv1.ReadAuthorEntryRequest]{
		Msg: &notebookv1.ReadAuthorEntryRequest{
			Id: 123,
		},
	}

	_, err := notebookService.ReadAuthorEntry(c, req)

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

	req.Header().Set("x-principal-id", "abc123")

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

	req.Header().Set("x-principal-id", "abc123")

	result, err := notebookService.WriteToEntry(c, req)

	assert.Nil(t, err)
	assert.Equal(t, "New entry", result.Msg.Entry.Text)
}

func TestDeleteEntry(t *testing.T) {
	notebookStore := &MockStore{}
	notebookService := &Service{
		Store: notebookStore,
	}
	c := context.TODO()

	req := &connect.Request[notebookv1.DeleteEntryRequest]{
		Msg: &notebookv1.DeleteEntryRequest{
			Id: 1,
		},
	}

	req.Header().Set("x-principal-id", "abc123")

	result, err := notebookService.DeleteEntry(c, req)

	assert.Nil(t, err)
	assert.True(t, result.Msg.Entry.DeleteTime.IsValid())
}

func TestUndeleteDeleteEntry(t *testing.T) {
	notebookStore := &MockStore{}
	notebookService := &Service{
		Store: notebookStore,
	}
	c := context.TODO()

	req := &connect.Request[notebookv1.UnDeleteEntryRequest]{
		Msg: &notebookv1.UnDeleteEntryRequest{
			Id: 1,
		},
	}

	req.Header().Set("x-principal-id", "abc123")

	result, err := notebookService.UndeleteEntry(c, req)

	assert.Nil(t, err)
	assert.False(t, result.Msg.Entry.DeleteTime.IsValid())
}
