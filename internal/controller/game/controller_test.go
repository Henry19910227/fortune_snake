package game

import (
	"context"
	"encoding/json"
	"flag"
	gameConfig "game_server_slots_fortune_snake/config/game"
	"game_server_slots_fortune_snake/config/system"
	"game_server_slots_fortune_snake/constants"
	"game_server_slots_fortune_snake/db"
	repoFactory "game_server_slots_fortune_snake/internal/factory/repository"
	serviceFactory "game_server_slots_fortune_snake/internal/factory/service"
	betModel "game_server_slots_fortune_snake/internal/model/controller/game/bet"
	playerModel "game_server_slots_fortune_snake/internal/model/entity/player"
	"game_server_slots_fortune_snake/internal/server"
	"log"
	"testing"
)

func GetController() *controller {
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

	// 7. 初始化 Redis
	redisDB, err := db.NewRedisDB(appConfig.Config().Redis)
	if err != nil {
		log.Fatalf("❌ Redis 初始化失败: %v", err)
	}

	// 初始化工廠
	repoFact := repoFactory.New(mysqlDB.DB(), redisDB.RDB(), gameConfig.New())
	serviceFact := serviceFactory.New(repoFact)

	gameSvc := serviceFact.GameService()
	gameDemoSvc := serviceFact.GameDemoService()
	weightSvc := serviceFact.WeightService()
	resultFreeSvc := serviceFact.ResultFreeService()
	resultLoader := serviceFact.ResultLoader()
	resultFreeLoader := serviceFact.ResultFreeLoader()
	reelsSvc := serviceFact.ReelsService()
	reelsFreeSvc := serviceFact.ReelsFreeService()
	settleSvc := serviceFact.SettleService()
	return &controller{gameSvc, gameDemoSvc, weightSvc, resultFreeSvc, resultLoader, resultFreeLoader, reelsSvc, reelsFreeSvc, settleSvc}
}

func TestGameController_Bet(t *testing.T) {
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

	// 創建 game controller
	gameController := GetController()
	gameController.Bet(&ctx)
}

func TestGameController_FreeModeInDemo(t *testing.T) {
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

	// 創建 game controller
	gameController := GetController()
	gameController.FreeModeInDemo(engineCtx, player, [][]int{{1, 1, 1}, {1, 1, 1, 1}, {1, 1, 1}})
}
