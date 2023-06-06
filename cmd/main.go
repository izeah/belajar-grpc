package main

import (
	"log"
	"net"

	"belajar-grpc/cmd/services"
	productPb "belajar-grpc/pb/product"

	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":50001")
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	productService := services.ProductService{}
	productPb.RegisterProductServiceServer(grpcServer, &productService)

	log.Printf("Server running on port %v", listener.Addr())
	if err = grpcServer.Serve(listener); err != nil {
		panic(err)
	}
}
