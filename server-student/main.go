package main

import (
	"log"
	"net"
	database "w00k/go/grpc/database"
	"w00k/go/grpc/server"
	"w00k/go/grpc/studentpb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	list, err := net.Listen("tcp", "5060")
	if err != nil {
		log.Fatal(err)
	}

	repo, err := database.NewPostgresRepository("postgres://postgres:postgres@localhost:54321/postgres?sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}

	server := server.NewServer(repo)
	s := grpc.NewServer()
	studentpb.RegisterStudentServiceServer(s, server)
	reflection.Register(s)

	if err := s.Serve(list); err != nil {
		log.Fatal(err)
	}
}
