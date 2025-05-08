package result

import (
	"flag"
	"fmt"
	"game_server_slots_fortune_snake/config"
	"game_server_slots_fortune_snake/db"
	bucketRepository "game_server_slots_fortune_snake/internal/repository/bucket"
	resultRepository "game_server_slots_fortune_snake/internal/repository/result"
	settleRepository "game_server_slots_fortune_snake/internal/repository/settle"
	symbolRepository "game_server_slots_fortune_snake/internal/repository/symbol"
	reelService "game_server_slots_fortune_snake/internal/service/reels"
	resultService "game_server_slots_fortune_snake/internal/service/result"
	settleService "game_server_slots_fortune_snake/internal/service/settle"
	"log"
	"testing"
)

func TestResultController_Generate(t *testing.T) {
	// 加載 yaml
	configFile := flag.String("config", "config.yaml", "YAML configuration file name")
	flag.Parse()
	cfg := config.New(configFile)

	// 加载配置文件
	appConfig := config.New(configFile)

	// 初始化 MySQL DB
	mysqlDB, err := db.NewMysqlDB(appConfig.Config().Database)
	if err != nil {
		log.Fatalf("❌ MySQL 初始化失败: %v", err)
	}
	defer mysqlDB.Close()

	symRepo := symbolRepository.New()
	settRepo := settleRepository.New()
	resultRepo := resultRepository.New(mysqlDB.DB())
	bucketRepo := bucketRepository.New(cfg.BaseBucketConfig())

	reelSvc := reelService.New(symRepo)
	settleSvc := settleService.New(settRepo)
	resultSvc := resultService.New(resultRepo, bucketRepo)

	con := New(reelSvc, settleSvc, resultSvc)
	con.Generate()

	results := bucketRepo.Items(200, 500)
	for _, result := range results {
		fmt.Printf("rate: %v, symbols: %v\n", result.Rate, result.Symbols)
	}
}
