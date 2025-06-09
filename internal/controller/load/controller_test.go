package load

import (
	"flag"
	gameCfg "game_server_slots_fortune_snake/config/game"
	"game_server_slots_fortune_snake/config/system"
	. "game_server_slots_fortune_snake/constants"
	"game_server_slots_fortune_snake/db"
	bucketRepository "game_server_slots_fortune_snake/internal/repository/bucket"
	resultLoaderRepository "game_server_slots_fortune_snake/internal/repository/result_loader"
	weightRepository "game_server_slots_fortune_snake/internal/repository/weight"
	resultLoaderService "game_server_slots_fortune_snake/internal/service/result_loader"
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
	resultLoaderRepo := resultLoaderRepository.New(mysqlDB.DB(), SpinModeBase)
	resultLoaderFreeRepo := resultLoaderRepository.New(mysqlDB.DB(), SpinModeFree)
	bucketRepo := bucketRepository.New(cfg.BaseBucketConfig())
	bucketFreeRepo := bucketRepository.New(cfg.FreeBucketConfig())
	// 初始化 service
	weightSvc := weightService.New(weightRepo)
	resultLoaderSvc := resultLoaderService.New(resultLoaderRepo, bucketRepo, SpinModeBase)
	resultFreeLoaderSvc := resultLoaderService.New(resultLoaderFreeRepo, bucketFreeRepo, SpinModeFree)
	// 初始化 controller
	loadController := New(weightSvc, resultLoaderSvc, resultFreeLoaderSvc)
	loadController.Load()
}
