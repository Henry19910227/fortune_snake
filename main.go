package main

import (
	"flag"
	"game_server_slots_fortune_snake/config"
	"game_server_slots_fortune_snake/pkg"
	"game_server_slots_fortune_snake/utils"
)

func main() {
	// 0. 加載 yaml
	configFile := flag.String("config", "config.yaml", "YAML configuration file name")
	flag.Parse()

	// 1. 加载配置文件
	appConfig := config.New(configFile)

	// 2. 初始化语言包
	_ = pkg.InitLocalize(appConfig.Config().Language.Default)

	// 3. 初始化分布式雪花 ID
	_, _ = utils.NewSnowFlake(appConfig.Config().Server.ServerNode)

}
