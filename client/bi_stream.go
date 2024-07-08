package main

import (
	"golang.org/x/net/context"
	pb "grpc-demo/proto"
	"io"
	"log"
	"time"
)

func callHelloBiDirectionalStream(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("Bidirectional Streaming Strrted")
	stream, err := client.SayHelloByDirectionalStreaming(context.Background())
	if err != nil {
		log.Fatal("无法发送消息: %v", err)
	}
	waitc := make(chan struct{})
	go func() {
		for {
			message, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal("发送未知错误 ", err)
			}
			log.Printf(message.GetMessage())
		}
		close(waitc)
	}()
	for _, name := range names.Name {
		req := &pb.HelloRequest{
			Message: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatal("发生未知的错误 %v", err)
		}
		time.Sleep(2 * time.Second)
	}
	stream.CloseSend()
	<-waitc
	log.Printf("Bidirectional Streaming finished")
}
