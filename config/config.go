package config

import (
	"fmt"
	"game_server_slots_fortune_snake/internal/model"
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

type config struct {
	config *model.Config
}

func New(configFile *string) Config {
	// 检查配置文件参数是否为空
	if *configFile == "" {
		fmt.Println("❌ The configuration file name must be included")
		os.Exit(1)
	}
	data, err := os.ReadFile(*configFile)
	if err != nil {
		log.Fatalf("❌ 读取配置文件失败: %v", err)
	}

	var cfg model.Config
	// 将 YAML 文件反序列化到 cfg 变量中
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		log.Fatalf("❌ 解析配置文件失败: %v", err)
	}
	log.Println("✅ 配置加载成功")
	return &config{config: &cfg}
}

func (c *config) Config() *model.Config {
	return c.config
}

func (c *config) BaseResultConfig() []*model.ResultConfig {
	return []*model.ResultConfig{
		{LowerLimit: 0, UpperLimit: 0, MaxCapacity: 5000},
		{LowerLimit: 0, UpperLimit: 0.2, MaxCapacity: 0},
		{LowerLimit: 0.2, UpperLimit: 0.3, MaxCapacity: 0},
		{LowerLimit: 0.3, UpperLimit: 0.4, MaxCapacity: 0},
		{LowerLimit: 0.4, UpperLimit: 0.5, MaxCapacity: 0},
		{LowerLimit: 0.5, UpperLimit: 0.6, MaxCapacity: 200},
		{LowerLimit: 0.6, UpperLimit: 0.7, MaxCapacity: 0},
		{LowerLimit: 0.7, UpperLimit: 0.8, MaxCapacity: 0},
		{LowerLimit: 0.8, UpperLimit: 0.9, MaxCapacity: 0},
		{LowerLimit: 0.9, UpperLimit: 1, MaxCapacity: 200},
		{LowerLimit: 1, UpperLimit: 2, MaxCapacity: 200},
		{LowerLimit: 2, UpperLimit: 3, MaxCapacity: 200},
		{LowerLimit: 3, UpperLimit: 4, MaxCapacity: 200},
		{LowerLimit: 4, UpperLimit: 5, MaxCapacity: 200},
		{LowerLimit: 5, UpperLimit: 6, MaxCapacity: 200},
		{LowerLimit: 6, UpperLimit: 7, MaxCapacity: 200},
		{LowerLimit: 7, UpperLimit: 8, MaxCapacity: 200},
		{LowerLimit: 8, UpperLimit: 9, MaxCapacity: 200},
		{LowerLimit: 9, UpperLimit: 10, MaxCapacity: 200},
		{LowerLimit: 10, UpperLimit: 15, MaxCapacity: 150},
		{LowerLimit: 15, UpperLimit: 20, MaxCapacity: 120},
		{LowerLimit: 20, UpperLimit: 30, MaxCapacity: 100},
		{LowerLimit: 30, UpperLimit: 40, MaxCapacity: 90},
		{LowerLimit: 40, UpperLimit: 50, MaxCapacity: 20},
		{LowerLimit: 50, UpperLimit: 100, MaxCapacity: 10},
		{LowerLimit: 100, UpperLimit: 200, MaxCapacity: 10},
		{LowerLimit: 200, UpperLimit: 500, MaxCapacity: 5},
		{LowerLimit: 500, UpperLimit: 1000, MaxCapacity: 2},
		{LowerLimit: 1000, UpperLimit: 2500, MaxCapacity: 1},
	}
}
