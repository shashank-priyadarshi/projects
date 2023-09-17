package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	pb "dbconn/db"
)

// main is the entry point for the program
func main() {
	// Listen for incoming TCP connections on port 8080
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Create a new gRPC server
	s := grpc.NewServer()

	// Register the database service with the server
	pb.RegisterDatabaseServiceServer(s, pb.UnimplementedDatabaseServiceServer{})

	// Start the gRPC server
	log.Println("Starting server on :8080")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
