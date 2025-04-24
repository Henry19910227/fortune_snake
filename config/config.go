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
