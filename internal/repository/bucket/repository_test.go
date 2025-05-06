package result

import (
	"flag"
	"game_server_slots_fortune_snake/config"
	"game_server_slots_fortune_snake/internal/model/entity/result"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestResultRepo_Save(t *testing.T) {
	// 創建配置檔組件
	configFile := flag.String("config", "config.yaml", "YAML configuration file name")
	flag.Parse()
	cfg := config.New(configFile)
	resultRepo := New(cfg.BaseBucketConfig())

	item := &result.Item{
		ID:      1,
		Rate:    1,
		Symbols: "",
	}

	// 將盤面賠率儲存至桶內
	assert.Equal(t, 7708, resultRepo.Quota())
	resultRepo.Save(item)
	assert.Equal(t, 7707, resultRepo.Quota())
	bucket := resultRepo.Items(0.9, 1.0)
	assert.Equal(t, 1, len(bucket))
}

func TestResultRepo_Save_2(t *testing.T) {
	// 創建配置檔組件
	configFile := flag.String("config", "config.yaml", "YAML configuration file name")
	flag.Parse()
	cfg := config.New(configFile)

	resultRepo := New(cfg.FreeBucketConfig())

	for resultRepo.Quota() > 0 {
		randomRate := GetFreeRandomRate()
		item := &result.Item{
			ID:      1,
			Rate:    randomRate,
			Symbols: "",
		}
		resultRepo.Save(item)
	}
	assert.Equal(t, 0, resultRepo.Quota())
	assert.Equal(t, 100, len(resultRepo.Items(25, 30)))
	assert.Equal(t, 2, len(resultRepo.Items(1000, 2450)))
}

func GetRandomRate() float64 {
	step := 0.1
	max := 2000.0
	steps := int(max/step) + 1 // 加 1 是因為包含 0.0

	index := rand.Intn(steps)            // 隨機選擇 0 ~ steps-1
	value := (float64(index) - 1) * step // 計算對應的值
	return value
}

func GetFreeRandomRate() float64 {
	steps := int(2500) + 1    // 加 1 是因為包含 0
	index := rand.Intn(steps) // 隨機選擇 0 ~ steps-1
	value := float64(index)   // 計算對應的值
	return value
}
