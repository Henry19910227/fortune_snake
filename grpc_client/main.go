package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc_client/proto"
	"time"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8181", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		fmt.Printf("did not connect: %v\n", err)
	}
	defer conn.Close()
	client := proto.NewMessageServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	resp, err := client.SendMessage(ctx, &proto.MessageRequest{Action: "player", PlayerId: 10376293541461622785, Data: ""})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resp.Code, resp.Message, resp.Data)
}
