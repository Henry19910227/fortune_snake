package result

import (
	"flag"
	"fmt"
	gameConfig "game_server_slots_fortune_snake/config/game"
	"game_server_slots_fortune_snake/config/system"
	"game_server_slots_fortune_snake/constants"
	"game_server_slots_fortune_snake/db"
	bucketRepository "game_server_slots_fortune_snake/internal/repository/bucket"
	resultLoaderRepository "game_server_slots_fortune_snake/internal/repository/result_loader"
	settleRepository "game_server_slots_fortune_snake/internal/repository/settle"
	symbolRepository "game_server_slots_fortune_snake/internal/repository/symbol"
	reelService "game_server_slots_fortune_snake/internal/service/reels"
	resultLoaderService "game_server_slots_fortune_snake/internal/service/result_loader"
	settleService "game_server_slots_fortune_snake/internal/service/settle"
	"log"
	"testing"
)

func TestResultController_Generate(t *testing.T) {
	// 加載 yaml
	configFile := flag.String("config", "config.yaml", "YAML configuration file name")
	flag.Parse()

	// 加载配置文件
	appConfig := system.New(configFile)

	// 加載game配置文件
	gameCfg := gameConfig.New()

	// 初始化 MySQL DB
	mysqlDB, err := db.NewMysqlDB(appConfig.Config().Database)
	if err != nil {
		log.Fatalf("❌ MySQL 初始化失败: %v", err)
	}
	defer mysqlDB.Close()

	// repo 建立
	symRepo := symbolRepository.New(gameCfg.SymbolConfig())
	settRepo := settleRepository.New()
	resultRepo := resultLoaderRepository.SpinMode(constants.SpinModeNormal).Init(mysqlDB.DB())
	bucketRepo := bucketRepository.New(gameCfg.BaseBucketConfig())

	// service 建立
	reelSvc := reelService.New(symRepo)
	settleSvc := settleService.New(settRepo)
	resultSvc := resultLoaderService.SpinMode(constants.SpinModeNormal).Init(resultRepo, bucketRepo)

	// controller 建立
	resultController := New(reelSvc, settleSvc, resultSvc)
	// 生產盤面數據
	resultController.Generate()
	// 檢查數據
	results := bucketRepo.Items(200, 500)
	for _, result := range results {
		fmt.Printf("rate: %v, symbols: %v\n", result.Rate, result.Symbols)
	}
}
