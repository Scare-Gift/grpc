package main

import (
	pb "grpc-demo/proto"
	"log"
	"time"
)

func (s *helloServer) SayHelloServerStreaming(req *pb.NamesList, stream pb.GreetService_SayHelloServerStreamingServer) error {
	log.Printf("接收一个带有名称列表的请求：%v", req.Name)
	for _, name := range req.Name {
		res := &pb.HelloResponse{
			Message: "hello" + "  " + name,
		}
		if err := stream.Send(res); err != nil {
			return err
		}
		time.Sleep(2 * time.Second)
	}
	return nil
}
