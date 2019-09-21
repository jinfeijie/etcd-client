package etcd

import (
	"context"
	"github.com/etcd-io/etcd/clientv3"
	"time"
)

var (
	endpoints []string
	Cli    *clientv3.Client
	Cancel context.CancelFunc
)

type EtcdCnf struct {
	Endpoints []string
}

var Etcd etcd
type etcd struct {}

func NewEtcd() clientv3.Client {
	return Etcd.Open()
}

func SetEtcdCnf(cnf EtcdCnf) {
	endpoints = cnf.Endpoints
	if endpoints[0] == "" {
		// use 127.0.0.1:2379 by default
		endpoints = append(endpoints, "127.0.0.1:2379")
	}
}

func (etcd *etcd) Open() clientv3.Client {
	Cli, _ = clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: 5 * time.Second,
	})

	_, Cancel = context.WithTimeout(context.Background(), 5*time.Second)
	return *Cli
}

func Close() {
	defer Cli.Close()
	defer Cancel()
}
