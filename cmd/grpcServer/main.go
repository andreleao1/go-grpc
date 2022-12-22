package main

import (
	"database/sql"
	"net"

	"exemple.com/grpc/internal/database"
	"exemple.com/grpc/internal/pb"
	"exemple.com/grpc/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"github.com/mattn/go-sqlite3"
)

func main() {
	db, error := sql.Open("sqlite3", "./db.sqlite")
	if error != nil {
		panic(error)
	}
	defer db.Close()

	categoryDb := database.NewCategory(db)
	categoryService := service.NewCategoryService(*categoryDb)

	grpcServer := grpc.NewServer()
	pb.RegisterCategoryServiceServer(grpcServer, categoryService)
	reflection.Register(grpcServer)

	list, error := net.Listen("tcp", ":50051")
	if error != nil {
		panic(error)
	}

	if error := grpcServer.Serve(list); error != nil {
		panic(error)
	}
}
