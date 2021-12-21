package main

import (
	"context"
	"errors"
	"log"
	"net"
	"simplecalculator/GRPC/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	log.Println("Starting server.....")
	listener, error := net.Listen("tcp", ":8180")
	if error != nil {
		panic(error)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterBasicCalculatorServer(grpcServer, &CalculatorImpl{})
	reflection.Register(grpcServer)
	log.Println("Registered Basic calculator service")

	list_err := grpcServer.Serve(listener)
	if list_err != nil {
		panic(list_err)
	}
	log.Println("Server started")
}

type CalculatorImpl struct{}

func (calculator *CalculatorImpl) Addition(ctx context.Context, in *proto.CalculationRequest) (*proto.CalculatedResponse, error) {

	num1 := in.GetNumber1()
	num2 := in.GetNumber2()
	log.Println("Received Number1 '", num1, "' Number2 '", num2, "'")
	result := num1 + num2
	log.Println("Result of Addition '", result, "'")
	return &proto.CalculatedResponse{Result: result}, nil
}

func (calculator *CalculatorImpl) Substraction(ctx context.Context, in *proto.CalculationRequest) (*proto.CalculatedResponse, error) {

	num1 := in.GetNumber1()
	num2 := in.GetNumber2()
	log.Println("Received Number1 '", num1, "' Number2 '", num2, "'")
	result := num1 - num2
	log.Println("Result of Substraction '", result, "'")
	return &proto.CalculatedResponse{Result: result}, nil
}

func (calculator *CalculatorImpl) Multiplication(ctx context.Context, in *proto.CalculationRequest) (*proto.CalculatedResponse, error) {

	num1 := in.GetNumber1()
	num2 := in.GetNumber2()
	log.Println("Received Number1 '", num1, "' Number2 '", num2, "'")
	result := num1 * num2
	log.Println("Result of Multiplication '", result, "'")
	return &proto.CalculatedResponse{Result: result}, nil
}

func (calculator *CalculatorImpl) Division(ctx context.Context, in *proto.CalculationRequest) (*proto.CalculatedResponse, error) {

	num1 := in.GetNumber1()
	num2 := in.GetNumber2()
	log.Println("Received Number1 '", num1, "' Number2 '", num2, "'")

	if num2 == 0 {
		log.Println("Divisor cannt be 0")
		return &proto.CalculatedResponse{Result: 0}, errors.New("Divisor cannt be 0")
	}

	result := num1 / num2
	log.Println("Result of Division '", result, "'")
	return &proto.CalculatedResponse{Result: result}, nil
}
