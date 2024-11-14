package main

import (
	"context"
	"log"
	"time"

	pb "github.com/Kyadlapati/hospitality/proto"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	client := pb.NewCarCommuteServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.NotifyResidents(ctx, &pb.NotifyRequest{
		ApartmentNumber: "2011",
		Message:         "Come to building 15",
	})
	if err != nil {
		log.Fatalf("Error calling NotifyResidents: %v", err)
	}
	log.Printf("Notification response: %v", res.Success)
}
