package main

import (
	"context"
	pb "grpc-demo/proto"
	"log"
	"time"
)

func callSayHelloClientStream(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("Client streaming started")
	stream, err := client.SayHelloClientStreaming(context.Background())
	if err != nil {
		log.Fatalf("无法发送names：%v", err)
	}
	for _, name := range names.Name {
		req := &pb.HelloRequest{
			Message: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("发送消息时发生错误:%v", err)
		}
		log.Printf("发生携带name的请求: %v", name)
		time.Sleep(2 * time.Second)
	}
	res, err := stream.CloseAndRecv()
	log.Printf("streaming fished")
	if err != nil {
		log.Fatalf("接收响应时发生错误:%v", err)
	}
	log.Printf("%v", res.Message)
}
