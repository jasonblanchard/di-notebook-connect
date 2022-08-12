package notebookv1

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/bufbuild/connect-go"
	notebookv1 "github.com/jasonblanchard/di-notebook-connect/gen/proto/go/notebookapis/notebook/v1"
	notebookstore "github.com/jasonblanchard/di-notebook-connect/gen/sqlc/notebook"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Store interface {
	GetEntryByIdAndAuthor(ctx context.Context, params notebookstore.GetEntryByIdAndAuthorParams) (notebookstore.GetEntryByIdAndAuthorRow, error)
	CreateEntry(ctx context.Context, params notebookstore.CreateEntryParams) (int32, error)
	UpdateEntryText(ctx context.Context, params notebookstore.UpdateEntryTextParams) (notebookstore.UpdateEntryTextRow, error)
}

type Service struct {
	Store Store
}

func (s *Service) GetAuthorEntry(ctx context.Context, req *connect.Request[notebookv1.GetAuthorEntryRequest]) (*connect.Response[notebookv1.GetAuthorEntryResponse], error) {
	principalId := req.Header().Get("principalId")

	if principalId == "" {
		return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("no principalId"))
	}

	entryRecord, err := s.Store.GetEntryByIdAndAuthor(ctx, notebookstore.GetEntryByIdAndAuthorParams{
		ID:        req.Msg.Id,
		CreatorID: principalId,
	})

	if err != nil {
		return nil, connect.NewError(connect.CodeUnknown, fmt.Errorf("error getting from store: %w", err))
	}

	if entryRecord.ID == 0 {
		return nil, connect.NewError(connect.CodeNotFound, fmt.Errorf("no entry found for id %v principalId %v", req.Msg.Id, principalId))
	}

	res := connect.NewResponse(&notebookv1.GetAuthorEntryResponse{
		Entry: &notebookv1.Entry{
			Id:        entryRecord.ID,
			Text:      entryRecord.Text.String,
			CreatorId: entryRecord.CreatorID,
			CreatedAt: &timestamppb.Timestamp{
				Seconds: entryRecord.CreatedAt.Unix(),
			},
			UpdatedAt: &timestamppb.Timestamp{
				Seconds: entryRecord.UpdatedAt.Time.Unix(),
			},
		},
	})
	return res, nil
}

func (s *Service) BeginNewEntry(ctx context.Context, req *connect.Request[notebookv1.BeginNewEntryRequest]) (*connect.Response[notebookv1.BeginNewEntryResponse], error) {
	principalId := req.Header().Get("principalId")

	if principalId == "" {
		return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("no principalId"))
	}

	id, err := s.Store.CreateEntry(ctx, notebookstore.CreateEntryParams{
		Text: sql.NullString{
			String: req.Msg.Text,
		},
		CreatorID: principalId,
		CreatedAt: time.Now(),
	})

	if err != nil {
		return nil, connect.NewError(connect.CodeUnknown, fmt.Errorf("error writing to store: %w", err))
	}

	entryRecord, err := s.Store.GetEntryByIdAndAuthor(ctx, notebookstore.GetEntryByIdAndAuthorParams{
		ID:        id,
		CreatorID: principalId,
	})

	if err != nil {
		return nil, connect.NewError(connect.CodeUnknown, fmt.Errorf("error getting from store: %w", err))
	}

	res := connect.NewResponse(&notebookv1.BeginNewEntryResponse{
		Entry: &notebookv1.Entry{
			Id:        entryRecord.ID,
			Text:      entryRecord.Text.String,
			CreatorId: entryRecord.CreatorID,
			CreatedAt: &timestamppb.Timestamp{
				Seconds: entryRecord.CreatedAt.Unix(),
			},
			UpdatedAt: &timestamppb.Timestamp{
				Seconds: entryRecord.UpdatedAt.Time.Unix(),
			},
		},
	})
	return res, nil
}

func (s *Service) WriteToEntry(ctx context.Context, req *connect.Request[notebookv1.WriteToEntryRequest]) (*connect.Response[notebookv1.WriteToEntryResponse], error) {
	principalId := req.Header().Get("principalId")

	if principalId == "" {
		return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("no principalId"))
	}

	updatedEntryRecord, err := s.Store.UpdateEntryText(ctx, notebookstore.UpdateEntryTextParams{
		ID: req.Msg.Id,
		Text: sql.NullString{
			String: req.Msg.Text,
		},
		UpdatedAt: sql.NullTime{
			Time: time.Now(),
		},
	})

	if err != nil {
		return nil, connect.NewError(connect.CodeUnknown, fmt.Errorf("error updating in store: %w", err))
	}

	res := connect.NewResponse(&notebookv1.WriteToEntryResponse{
		Entry: &notebookv1.Entry{
			Id:        updatedEntryRecord.ID,
			Text:      updatedEntryRecord.Text.String,
			CreatorId: updatedEntryRecord.CreatorID,
			CreatedAt: &timestamppb.Timestamp{
				Seconds: updatedEntryRecord.CreatedAt.Unix(),
			},
			UpdatedAt: &timestamppb.Timestamp{
				Seconds: updatedEntryRecord.UpdatedAt.Time.Unix(),
			},
		},
	})

	return res, nil
}
