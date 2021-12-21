package main

import (
	"errors"
	"log"
	"net/http"
	"simplecalculator/GRPC/proto"
	"strconv"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

func main() {
	log.Println("Starting client....")
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/add/{num1}/{num2}", add)
	router.HandleFunc("/substract/{num1}/{num2}", substract)
	router.HandleFunc("/multiply/{num1}/{num2}", multiply)
	router.HandleFunc("/divide/{num1}/{num2}", divide)
	log.Println("Registered Endpoints")

	log.Fatal(http.ListenAndServe(":8181", router))
	log.Println("Client Started")

}

func add(resp http.ResponseWriter, req *http.Request) {

	log.Println("Parsing Request.")
	num1, num2, parse_error := parseRequest(req)
	if parse_error != nil {
		panic(parse_error)
	}

	log.Println("Connecting to GRPC server.....")
	connection, conn_err := grpc.Dial("localhost:8180", grpc.WithInsecure())
	if conn_err != nil {
		panic(connection)
	}

	calc_client := proto.NewBasicCalculatorClient(connection)
	log.Println("connection Established")
	response, calc_error := calc_client.Addition(req.Context(), &proto.CalculationRequest{Number1: num1, Number2: num2})

	writeResponse(response, resp, calc_error)

}

func substract(resp http.ResponseWriter, req *http.Request) {
	log.Println("Parsing Request.")
	num1, num2, parse_error := parseRequest(req)
	if parse_error != nil {
		panic(parse_error)
	}

	log.Println("Connecting to GRPC server.....")
	connection, conn_err := grpc.Dial("localhost:8180", grpc.WithInsecure())
	if conn_err != nil {
		panic(connection)
	}

	calc_client := proto.NewBasicCalculatorClient(connection)
	log.Println("connection Established")
	response, calc_error := calc_client.Substraction(req.Context(), &proto.CalculationRequest{Number1: num1, Number2: num2})

	writeResponse(response, resp, calc_error)
}

func multiply(resp http.ResponseWriter, req *http.Request) {
	log.Println("Parsing Request.")
	num1, num2, parse_error := parseRequest(req)
	if parse_error != nil {
		panic(parse_error)
	}

	log.Println("Connecting to GRPC server.....")
	connection, conn_err := grpc.Dial("localhost:8180", grpc.WithInsecure())
	if conn_err != nil {
		panic(connection)
	}

	calc_client := proto.NewBasicCalculatorClient(connection)
	log.Println("connection Established")
	response, calc_error := calc_client.Multiplication(req.Context(), &proto.CalculationRequest{Number1: num1, Number2: num2})

	writeResponse(response, resp, calc_error)

}

func divide(resp http.ResponseWriter, req *http.Request) {

	log.Println("Parsing Request.")
	num1, num2, parse_error := parseRequest(req)
	if parse_error != nil {
		panic(parse_error)
	}

	log.Println("Connecting to GRPC server.....")
	connection, conn_err := grpc.Dial("localhost:8180", grpc.WithInsecure())
	if conn_err != nil {
		panic(connection)
	}

	calc_client := proto.NewBasicCalculatorClient(connection)
	log.Println("connection Established")
	response, calc_error := calc_client.Division(req.Context(), &proto.CalculationRequest{Number1: num1, Number2: num2})

	writeResponse(response, resp, calc_error)

}

func writeResponse(response *proto.CalculatedResponse, writer http.ResponseWriter, calc_err error) {
	writer.Header().Add("Content-Type", "text/plain")
	if calc_err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		var error string = calc_err.Error()
		log.Fatal("error - ", error)
		writer.Write([]byte(error))
	} else {
		result := strconv.Itoa(int(response.Result))
		writer.WriteHeader(http.StatusAccepted)
		writer.Write(([]byte)(result))
	}
}

func parseRequest(req *http.Request) (int64, int64, error) {
	reqParams := mux.Vars(req)
	num1, num1Found := reqParams["num1"]

	if !num1Found {
		log.Println("Parameter cannot be empty")
		return 0, 0, errors.New("Parameter cannot be empty")
	}

	num2, num2Found := reqParams["num2"]
	if !num2Found {
		log.Println("Parameter cannot be empty")
		return 0, 0, errors.New("Parameter cannot be empty")
	}

	convertedNum1, _ := strconv.ParseInt(num1, 10, 64)
	convertedNum2, _ := strconv.ParseInt(num2, 10, 64)
	return convertedNum1, convertedNum2, nil
}
