package tripservice

import (
	"context"

	trippb "github.com/iralance/go-looklook/basic/grpc/proto/gen/go"
)

type ServiceServer struct {
	*trippb.UnimplementedTripServiceServer
}

func (t ServiceServer) GetTrip(ctx context.Context, request *trippb.GetTripRequest) (*trippb.GetTripResponse, error) {
	return &trippb.GetTripResponse{
		Id: request.Id,
		Trip: &trippb.Trip{
			Start:       "abc",
			End:         "张三",
			DurationSec: 3600,
			FeeCent:     10000,
			StartPos: &trippb.Location{
				Latitude:  30,
				Longitude: 120,
			},
			EndPos: &trippb.Location{
				Latitude:  35,
				Longitude: 115,
			},
			PathLocations: []*trippb.Location{
				{
					Latitude:  31,
					Longitude: 119,
				},
				{
					Latitude:  32,
					Longitude: 118,
				},
			},
			Status: trippb.TripStatus_PAID,
		},
	}, nil
}
