package main

import (
	"fmt"
	greet "grpc-service/grpc-service/pb/go"
	"grpc-service/grpc-service/pb/handler"

	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	port := ":8080"
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("error:%v", err)
	} else {
		log.Printf("Successfully connected to port %v\n", port)
	}

	x := handler.Greet{}
	grpc := grpc.NewServer()
	reflection.Register(grpc)

	greet.RegisterServer(grpc, x)
	if err := grpc.Serve(listen); err != nil {
		fmt.Println(err)
	}
}
