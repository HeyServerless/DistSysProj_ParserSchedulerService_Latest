package grpcservice

import (
	"context"
	"log"
	"strconv"

	// path to your proto file Hello.proto
	pbAdd "github.com/parserSchedulerService/grpcclient/ADD"

	pbSubtract "github.com/parserSchedulerService/grpcclient/SUBTRACT"

	pbDivision "github.com/parserSchedulerService/grpcclient/DIVISION"

	pbMultiply "github.com/parserSchedulerService/grpcclient/MULTIPLY"

	"google.golang.org/grpc"
)

const (
	// IPADDRESS = "192.168.59.101" // minikube
	IPADDRESS = "52.7.218.132" //cluster ip
	// IPADDRESS = "localhost"
)

func CallGrpcAdd(port int, A float32, B float32) (float32, error) {
	// Create a connection to the gRPC server
	log.Println("Calling Add")
	log.Println("IPADDRESS: ", IPADDRESS+":"+strconv.Itoa(port))
	conn, err := grpc.Dial(IPADDRESS+":"+strconv.Itoa(port), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial server: %v", err)
	}
	defer conn.Close()

	// Create a client for the YourService service
	client := pbAdd.NewAddServiceClient(conn)

	// Call a method on the client
	resp, err := client.AddMethod(context.Background(), &pbAdd.AddRequest{
		A: A,
		B: B,
	})
	if err != nil {
		log.Fatalf("Failed to call AddMethod: %v", err)
	}

	// Process the response from the server
	log.Printf("Response from server: %v", resp.Sum)
	return float32(resp.Sum), nil
}

func CallGrpcMultiply(port int, A float32, B float32) (float32, error) {
	// Create a connection to the gRPC server
	log.Println("Calling Multiply")
	log.Println("IPADDRESS: ", IPADDRESS+":"+strconv.Itoa(port))
	conn, err := grpc.Dial(IPADDRESS+":"+strconv.Itoa(port), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial server: %v", err)
	}
	defer conn.Close()

	// Create a client for the YourService service
	client := pbMultiply.NewMultiplyServiceClient(conn)

	// Call a method on the client
	resp, err := client.MultiplyMethod(context.Background(), &pbMultiply.MultiplyRequest{
		A: A,
		B: B,
	})
	if err != nil {
		log.Fatalf("Failed to call MultiplyMethod: %v", err)
	}

	// Process the response from the server
	log.Printf("Response from server: %v", resp.Result)
	return resp.Result, nil
}

func CallGrpcSubtract(port int, A float32, B float32) (float32, error) {
	// Create a connection to the gRPC server
	log.Println("Calling Subtract")
	log.Println("IPADDRESS: ", IPADDRESS+":"+strconv.Itoa(port))
	conn, err := grpc.Dial(IPADDRESS+":"+strconv.Itoa(port), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial server: %v", err)
	}
	defer conn.Close()

	// Create a client for the YourService service
	client := pbSubtract.NewSubtractServiceClient(conn)

	// Call a method on the client
	resp, err := client.SubtractMethod(context.Background(), &pbSubtract.SubtractRequest{
		A: A,
		B: B,
	})
	if err != nil {
		log.Fatalf("Failed to call SubtractMethod: %v", err)
	}

	// Process the response from the server
	log.Printf("Response from server: %v", resp.Result)
	return resp.Result, nil
}

func CallGrpcDivision(port int, A float32, B float32) (float32, error) {
	// Create a connection to the gRPC server
	log.Println("Calling Division")
	log.Println("IPADDRESS: ", IPADDRESS+":"+strconv.Itoa(port))
	conn, err := grpc.Dial(IPADDRESS+":"+strconv.Itoa(port), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial server: %v", err)
	}
	defer conn.Close()

	// Create a client for the YourService service
	client := pbDivision.NewDivisionServiceClient(conn)

	// Call a method on the client
	resp, err := client.DivisionMethod(context.Background(), &pbDivision.DivisionRequest{
		A: A,
		B: B,
	})
	if err != nil {
		log.Fatalf("Failed to call DivisionMethod: %v", err)
	}

	// Process the response from the server
	log.Printf("Response from server: %v", resp.Result)
	return resp.Result, nil
}
