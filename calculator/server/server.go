package main

import pb "github.com/eranga-stack/grpc-go/calculator/proto"

type Server struct {
	pb.CalculatorServiceServer
}
