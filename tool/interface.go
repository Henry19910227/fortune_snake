package tool

import (
	"github.com/elastic/go-elasticsearch/v8"
	"go.uber.org/zap"
)

type Logger interface {
	Logs(level string, writeToElk bool, msg string, fields ...zap.Field)
	SetELK(elkClient *elasticsearch.Client)
	Close()
}
