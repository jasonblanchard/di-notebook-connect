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
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Store interface {
	GetEntryByIdAndAuthor(ctx context.Context, params notebookstore.GetEntryByIdAndAuthorParams) (notebookstore.GetEntryByIdAndAuthorRow, error)
	CreateEntry(ctx context.Context, params notebookstore.CreateEntryParams) (int32, error)
	UpdateEntryText(ctx context.Context, params notebookstore.UpdateEntryTextParams) (notebookstore.UpdateEntryTextRow, error)
	DeleteEntryByIdAndAuthor(ctx context.Context, params notebookstore.DeleteEntryByIdAndAuthorParams) (notebookstore.Entry, error)
	UnDeleteEntryByIdAndAuthor(ctx context.Context, params notebookstore.UnDeleteEntryByIdAndAuthorParams) (notebookstore.Entry, error)
	ListEntriesByAuthor(ctx context.Context, params notebookstore.ListEntriesByAuthorParams) ([]notebookstore.Entry, error)
}

type Service struct {
	Store  Store
	Logger *zap.SugaredLogger
}

func (s *Service) ReadAuthorEntry(ctx context.Context, req *connect.Request[notebookv1.ReadAuthorEntryRequest]) (*connect.Response[notebookv1.ReadAuthorEntryResponse], error) {
	principalId := req.Header().Get("x-principal-id")

	if principalId == "" {
		return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("no principalId"))
	}

	entryRecord, err := s.Store.GetEntryByIdAndAuthor(ctx, notebookstore.GetEntryByIdAndAuthorParams{
		ID:        req.Msg.Id,
		CreatorID: principalId,
	})

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, connect.NewError(connect.CodeNotFound, fmt.Errorf(connect.CodeNotFound.String()))
		}
		s.Logger.Errorf("error getting from store: %s", err)
		return nil, connect.NewError(connect.CodeUnknown, fmt.Errorf(connect.CodeUnknown.String()))
	}

	res := connect.NewResponse(&notebookv1.ReadAuthorEntryResponse{
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

func (s *Service) ListEntries(ctx context.Context, req *connect.Request[notebookv1.ListEntriesRequest]) (*connect.Response[notebookv1.ListEntriesResponse], error) {
	principalId := req.Header().Get("x-principal-id")

	if principalId == "" {
		return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("no principalId"))
	}

	// Get entries after token
	entryRecords, err := s.Store.ListEntriesByAuthor(ctx, notebookstore.ListEntriesByAuthorParams{
		CreatorID: principalId,
		ID:        req.Msg.PageToken,
		Limit:     req.Msg.PageSize,
	})

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, connect.NewError(connect.CodeNotFound, fmt.Errorf(connect.CodeNotFound.String()))
		}
		s.Logger.Errorf("error getting from store: %s", err)
		return nil, connect.NewError(connect.CodeUnknown, fmt.Errorf(connect.CodeUnknown.String()))
	}

	// Get pagination info
	entries := []*notebookv1.Entry{}

	for _, entryRecord := range entryRecords {
		entries = append(entries, &notebookv1.Entry{
			Id:        entryRecord.ID,
			Text:      entryRecord.Text.String,
			CreatorId: entryRecord.CreatorID,
			CreatedAt: &timestamppb.Timestamp{
				Seconds: entryRecord.CreatedAt.Unix(),
			},
			UpdatedAt: &timestamppb.Timestamp{
				Seconds: entryRecord.UpdatedAt.Time.Unix(),
			},
		})
	}

	res := connect.NewResponse(&notebookv1.ListEntriesResponse{
		Entries:       entries,
		NextPageToken: 123,
		TotalSize:     10,
	})

	return res, nil
}

func (s *Service) BeginNewEntry(ctx context.Context, req *connect.Request[notebookv1.BeginNewEntryRequest]) (*connect.Response[notebookv1.BeginNewEntryResponse], error) {
	principalId := req.Header().Get("x-principal-id")

	if principalId == "" {
		return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("no principal"))
	}

	id, err := s.Store.CreateEntry(ctx, notebookstore.CreateEntryParams{
		Text: sql.NullString{
			String: req.Msg.Text,
			Valid:  true,
		},
		CreatorID: principalId,
		CreatedAt: time.Now(),
	})

	if err != nil {
		s.Logger.Errorf("error creating entry: %s", err)
		return nil, connect.NewError(connect.CodeUnknown, fmt.Errorf(connect.CodeUnknown.String()))
	}

	entryRecord, err := s.Store.GetEntryByIdAndAuthor(ctx, notebookstore.GetEntryByIdAndAuthorParams{
		ID:        id,
		CreatorID: principalId,
	})

	if err != nil {
		s.Logger.Errorf("error getting from store: %s", err)
		return nil, connect.NewError(connect.CodeUnknown, fmt.Errorf(connect.CodeUnknown.String()))
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
	principalId := req.Header().Get("x-principal-id")

	if principalId == "" {
		return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("no principal"))
	}

	updatedEntryRecord, err := s.Store.UpdateEntryText(ctx, notebookstore.UpdateEntryTextParams{
		ID: req.Msg.Id,
		Text: sql.NullString{
			String: req.Msg.Text,
			Valid:  true,
		},
		UpdatedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
	})

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, connect.NewError(connect.CodeNotFound, fmt.Errorf(connect.CodeNotFound.String()))
		}
		s.Logger.Errorf("error updating in store: %s", err)
		return nil, connect.NewError(connect.CodeUnknown, fmt.Errorf(connect.CodeUnknown.String()))
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

func (s *Service) DeleteEntry(ctx context.Context, req *connect.Request[notebookv1.DeleteEntryRequest]) (*connect.Response[notebookv1.DeleteEntryResponse], error) {
	principalId := req.Header().Get("x-principal-id")

	if principalId == "" {
		return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("no principal"))
	}

	deletedEntry, err := s.Store.DeleteEntryByIdAndAuthor(ctx, notebookstore.DeleteEntryByIdAndAuthorParams{
		ID:        req.Msg.Id,
		CreatorID: principalId,
		DeleteTime: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
	})

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, connect.NewError(connect.CodeNotFound, fmt.Errorf(connect.CodeNotFound.String()))
		}
		s.Logger.Errorf("error updating in store: %s", err)
		return nil, connect.NewError(connect.CodeUnknown, fmt.Errorf(connect.CodeUnknown.String()))
	}

	res := connect.NewResponse(&notebookv1.DeleteEntryResponse{
		Entry: &notebookv1.Entry{
			Id:        deletedEntry.ID,
			Text:      deletedEntry.Text.String,
			CreatorId: deletedEntry.CreatorID,
			CreatedAt: &timestamppb.Timestamp{
				Seconds: deletedEntry.CreatedAt.Unix(),
			},
			UpdatedAt: &timestamppb.Timestamp{
				Seconds: deletedEntry.UpdatedAt.Time.Unix(),
			},
			DeleteTime: &timestamppb.Timestamp{
				Seconds: deletedEntry.UpdatedAt.Time.Unix(),
			},
		},
	})

	return res, nil
}

func (s *Service) UndeleteEntry(ctx context.Context, req *connect.Request[notebookv1.UnDeleteEntryRequest]) (*connect.Response[notebookv1.UnDeleteEntryResponse], error) {
	principalId := req.Header().Get("x-principal-id")

	if principalId == "" {
		return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("no principal"))
	}

	undeletedEntry, err := s.Store.UnDeleteEntryByIdAndAuthor(ctx, notebookstore.UnDeleteEntryByIdAndAuthorParams{
		ID:        req.Msg.Id,
		CreatorID: principalId,
	})

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, connect.NewError(connect.CodeNotFound, fmt.Errorf(connect.CodeNotFound.String()))
		}
		s.Logger.Errorf("error updating in store: %s", err)
		return nil, connect.NewError(connect.CodeUnknown, fmt.Errorf(connect.CodeUnknown.String()))
	}

	res := connect.NewResponse(&notebookv1.UnDeleteEntryResponse{
		Entry: &notebookv1.Entry{
			Id:        undeletedEntry.ID,
			Text:      undeletedEntry.Text.String,
			CreatorId: undeletedEntry.CreatorID,
			CreatedAt: &timestamppb.Timestamp{
				Seconds: undeletedEntry.CreatedAt.Unix(),
			},
			UpdatedAt: &timestamppb.Timestamp{
				Seconds: undeletedEntry.UpdatedAt.Time.Unix(),
			},
		},
	})

	return res, nil
}
