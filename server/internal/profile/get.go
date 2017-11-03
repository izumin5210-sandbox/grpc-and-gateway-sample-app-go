package profile

import (
	"context"

	"github.com/izumin5210-sandbox/grpc-and-gateway-sample-app-go/api"
)

func (s *server) Get(context.Context, *api.GetProfileRequest) (*api.Profile, error) {
	return &api.Profile{
		Name:     "Masayuki Izumi",
		Location: "Tokyo",
		WorkingHistories: []*api.WorkingHistory{
			{
				Company: "Wantedly, Inc.",
			},
		},
	}, nil
}
