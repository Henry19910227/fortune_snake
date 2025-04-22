package server

import (
	"context"
	"fmt"
	"game_server_slots_fortune_snake/model"
	pb "game_server_slots_fortune_snake/proto"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

type Server struct {
	pb.UnimplementedMessageServiceServer
	grpcServer *grpc.Server
	listener   net.Listener
	config     model.ServerConfig
}

func NewServer(config model.ServerConfig) *Server {
	address := fmt.Sprintf("%s:%d", config.Ip, config.Port)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return nil
	}
	grpcServer := grpc.NewServer()
	pb.RegisterMessageServiceServer(grpcServer, &Server{})

	return &Server{
		grpcServer: grpcServer,
		listener:   lis,
		config:     config,
	}
}

func (s *Server) RunWithRetry(ctx context.Context) {
	var retryCount int

	for {
		select {
		case <-ctx.Done():
			log.Println("收到退出信号，停止启动 gRPC 服务")
			return
		default:
		}
		// 啟動成功
		err := s.Run(ctx)
		if err == nil {
			break
		}
		// 啟動失敗重試
		retryCount++
		log.Printf("❌ 服务启动失败, 正在重试第 %d 次, 错误详情: %v", retryCount, err)
		select {
		case <-ctx.Done():
			log.Println("收到退出信号，停止重试 gRPC 服务启动")
			return
		case <-time.After(5 * time.Second):
		}
	}
}

func (s *Server) Run(ctx context.Context) error {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("gRPC 發生 panic: %v", r)
		}
	}()

	go func() {
		<-ctx.Done()
		log.Println("收到關閉信號，優雅關閉 gRPC 服務...")
		s.grpcServer.GracefulStop()
	}()

	log.Printf("啟動 gRPC 服務：%s:%d", s.config.Ip, s.config.Port)
	return s.grpcServer.Serve(s.listener)
}
