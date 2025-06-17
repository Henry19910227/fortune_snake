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
	"game_server_slots_fortune_snake/internal/model/repository/game/save"
	"game_server_slots_fortune_snake/util"
	"log"
	"testing"
)

func TestRepository_Save(t *testing.T) {
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
	repo := New(redisDB.RDB())

	gameResult := betModel.NewGameResult(SpinModeBase)
	gameResult.SpinResult = &betModel.SpinResult{}
	gameResult.TotalScore = 10000
	data, err := json.Marshal(gameResult)
	if err != nil {
		log.Fatal(err)
		return
	}

	param := save.Param{}
	param.Ctx = context.Background()
	param.GameMode = GameModeDemo
	param.PlayerID = 3345678
	param.GameResult = util.PointerString(string(data))
	param.Bet = util.PointerInt(1)
	param.Value = util.PointerInt(1000)
	param.Bonus = util.PointerBool(false)
	if err := repo.Save(param); err != nil {
		log.Fatalf("❌: %v", err)
	}
}

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
	repo := New(redisDB.RDB())

	gameResult := betModel.NewGameResult(SpinModeBase)
	gameResult.SpinResult = &betModel.SpinResult{}
	gameResult.TotalScore = 10000
	data, err := json.Marshal(gameResult)
	if err != nil {
		log.Fatal(err)
		return
	}
	if err := repo.SaveGameResult(context.Background(), GameModeDemo, 123, string(data)); err != nil {
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
	repo := New(redisDB.RDB())

	if err := repo.SaveBet(context.Background(), GameModeDemo, 123, 1); err != nil {
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
	repo := New(redisDB.RDB())

	if err := repo.SaveValue(context.Background(), GameModeDemo, 123, 1000); err != nil {
		log.Fatalf("❌: %v", err)
	}
}

func TestRepository_GetParam(t *testing.T) {
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
	repo := New(redisDB.RDB())

	output, err := repo.GetParam(context.Background(), GameModeDemo, 3345678)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(output.Bet)
	fmt.Println(output.Value)
	fmt.Println(output.Bonus)
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
	repo := New(redisDB.RDB())

	data, err := repo.GetGameResult(context.Background(), GameModeDemo, 123)
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
