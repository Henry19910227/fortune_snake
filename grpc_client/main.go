package main

import (
	"context"
	"encoding/json"
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
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()
	m := Param{
		Bet:   1,
		Value: 1000,
	}
	b, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)
		return
	}
	resp, err := client.SendMessage(ctx, &proto.MessageRequest{Action: "bet", PlayerId: 6917529027641081862, Data: string(b)})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resp.Code, resp.Message, resp.Data)
}
