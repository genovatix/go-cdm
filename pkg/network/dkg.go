package network

import (
	"go.etcd.io/etcd/client/v3"
)

type DKGManager struct {
	EtcdClient *clientv3.Client
}
