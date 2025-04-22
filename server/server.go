package server

import (
	"context"
	"encoding/json"
	"fmt"
	"game_server_slots_fortune_snake/constants"
	"game_server_slots_fortune_snake/model"
	"game_server_slots_fortune_snake/pkg"
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
	engine     *Engine
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
		engine:     New(),
	}
}

func (s *Server) SendMessage(ctx context.Context, req *pb.MessageRequest) (*pb.MessageResponse, error) {
	res := &pb.MessageResponse{}
	// 路由解析
	path := s.engine.resolver([]byte(req.Action))
	// 路由選擇
	handlers, ok := s.engine.route.Get(path)
	if !ok {
		res.Code = constants.CodeBadRequest
		res.Message = pkg.LocalizeInstance().LocalizeMessage("Failure")
		return res, nil
	}
	// 創建 context
	engineCtx := Context{
		ctx:      ctx,
		engine:   s.engine,
		handlers: handlers,
		keys:     make(map[string]interface{}),
		index:    -1,
		data:     []byte(res.Data),
		output:   []byte{},
	}
	// 執行路由
	engineCtx.Next()
	// 將 output 轉換為 model
	resModel := &model.MessageResponse{}
	if err := json.Unmarshal(engineCtx.output, resModel); err != nil {
		res.Code = constants.CodeBadRequest
		res.Message = pkg.LocalizeInstance().LocalizeMessage("Failure")
		res.Data = err.Error()
		return res, nil
	}
	res.Code = resModel.Code
	res.Message = resModel.Message
	res.Data = resModel.Data
	return res, nil
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
