package snow_flake

import (
	"github.com/bwmarrin/snowflake"
	"time"
)

type repository struct {
	node *snowflake.Node
}

func New(workerID int64) (Repository, error) {
	// 设置起始时间（自定义一个固定时间点）
	snowflake.Epoch = time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC).UnixNano() / 1e6

	// 初始化 Snowflake 节点
	node, err := snowflake.NewNode(workerID)
	if err != nil {
		return nil, err
	}
	return &repository{node: node}, nil
}

func (r *repository) GenerateID() uint64 {
	if r.node == nil {
		return 0
	}
	return uint64(r.node.Generate().Int64())
}
