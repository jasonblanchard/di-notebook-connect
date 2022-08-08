package main

import (
	"context"
	"database/sql"

	notebookstore "github.com/jasonblanchard/di-notebook-connect/gen/sqlc/notebook"
)

type MockNotebookStore struct{}

func (s *MockNotebookStore) GetEntry(ctx context.Context, id int32) (notebookstore.GetEntryRow, error) {
	result := &notebookstore.GetEntryRow{
		ID: 123,
		Text: sql.NullString{
			String: "Entry from mock store",
		},
		CreatorID: "abc123",
	}

	return *result, nil
}
