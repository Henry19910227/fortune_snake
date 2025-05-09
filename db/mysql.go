package db

import (
	"fmt"
	model "game_server_slots_fortune_snake/internal/model/config/system"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

type mysqlDB struct {
	db *gorm.DB
}

func NewMysqlDB(config model.DatabaseConfig) (MysqlDB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.Database,
		config.Charset,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	// 设置连接池
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxOpenConns(config.MaxOpenConns)
	sqlDB.SetMaxIdleConns(config.MaxIdleConns)
	sqlDB.SetConnMaxLifetime(time.Hour)
	log.Println("✅ MySQL 连接成功")
	return &mysqlDB{db: db}, nil
}

func (m *mysqlDB) DB() *gorm.DB {
	return m.db
}

func (m *mysqlDB) Close() {
	if m.db == nil {
		return
	}
	sqlDB, err := m.db.DB()
	if err == nil {
		_ = sqlDB.Close()
		log.Println("🛑 MySQL 连接已关闭")
	}
}
