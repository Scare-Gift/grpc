package main

import (
	pb "grpc-demo/proto"
	"io"
	"log"
)

func (s *helloServer) SayHelloClientStreaming(stream pb.GreetService_SayHelloClientStreamingServer) error {
	var messages []string
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.MessagesList{Message: messages})
		}
		if err != nil {
			return err
		}
		log.Printf("接收带有name的请求：%v", req.GetMessage())
		messages = append(messages, "hello", req.GetMessage())
	}
	return nil
}
