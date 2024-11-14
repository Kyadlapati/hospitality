//go:build server

package main

import (
	"context"
	"log"
	"net"

	pb "github.com/Kyadlapati/hospitality/proto"

	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedCarCommuteServiceServer
}

func (s *Server) NotifyResidents(ctx context.Context, req *pb.NotifyRequest) (*pb.NotifyResponse, error) {
	log.Printf("Sending notification to apartment %s: %s", req.ApartmentNumber, req.Message)
	return &pb.NotifyResponse{Success: true}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen on port 50051: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterCarCommuteServiceServer(grpcServer, &Server{})

	log.Println("gRPC server is running on port 50051")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
