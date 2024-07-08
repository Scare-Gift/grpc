package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "grpc-demo/proto"
	"log"
)

const (
	port = ":9091"
)

func main() {
	conn, err := grpc.NewClient("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("无法连接: ", err)
	}
	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)

	names := &pb.NamesList{
		Name: []string{"gsl", "cz", "king"},
	}

	//CallSayHello(client)
	//CallSayHelloServerStream(client, names)
	//callSayHelloClientStream(client, names)
	callHelloBiDirectionalStream(client, names)
}
