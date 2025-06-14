package game

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	gameConfig "game_server_slots_fortune_snake/config/game"
	"game_server_slots_fortune_snake/config/system"
	"game_server_slots_fortune_snake/constants"
	"game_server_slots_fortune_snake/db"
	"game_server_slots_fortune_snake/internal/controller/load"
	repoFactory "game_server_slots_fortune_snake/internal/factory/repository"
	serviceFactory "game_server_slots_fortune_snake/internal/factory/service"
	"game_server_slots_fortune_snake/internal/model"
	betModel "game_server_slots_fortune_snake/internal/model/controller/game/bet"
	playerModel "game_server_slots_fortune_snake/internal/model/entity/player"
	"game_server_slots_fortune_snake/internal/server"
	"log"
	"testing"
)

func TestGameController_Bet(t *testing.T) {
	// 加載 yaml
	configFile := flag.String("config", "config.yaml", "YAML configuration file name")
	flag.Parse()

	// 加载配置文件
	appConfig := system.New(configFile)

	// 初始化 MySQL DB
	mysqlDB, err := db.NewMysqlDB(appConfig.Config().Database)
	if err != nil {
		log.Fatalf("❌ MySQL 初始化失败: %v", err)
	}
	defer mysqlDB.Close()

	// 初始化 Redis
	redisDB, err := db.NewRedisDB(appConfig.Config().Redis)
	if err != nil {
		log.Fatalf("❌ Redis 初始化失败: %v", err)
	}
	defer redisDB.Close()

	// 初始化工廠
	repoFact, _ := repoFactory.New(mysqlDB.DB(), redisDB.RDB(), gameConfig.New(), appConfig.Config())
	serviceFact := serviceFactory.New(repoFact)

	// 初始化 service 模塊
	gameSvc := serviceFact.GameService()
	weightSvc := serviceFact.WeightService()
	resultFreeSvc := serviceFact.ResultFreeService()
	resultLoader := serviceFact.ResultLoader()
	resultFreeLoader := serviceFact.ResultFreeLoader()
	reelsSvc := serviceFact.ReelsService()
	reelsFreeSvc := serviceFact.ReelsFreeService()
	settleSvc := serviceFact.SettleService()

	// 載入權重、盤面
	loadController := load.New(weightSvc, resultLoader, resultFreeLoader)
	loadController.Load()

	// 初始化 game controller
	gameController := &controller{gameSvc, weightSvc, resultFreeSvc, resultLoader, resultFreeLoader, reelsSvc, reelsFreeSvc, settleSvc}

	// 創建 param
	m := betModel.Param{Bet: 1, Value: 1000}
	b, _ := json.Marshal(m)

	// 創建 player
	player := &playerModel.Session{}
	player.PlayerId = 123
	player.Mode = constants.GameModeDemo
	player.GameRtp = 96

	// 創建 context
	ctx := server.NewContext(nil)
	ctx.SetData(b)
	ctx.Set("session", player)
	ctx.Set("ctx", context.Background())

	// 執行 bet
	gameController.Bet(&ctx)

	resModel := &model.MessageResponse{}
	_ = json.Unmarshal(ctx.Result(), resModel)
	fmt.Println(resModel.Data)
}

func TestGameController_FreeModeInDemo(t *testing.T) {
	// 加載 yaml
	configFile := flag.String("config", "config.yaml", "YAML configuration file name")
	flag.Parse()

	// 加载配置文件
	appConfig := system.New(configFile)

	// 初始化 MySQL DB
	mysqlDB, err := db.NewMysqlDB(appConfig.Config().Database)
	if err != nil {
		log.Fatalf("❌ MySQL 初始化失败: %v", err)
	}
	defer mysqlDB.Close()

	// 初始化 Redis
	redisDB, err := db.NewRedisDB(appConfig.Config().Redis)
	if err != nil {
		log.Fatalf("❌ Redis 初始化失败: %v", err)
	}
	defer redisDB.Close()

	// 初始化工廠
	repoFact, _ := repoFactory.New(mysqlDB.DB(), redisDB.RDB(), gameConfig.New(), appConfig.Config())
	serviceFact := serviceFactory.New(repoFact)

	// 初始化 game controller
	gameSvc := serviceFact.GameService()
	weightSvc := serviceFact.WeightService()
	resultFreeSvc := serviceFact.ResultFreeService()
	resultLoader := serviceFact.ResultLoader()
	resultFreeLoader := serviceFact.ResultFreeLoader()
	reelsSvc := serviceFact.ReelsService()
	reelsFreeSvc := serviceFact.ReelsFreeService()
	settleSvc := serviceFact.SettleService()
	gameController := &controller{gameSvc, weightSvc, resultFreeSvc, resultLoader, resultFreeLoader, reelsSvc, reelsFreeSvc, settleSvc}

	// 創建 param
	m := betModel.Param{Bet: 1, Value: 1000}
	b, _ := json.Marshal(m)

	// 創建 player
	player := &playerModel.Session{}
	player.PlayerId = 123
	player.Mode = constants.GameModeDemo
	player.GameRtp = 96

	// 創建 context
	engineCtx := &server.Context{}
	engineCtx.SetData(b)

	// 執行 FreeModeInDemo
	gameController.FreeModeInDemo(engineCtx, [][]int{{1, 1, 1}, {1, 1, 1, 1}, {1, 1, 1}}, false)
}
