package main

import (
	"context"
	pb "day3/server/proto/consignment"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

const (
	port = ":50051"
)

type repository interface {
	Create(command *pb.Command) (*pb.Command, error)
	GetAll() []*pb.Command
}

type Repository struct {
	commands []*pb.Command
}

func (r *Repository) Create(command *pb.Command) (*pb.Command, error) {
	updatedCommands := append(r.commands, command)
	r.commands = updatedCommands
	return command, nil
}

func (r *Repository) GetAll() []*pb.Command {
	return r.commands
}

type service struct {
	repo repository
}

func (s *service) CreateCommand(ctx context.Context, req *pb.Command) (*pb.Response, error) {
	command, err := s.repo.Create(req)
	if err != nil {
		return nil, err
	}
	log.Println("Request to create command")
	return &pb.Response{
		Created: true,
		Command: command,
	}, nil
}

func (s *service) GetAllCommands(ctx context.Context, req *pb.GetRequest) (*pb.Response, error) {
	commands := s.repo.GetAll()
	log.Println("Request to get all command")
	return &pb.Response{Commands: commands}, nil
}

func main() {
	repo := &Repository{}

	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen port: %v", err)
	}
	server := grpc.NewServer()
	ourService := &service{repo: repo}
	pb.RegisterShippingServiceServer(server, ourService)

	reflection.Register(server)

	log.Printf("gRPPC server running on port: %v", port)
	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to server from port: %v", err)
	}
}
