package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/LanfordCai/ava/api/rpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	startRPCServer()
}

func startRPCServer() {
	port := os.Getenv("AVA_API_PORT")
	if port == "" {
		port = "5000"
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("failed to listen port %s, error: %v", port, err)
	}
	s := rpc.NewAddressValidatorServer()

	if env() == "DEV" {
		reflection.Register(s)
	}

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func env() string {
	switch os.Getenv("AVA_ENV") {
	case "PROD":
		return "PROD"
	default:
		return "DEV"
	}
}
