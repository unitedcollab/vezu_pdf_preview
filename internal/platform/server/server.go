package server

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

func NewServer() *grpc.Server {
	server := grpc.NewServer()

	reflection.Register(server)
	grpc_health_v1.RegisterHealthServer(server, health.NewServer())

	return server
}
