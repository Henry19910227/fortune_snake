package utils

import (
	"fmt"
	"time"

	"github.com/bwmarrin/snowflake"
)

type snowFlake struct {
	node *snowflake.Node
}

func NewSnowFlake(workerID int64) (SnowFlake, error) {
	// 设置起始时间（自定义一个固定时间点）
	snowflake.Epoch = time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC).UnixNano() / 1e6

	// 初始化 Snowflake 节点
	node, err := snowflake.NewNode(workerID)
	if err != nil {
		fmt.Println("初始化 Snowflake 失败:", err)
		return nil, err
	}
	return &snowFlake{node: node}, nil
}

func (sf *snowFlake) GenerateID() (uint64, error) {
	if sf.node == nil {
		return 0, fmt.Errorf("snowflake 未初始化，请先调用 NewSnowFlake()")
	}
	return uint64(sf.node.Generate().Int64()), nil
}
