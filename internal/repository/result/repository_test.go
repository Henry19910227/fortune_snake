package result

import (
	"flag"
	"game_server_slots_fortune_snake/config"
	"game_server_slots_fortune_snake/db"
	"game_server_slots_fortune_snake/internal/model/entity/result"
	"game_server_slots_fortune_snake/internal/model/repository/result/create_items"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"time"
)

// TestRepository_CreatItems 測試插入多筆數據場景
// rate: 203, symbols: [[6,0,1],[0,0,0,0],[1,2,6]]
// rate: 212, symbols: [[5,6,1],[0,0,0,0],[6,6,0]]
// rate: 212, symbols: [[6,1,6],[0,0,0,0],[6,0,3]]
// rate: 209, symbols: [[1,2,6],[0,0,0,0],[3,0,6]]
// rate: 260, symbols: [[1,4,4],[0,0,0,0],[1,4,4]]
func TestRepository_CreatItems(t *testing.T) {
	// 加載 yaml
	configFile := flag.String("config", "config.yaml", "YAML configuration file name")
	flag.Parse()

	// 加载配置文件
	appConfig := config.New(configFile)

	// 初始化 MySQL DB
	mysqlDB, err := db.NewMysqlDB(appConfig.Config().Database)
	if err != nil {
		log.Fatalf("❌ MySQL 初始化失败: %v", err)
	}
	defer mysqlDB.Close()

	// 準備數據
	items := make([]*result.Item, 0)
	items = append(items, &result.Item{Rate: 203, Symbols: "[[6,0,1],[0,0,0,0],[1,2,6]]", CreatedAt: time.Now().UnixMilli(), UpdatedAt: time.Now().UnixMilli()})
	items = append(items, &result.Item{Rate: 212, Symbols: "[[5,6,1],[0,0,0,0],[6,6,0]]", CreatedAt: time.Now().UnixMilli(), UpdatedAt: time.Now().UnixMilli()})
	items = append(items, &result.Item{Rate: 212, Symbols: "[[6,1,6],[0,0,0,0],[6,0,3]]", CreatedAt: time.Now().UnixMilli(), UpdatedAt: time.Now().UnixMilli()})
	items = append(items, &result.Item{Rate: 209, Symbols: "[[1,2,6],[0,0,0,0],[3,0,6]]", CreatedAt: time.Now().UnixMilli(), UpdatedAt: time.Now().UnixMilli()})
	items = append(items, &result.Item{Rate: 260, Symbols: "[[1,4,4],[0,0,0,0],[1,4,4]]", CreatedAt: time.Now().UnixMilli(), UpdatedAt: time.Now().UnixMilli()})

	// 創建 repo
	repo := New(mysqlDB.DB())

	// 插入數據
	err = repo.CreateItems(create_items.NewInput(items))
	assert.Nil(t, err)
}
