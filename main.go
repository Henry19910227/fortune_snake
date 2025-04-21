package main

import (
	"flag"
	"game_server_slots_fortune_snake/config"
	"game_server_slots_fortune_snake/db"
	"game_server_slots_fortune_snake/pkg"
	"game_server_slots_fortune_snake/utils"
	"log"
)

func main() {
	// 0. 加載 yaml
	configFile := flag.String("config", "config.yaml", "YAML configuration file name")
	flag.Parse()

	// 1. 加载配置文件
	appConfig := config.New(configFile)

	// 2. 初始化语言包
	_ = pkg.InitLocalize(appConfig.Config().Language.Default)

	// 3. 初始化分布式雪花 ID
	_, err := utils.NewSnowFlake(appConfig.Config().Server.ServerNode)
	if err != nil {
		log.Fatalf("❌ 分布式雪花 ID 初始化失败: %v", err)
	}

	// 4. 创建全局 context，用于控制所有服务的退出，并传递给需要优雅关闭的服务

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
}
