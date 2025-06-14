package game

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"game_server_slots_fortune_snake/config/system"
	. "game_server_slots_fortune_snake/constants"
	"game_server_slots_fortune_snake/db"
	betModel "game_server_slots_fortune_snake/internal/model/controller/game/bet"
	"log"
	"testing"
)

func TestRepository_SaveGameResult(t *testing.T) {
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
	repo := New(redisDB.RDB()).Mode(GameModeDemo)

	gameResult := betModel.NewGameResult(SpinModeBase)
	gameResult.SpinResult = &betModel.SpinResult{}
	gameResult.TotalScore = 10000
	data, err := json.Marshal(gameResult)
	if err != nil {
		log.Fatal(err)
		return
	}
	if err := repo.SaveGameResult(context.Background(), 123, string(data)); err != nil {
		log.Fatalf("❌: %v", err)
	}
}

func TestRepository_SaveBet(t *testing.T) {
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
	repo := New(redisDB.RDB()).Mode(GameModeDemo)

	if err := repo.SaveBet(context.Background(), 123, 1); err != nil {
		log.Fatalf("❌: %v", err)
	}
}

func TestRepository_SaveValue(t *testing.T) {
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
	repo := New(redisDB.RDB()).Mode(GameModeDemo)

	if err := repo.SaveValue(context.Background(), 123, 1000); err != nil {
		log.Fatalf("❌: %v", err)
	}
}

func TestRepository_GetGameResult(t *testing.T) {
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
	repo := New(redisDB.RDB()).Mode(GameModeDemo)

	data, err := repo.GetGameResult(context.Background(), 123)
	if err != nil {
		log.Fatalf("❌: %v", err)
	}

	gameResult := &betModel.GameResult{}
	if err = json.Unmarshal([]byte(data), gameResult); err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println(gameResult.TotalScore)
}
