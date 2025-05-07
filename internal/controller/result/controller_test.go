package result

import (
	"flag"
	"fmt"
	"game_server_slots_fortune_snake/config"
	bucketRepository "game_server_slots_fortune_snake/internal/repository/bucket"
	resultRepository "game_server_slots_fortune_snake/internal/repository/result"
	settleRepository "game_server_slots_fortune_snake/internal/repository/settle"
	symbolRepository "game_server_slots_fortune_snake/internal/repository/symbol"
	reelService "game_server_slots_fortune_snake/internal/service/reels"
	resultService "game_server_slots_fortune_snake/internal/service/result"
	settleService "game_server_slots_fortune_snake/internal/service/settle"
	"testing"
)

func TestResultController_Generate(t *testing.T) {
	configFile := flag.String("config", "config.yaml", "YAML configuration file name")
	flag.Parse()
	cfg := config.New(configFile)

	symRepo := symbolRepository.New()
	settRepo := settleRepository.New()
	resultRepo := resultRepository.New(nil)
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
