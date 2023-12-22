package main

import (
	"log"
	"math/rand"
	"net"
	"time"

	wearablepb "github.com/mVedr/grpc_tut/gen/go/wearable/v1"
	"google.golang.org/grpc"
)

type wearableServer struct {
	wearablepb.UnimplementedWearableServiceServer
}

func (s *wearableServer) BeatsPerMinute(
	req *wearablepb.BeatsPerMinuteRequest,
	stream wearablepb.WearableService_BeatsPerMinuteServer) error {

	for {
		select {
		case <-stream.Context().Done():
			return nil // Change this line to return nil instead of an error
		default:
			time.Sleep(1 * time.Second)
			value := 30 + rand.Int31n(80)
			err := stream.SendMsg(&wearablepb.BeatsPerMinuteResponse{
				Value:  uint32(value),
				Minute: uint32(time.Now().Second()),
			})
			if err != nil {
				return nil // Change this line to return nil instead of an error
			}
		}
	}
}

func main() {
	lis, err := net.Listen("tcp", "localhost:9879")
	if err != nil {
		log.Fatalf("Failed to listen: %v \n", err)
	}
	grpcServer := grpc.NewServer()
	wearableServer := &wearableServer{} // Corrected the instance creation

	wearablepb.RegisterWearableServiceServer(grpcServer, wearableServer)
	grpcServer.Serve(lis)
}
