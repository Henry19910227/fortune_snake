package result_free

import (
	"context"
	"flag"
	"fmt"
	"game_server_slots_fortune_snake/config/system"
	"game_server_slots_fortune_snake/constants"
	"game_server_slots_fortune_snake/db"
	"log"
	"testing"
)

func TestRepository_SaveItems(t *testing.T) {
	// 加載 yaml
	configFile := flag.String("config", "config.yaml", "YAML configuration file name")
	flag.Parse()

	// 加载配置文件
	appConfig := system.New(configFile)

	redisDB, err := db.NewRedisDB(appConfig.Config().Redis)
	if err != nil {
		log.Fatalf("❌ Redis 初始化失败: %v", err)
	}
	defer redisDB.Close()

	// 準備數據
	items := make([]string, 0)
	items = append(items, "[[6,0,1],[0,0,0,0],[1,2,6]]")
	items = append(items, "[[5,6,1],[0,0,0,0],[6,6,0]]")
	items = append(items, "[[6,1,6],[0,0,0,0],[6,0,3]]")
	items = append(items, "[[1,2,6],[0,0,0,0],[3,0,6]]")
	items = append(items, "[[1,4,4],[0,0,0,0],[1,4,4]]")

	// 創建 repo
	repo := New(redisDB.RDB()).GameMode(constants.GameModeDemo)

	if err = repo.SaveItems(context.Background(), 123, items); err != nil {
		log.Fatalf("❌ Redis 存取失败: %v", err)
		return
	}
}

func TestRepository_PopFirstItem(t *testing.T) {
	// 加載 yaml
	configFile := flag.String("config", "config.yaml", "YAML configuration file name")
	flag.Parse()

	// 加载配置文件
	appConfig := system.New(configFile)

	redisDB, err := db.NewRedisDB(appConfig.Config().Redis)
	if err != nil {
		log.Fatalf("❌ Redis 初始化失败: %v", err)
	}
	defer redisDB.Close()

	// 創建 repo
	repo := New(redisDB.RDB()).GameMode(constants.GameModeDemo)

	result, err := repo.PopFirstItem(context.Background(), 123)
	if err != nil {
		log.Fatalf("❌: %v", err)
	}
	fmt.Println(result)
}
