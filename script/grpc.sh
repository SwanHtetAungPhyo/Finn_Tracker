#!/bin/bash

# Set the directory for your User Service project
PROJECT_DIR="user_service"
PROTO_DIR="$PROJECT_DIR/proto"
GENERATED_DIR="$PROJECT_DIR/generated"
PROTO_FILE="$PROTO_DIR/user_service.proto"

# Step 1: Create project directory if it doesn't exist
mkdir -p $PROTO_DIR
mkdir -p $GENERATED_DIR

# Step 2: Define the .proto file
echo "Creating user_service.proto..."
cat <<EOF > $PROTO_FILE
syntax = "proto3";

package user;

option go_package = "github.com/your-username/user_service/generated";

service UserService {
    rpc GetUser (GetUserRequest) returns (GetUserResponse);
}

message GetUserRequest {
    int32 id = 1;
}

message GetUserResponse {
    int32 id = 1;
    string name = 2;
    bool exists = 3;
}
EOF

# Step 3: Install dependencies
echo "Installing necessary Go dependencies..."

# Ensure Go modules are initialized
cd $PROJECT_DIR
go mod init github.com/your-username/user_service
go get google.golang.org/grpc
go get github.com/golang/protobuf/protoc-gen-go

# Step 4: Install Protocol Buffers and Go plugins if not installed
echo "Installing Protocol Buffers compiler and Go plugins..."

# Check if protoc-gen-go is in the PATH, otherwise install it
if ! command -v protoc-gen-go &> /dev/null; then
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
fi

# Check if protoc-gen-go-grpc is in the PATH, otherwise install it
if ! command -v protoc-gen-go-grpc &> /dev/null; then
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
fi

# Step 5: Generate Go code from .proto file
echo "Generating gRPC Go code from user_service.proto..."
protoc --go_out=$GENERATED_DIR --go-grpc_out=$GENERATED_DIR $PROTO_FILE

# Step 6: Create the main Go file for the User Service gRPC server
echo "Creating the main.go file for the gRPC server..."
cat <<EOF > $PROJECT_DIR/main.go
package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/your-username/user_service/generated"
	"google.golang.org/grpc"
)

type server struct {
	generated.UnimplementedUserServiceServer
}

// GetUser handles the GetUser request
func (s *server) GetUser(ctx context.Context, req *generated.GetUserRequest) (*generated.GetUserResponse, error) {
	if req.Id == 1 {
		// Simulate fetching user from a database
		return &generated.GetUserResponse{
			Id:     1,
			Name:   "John Doe",
			Exists: true,
		}, nil
	}
	return &generated.GetUserResponse{
		Id:     req.Id,
		Name:   "",
		Exists: false,
	}, nil
}

func main() {
	// Set up gRPC server
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	// Register the UserService implementation with the gRPC server
	generated.RegisterUserServiceServer(grpcServer, &server{})

	// Start the server
	fmt.Println("User Service is running on port 50051...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
EOF

# Step 7: Run the server
echo "Running the User Service gRPC server..."
cd $PROJECT_DIR
go run main.go
