package result_free

import (
	"context"
	"flag"
	"fmt"
	"game_server_slots_fortune_snake/config/system"
	"game_server_slots_fortune_snake/constants"
	"game_server_slots_fortune_snake/db"
	resultFreeRepo "game_server_slots_fortune_snake/internal/repository/result_free"
	"log"
	"testing"
)

func TestService_SaveItems(t *testing.T) {
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
	items := make([][][]int, 0)
	items = append(items, [][]int{{6, 0, 1}, {0, 0, 0, 0}, {1, 2, 3}})
	items = append(items, [][]int{{6, 0, 2}, {0, 0, 0, 0}, {1, 2, 3}})
	items = append(items, [][]int{{6, 0, 3}, {0, 0, 0, 0}, {1, 2, 3}})
	items = append(items, [][]int{{6, 0, 4}, {0, 0, 0, 0}, {1, 2, 3}})
	items = append(items, [][]int{{6, 0, 5}, {0, 0, 0, 0}, {1, 2, 3}})

	// 創建 repo
	repo := resultFreeRepo.New(redisDB.RDB()).GameMode(constants.GameModeDemo)
	svc := New(repo)

	if err = svc.GameMode(constants.GameModeDemo).SaveItems(context.Background(), 123, items); err != nil {
		log.Fatalf("❌ Redis 存取失败: %v", err)
		return
	}
}

func TestService_PopFirstItem(t *testing.T) {
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
	repo := resultFreeRepo.New(redisDB.RDB()).GameMode(constants.GameModeDemo)
	svc := New(repo)

	result, err := svc.GameMode(constants.GameModeDemo).PopFirstItem(context.Background(), 123)
	if err != nil {
		log.Fatalf("❌ Redis 存取失败: %v", err)
		return
	}
	fmt.Println(result)
}
