package main

import (
	"context"
	pb "grpc-demo/proto"
	"log"
	"time"
)

func CallSayHello(client pb.GreetServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := client.SayHello(ctx, &pb.NoParam{})
	if err != nil {
		log.Fatalf("SayHello err: %v", err)
	}
	log.Print("%" + res.Message)

}
