package result

import (
	"flag"
	"fmt"
	gameConfig "game_server_slots_fortune_snake/config/game"
	"game_server_slots_fortune_snake/config/system"
	"game_server_slots_fortune_snake/db"
	bucketRepository "game_server_slots_fortune_snake/internal/repository/bucket"
	resultRepository "game_server_slots_fortune_snake/internal/repository/result"
	settleRepository "game_server_slots_fortune_snake/internal/repository/settle"
	symbolRepository "game_server_slots_fortune_snake/internal/repository/symbol"
	reelService "game_server_slots_fortune_snake/internal/service/reels"
	reelsFreeService "game_server_slots_fortune_snake/internal/service/reels_free"
	resultService "game_server_slots_fortune_snake/internal/service/result"
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
	resultRepo := resultRepository.New(mysqlDB.DB())
	bucketRepo := bucketRepository.New(gameCfg.BaseBucketConfig())

	// service 建立
	reelSvc := reelService.New(symRepo)
	reelFreeSvc := reelsFreeService.New(symRepo)
	settleSvc := settleService.New(settRepo)
	resultSvc := resultService.New(resultRepo, bucketRepo)

	// controller 建立
	resultController := New(reelSvc, reelFreeSvc, settleSvc, resultSvc)
	// 生產盤面數據
	resultController.Generate()
	// 檢查數據
	results := bucketRepo.Items(200, 500)
	for _, result := range results {
		fmt.Printf("rate: %v, symbols: %v\n", result.Rate, result.Symbols)
	}
}

func TestResultController_Generate_Free(t *testing.T) {
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
	resultRepo := resultRepository.New(mysqlDB.DB())
	bucketRepo := bucketRepository.New(gameCfg.BaseBucketConfig())

	// service 建立
	reelSvc := reelService.New(symRepo)
	reelFreeSvc := reelsFreeService.New(symRepo)
	settleSvc := settleService.New(settRepo)
	resultSvc := resultService.New(resultRepo, bucketRepo)

	// controller 建立
	resultController := New(reelSvc, reelFreeSvc, settleSvc, resultSvc)
	// 生產盤面數據
	resultController.GenerateFree()
	// 檢查數據
	//results := bucketRepo.Items(200, 500)
	//for _, result := range results {
	//	fmt.Printf("rate: %v, symbols: %v\n", result.Rate, result.Symbols)
	//}
}
