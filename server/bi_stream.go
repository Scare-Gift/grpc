package main

import (
	pb "grpc-demo/proto"
	"io"
	"log"
)

func (s *helloServer) SayHelloBiDirectionalStreaming(stream pb.GreetService_SayHelloByDirectionalStreamingServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Printf("接收携带name的请求 : %v", req.Message)
		res := &pb.HelloResponse{
			Message: "Hello" + "  " + req.Message,
		}
		if err := stream.Send(res); err != nil {
			return err
		}
	}
}
