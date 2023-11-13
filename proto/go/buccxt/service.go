package service

import "google.golang.org/grpc"

type SpotServiceClient struct {
}

func NewSpotServiceClient(conn *grpc.ClientConn) *SpotServiceClient {
	return &SpotServiceClient{}
}

type FetchDepositAddressRequest struct {
	Code string
}
