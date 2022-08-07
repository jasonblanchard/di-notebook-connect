package notebookv1

import (
	"context"

	"github.com/bufbuild/connect-go"
	notebookv1 "github.com/jasonblanchard/di-notebook-connect/gen/proto/go/notebookapis/notebook/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Service struct{}

func (s *Service) GetEntry(ctx context.Context, req *connect.Request[notebookv1.GetEntryRequest]) (*connect.Response[notebookv1.GetEntryResponse], error) {
	res := connect.NewResponse(&notebookv1.GetEntryResponse{
		Entry: &notebookv1.Entry{
			Id:        "1",
			Text:      "I am an entry",
			CreatorId: "123",
			CreatedAt: &timestamppb.Timestamp{
				Seconds: 1659831878,
			},
			UpdatedAt: &timestamppb.Timestamp{
				Seconds: 1659831878,
			},
		},
	})
	return res, nil
}

func (s *Service) ListEntries(ctx context.Context, req *connect.Request[notebookv1.ListEntriesRequest]) (*connect.Response[notebookv1.ListEntriesResponse], error) {
	res := connect.NewResponse(&notebookv1.ListEntriesResponse{
		NextPageToken: "TODO",
		TotalSize:     123,
		HasNextPage:   true,
		Entries: []*notebookv1.Entry{
			{
				Id:        "123",
				Text:      "I am an entry",
				CreatorId: "123",
				CreatedAt: &timestamppb.Timestamp{
					Seconds: 1659831878,
				},
				UpdatedAt: &timestamppb.Timestamp{
					Seconds: 1659831878,
				},
			},
			{
				Id:        "456",
				Text:      "I am another entry",
				CreatorId: "123",
				CreatedAt: &timestamppb.Timestamp{
					Seconds: 1659831878,
				},
				UpdatedAt: &timestamppb.Timestamp{
					Seconds: 1659831878,
				},
			},
		},
	})

	return res, nil
}
