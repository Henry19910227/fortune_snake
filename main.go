package main

import (
	"context"
	"flag"
	"game_server_slots_fortune_snake/config"
	"game_server_slots_fortune_snake/db"
	"game_server_slots_fortune_snake/etcd"
	"game_server_slots_fortune_snake/pkg"
	"game_server_slots_fortune_snake/server"
	"game_server_slots_fortune_snake/utils"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	// 0. 加載 yaml
	configFile := flag.String("config", "config.yaml", "YAML configuration file name")
	flag.Parse()

	// 1. 加载配置文件
	appConfig := config.New(configFile)

	// 2. 初始化语言包
	pkg.InitLocalizeInstance(appConfig.Config().Language.Default)

	// 3. 初始化分布式雪花 ID
	_, err := utils.NewSnowFlake(appConfig.Config().Server.ServerNode)
	if err != nil {
		log.Fatalf("❌ 分布式雪花 ID 初始化失败: %v", err)
	}

	// 4. 创建全局 context，用于控制所有服务的退出，并传递给需要优雅关闭的服务
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var wg sync.WaitGroup

	// 5. 初始化 MySQL DB
	mysqlDB, err := db.NewMysqlDB(appConfig.Config().Database)
	if err != nil {
		log.Fatalf("❌ MySQL 初始化失败: %v", err)
	}
	defer mysqlDB.Close()

	// 6. 初始化 Redis
	redisDB, err := db.NewRedisDB(appConfig.Config().Redis)
	if err != nil {
		log.Fatalf("❌ Redis 初始化失败: %v", err)
	}
	defer redisDB.Close()

	// 7. 初始化 MongoDB
	mongoDB, err := db.NewMongoDB(appConfig.Config().Mongodb)
	if err != nil {
		log.Fatalf("❌ MongoDB 初始化失败: %v", err)
	}
	defer mongoDB.Close()

	// 8. 初始化 etcd 客户端
	if err := etcd.InitEtcd(appConfig.Config().Etcd); err != nil {
		log.Printf("⚠️ Etcd 连接失败: %v（服务器仍会继续运行）", err)
	} else if err := etcd.RegisterService(*appConfig.Config()); err != nil {
		log.Printf("⚠️ Etcd 小火箭服务注册失败: %v", err)
	} else {
		defer etcd.UnregisterService(appConfig.Config().Etcd)
	}

	// 9. 启动 gRPC 服务
	wg.Add(1)
	grpcServer := server.NewServer(appConfig.Config().Server)
	go func() {
		defer wg.Done()
		grpcServer.RunWithRetry(ctx)
	}()

	// 10. 捕获系统退出信号
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	<-sigCh
	log.Println("🛑 收到关闭信号，开始优雅关闭所有服务...")
	cancel() // 通知所有服务退出
	wg.Wait()
	log.Println("✅ 所有服务已关闭，程序退出")
}
