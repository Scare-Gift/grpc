package main

import (
	"context"
	pb "grpc-demo/proto"
	"io"
	"log"
)

func CallSayHelloServerStream(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("stream start")
	stream, err := client.SayHelloServerStreaming(context.Background(), names)
	if err != nil {
		log.Printf("发送失败", err)
	}
	for {
		message, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error while streaming %v", err)
		}
		log.Printf(message.GetMessage())
	}
	log.Fatalf("stream finished")
}
