package result_free

import (
	"flag"
	"game_server_slots_fortune_snake/config/system"
	"game_server_slots_fortune_snake/db"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

// TestRepository_LoadData 測試載入盤面資料至內存場景
func TestRepository_LoadData(t *testing.T) {
	// 加載 yaml
	configFile := flag.String("config", "config.yaml", "YAML configuration file name")
	flag.Parse()

	// 加载配置文件
	appConfig := system.New(configFile)

	// 初始化 MySQL DB
	mysqlDB, err := db.NewMysqlDB(appConfig.Config().Database)
	if err != nil {
		log.Fatalf("❌ MySQL 初始化失败: %v", err)
	}
	defer mysqlDB.Close()

	// 創建 repo
	repo := New(mysqlDB.DB())
	// 插入數據
	err = repo.LoadData()
	assert.Nil(t, err)
	// 獲取某個rate對應的隨機盤面
	item, err := repo.Random(1038)
	assert.Nil(t, err)
	assert.Equal(t, "[[5, 0, 6], [0, 0, 0, 0], [3, 0, 4]]", item.Symbols)
}
