package etcd

import (
	"context"
	"fmt"
	"game_server_slots_fortune_snake/internal/model"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log"
	"time"
)

var EtcdClient *clientv3.Client
var leaseID clientv3.LeaseID
var keepAliveCh <-chan *clientv3.LeaseKeepAliveResponse

// InitEtcd 连接 Etcd
func InitEtcd(config model.EtcdConfig) error {
	endpoints := config.Endpoints
	var err error
	for i := 1; i <= 3; i++ { // 3 次重试
		EtcdClient, err = clientv3.New(clientv3.Config{
			Endpoints:   endpoints,
			DialTimeout: time.Duration(config.Ttl) * time.Second,
		})
		if err == nil {
			log.Println("✅ Etcd 连接成功")
			return nil
		}
		log.Printf("⚠️ Etcd 连接失败（重试 %d/3）：%v", i, err)
		time.Sleep(time.Duration(config.Ttl) * time.Second)
	}
	return fmt.Errorf("❌ Etcd 连接失败，所有重试均已失败")
}

// RegisterService 注册 小火箭 服务器到 Etcd
func RegisterService(config model.Config) error {
	if EtcdClient == nil {
		return fmt.Errorf("etcd 未初始化")
	}
	key := config.Etcd.ServicePrefix
	value := fmt.Sprintf("%s:%d", config.Server.Ip, config.Server.Port)
	leaseResp, err := EtcdClient.Grant(context.Background(), int64(config.Etcd.Ttl))
	if err != nil {
		return fmt.Errorf("创建租约失败: %v", err)
	}
	leaseID = leaseResp.ID
	_, err = EtcdClient.Put(context.Background(), key, value, clientv3.WithLease(leaseID))
	if err != nil {
		return fmt.Errorf("注册服务失败: %v", err)
	}
	log.Printf("✅ 小火箭 服务已注册到 Etcd: %s -> %s", key, value)
	keepAliveCh, err = EtcdClient.KeepAlive(context.Background(), leaseID)
	if err != nil {
		return fmt.Errorf("续租失败: %v", err)
	}
	go func() {
		for kaResp := range keepAliveCh {
			if kaResp == nil {
				log.Println("⚠️ Etcd 续租失败，服务可能被删除")
				break
			}
		}
	}()
	return nil
}

// UnregisterService 取消服务注册
func UnregisterService(config model.EtcdConfig) {
	if EtcdClient == nil {
		return
	}
	key := config.ServicePrefix
	_, err := EtcdClient.Delete(context.Background(), key)
	if err != nil {
		log.Printf("❌ 取消注册失败: %v", err)
	} else {
		log.Println("🛑 服务已从 Etcd 注销")
	}
}
