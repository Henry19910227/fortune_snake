package load

import (
	"flag"
	gameCfg "game_server_slots_fortune_snake/config/game"
	"game_server_slots_fortune_snake/config/system"
	"game_server_slots_fortune_snake/db"
	bucketRepository "game_server_slots_fortune_snake/internal/repository/bucket"
	bucketFreeRepository "game_server_slots_fortune_snake/internal/repository/bucket_free"
	resultRepository "game_server_slots_fortune_snake/internal/repository/result"
	resultFreeRepository "game_server_slots_fortune_snake/internal/repository/result_free"
	weightRepository "game_server_slots_fortune_snake/internal/repository/weight"
	resultService "game_server_slots_fortune_snake/internal/service/result"
	resultFreeService "game_server_slots_fortune_snake/internal/service/result_free"
	weightService "game_server_slots_fortune_snake/internal/service/weight"
	"log"
	"testing"
)

func TestLoadController_Load(t *testing.T) {
	// 加載 yaml
	configFile := flag.String("config", "config.yaml", "YAML configuration file name")
	flag.Parse()
	// 加載配置文件
	cfg := gameCfg.New()
	appConfig := system.New(configFile)
	// 初始化 MySQL DB
	mysqlDB, err := db.NewMysqlDB(appConfig.Config().Database)
	if err != nil {
		log.Fatalf("❌ MySQL 初始化失败: %v", err)
	}
	defer mysqlDB.Close()
	// 初始化 repository
	weightRepo := weightRepository.New(cfg.RTPConfig())
	resultRepo := resultRepository.New(mysqlDB.DB())
	resultFreeRepo := resultFreeRepository.New(mysqlDB.DB())
	bucketRepo := bucketRepository.New(cfg.BaseBucketConfig())
	bucketFreeRepo := bucketFreeRepository.New(cfg.FreeBucketConfig())
	// 初始化 service
	weightSvc := weightService.New(weightRepo)
	resultSvc := resultService.New(resultRepo, bucketRepo)
	resultFreeSvc := resultFreeService.New(resultFreeRepo, bucketFreeRepo)
	// 初始化 controller
	loadController := New(weightSvc, resultSvc, resultFreeSvc)
	loadController.Load()
}
