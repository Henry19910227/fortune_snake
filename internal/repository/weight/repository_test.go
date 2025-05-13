package weight

import (
	"fmt"
	gameCfg "game_server_slots_fortune_snake/config/game"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRepository_LoadBaseWeight(t *testing.T) {
	cfg := gameCfg.New()
	repo := New(cfg.RTPConfig())
	repo.LoadBaseWeight()
	for _, stat := range repo.BaseWeight() {
		fmt.Println(stat)
	}
	assert.Equal(t, 148, len(repo.BaseWeight()))
}

func TestRepository_LoadFreeWeight(t *testing.T) {
	cfg := gameCfg.New()
	repo := New(cfg.RTPConfig())
	repo.LoadFreeWeight()
	for _, stat := range repo.FreeWeight() {
		fmt.Println(stat)
	}
	assert.Equal(t, 224, len(repo.FreeWeight()))
}

func TestRepository_LoadBaseWeightH(t *testing.T) {
	cfg := gameCfg.New()
	repo := New(cfg.RTPConfig())
	repo.LoadBaseWeightH()
	for _, stat := range repo.BaseWeightH() {
		fmt.Println(stat)
	}
	assert.Equal(t, 148, len(repo.BaseWeightH()))
}

func TestRepository_LoadFreeWeightH(t *testing.T) {
	cfg := gameCfg.New()
	repo := New(cfg.RTPConfig())
	repo.LoadFreeWeightH()
	for _, stat := range repo.FreeWeightH() {
		fmt.Println(stat)
	}
	assert.Equal(t, 224, len(repo.FreeWeightH()))
}

func TestRepository_RandomBaseWeightRate(t *testing.T) {
	cfg := gameCfg.New()
	repo := New(cfg.RTPConfig())
	repo.LoadBaseWeight()
	rate, err := repo.RandomBaseWeightRate(96)
	assert.Nil(t, err)
	fmt.Println(rate)
}

func TestRepository_RandomBaseWeightHRate(t *testing.T) {
	cfg := gameCfg.New()
	repo := New(cfg.RTPConfig())
	repo.LoadBaseWeightH()
	rate, err := repo.RandomBaseWeightHRate(96)
	assert.Nil(t, err)
	fmt.Println(rate)
}

func TestRepository_RandomFreeWeightRate(t *testing.T) {
	cfg := gameCfg.New()
	repo := New(cfg.RTPConfig())
	repo.LoadFreeWeight()
	rate, err := repo.RandomFreeWeightRate(96)
	assert.Nil(t, err)
	fmt.Println(rate)
}

func TestRepository_RandomFreeWeightHRate(t *testing.T) {
	cfg := gameCfg.New()
	repo := New(cfg.RTPConfig())
	repo.LoadFreeWeightH()
	rate, err := repo.RandomFreeWeightHRate(96)
	assert.Nil(t, err)
	fmt.Println(rate)
}
