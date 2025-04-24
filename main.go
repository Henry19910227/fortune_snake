package main

import (
	"context"
	"flag"
	"game_server_slots_fortune_snake/config"
	"game_server_slots_fortune_snake/db"
	"game_server_slots_fortune_snake/etcd"
	controllerFactory "game_server_slots_fortune_snake/internal/factory/controller"
	repoFactory "game_server_slots_fortune_snake/internal/factory/repository"
	serviceFactory "game_server_slots_fortune_snake/internal/factory/service"
	"game_server_slots_fortune_snake/internal/router/game"
	"game_server_slots_fortune_snake/internal/server"
	"game_server_slots_fortune_snake/pkg"
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

	// 初始化工廠
	repoFact := repoFactory.New(mysqlDB.DB(), redisDB.RDB())
	serviceFact := serviceFactory.New(repoFact)
	factory := controllerFactory.New(serviceFact)

	// 9. 創建 gRPC Engine
	wg.Add(1)
	engine := server.New(appConfig.Config().Server)

	// 添加路由解析器邏輯
	engine.PathResolver(func(b []byte) string {
		return "/" + string(b)
	})

	// 設定Base路由組
	baseGroup := engine.Group("/")
	baseGroup.Use(factory.MiddleController().UnMarshalData)
	// 添加路由
	game.SetRoute(baseGroup, factory)

	// 9. 启动 gRPC Engine
	go func() {
		defer wg.Done()
		engine.Run(ctx)
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
