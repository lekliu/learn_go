package etcd

import (
	"context"
	"main/registry"
	"testing"
	"time"
)

func TestRegister(t *testing.T) {
	registryInst, err := registry.InitRegistry(context.TODO(), "etcd",
		registry.WithAddrs([]string{"127.0. 0. 1:2379"}),
		registry.WithTimeout(time.Second),
		registry.WithRegistryPath("/my/etcd/"),
		registry.WithHeartBeat(5),
	)

	if err != nil {
		t.Errorf("init registry failed, err:%v", err)
		return
	}
	service := &registry.Service{
		Name: "comment_service",
	}
	service.Nodes = append(
		service.Nodes, &registry.Node{
			Ip:   "127.0. 0.1",
			Port: 8801,
		},
		&registry.Node{
			Ip:   "127.0. 0.2",
			Port: 8801,
		},
	)

	registryInst.Register(context.TODO(), service)
	go func() {
		time.Sleep(time.Second * 10)
		service.Nodes = append(service.Nodes, &registry.Node{
			Ip:   "127.0. 0.3",
			Port: 8801,
		},
			&registry.Node{
				Ip:   "127.0. 0.4",
				Port: 8801,
			},
		)
		registryInst.Register(context.TODO(), service)
	}()
	//TODO

}
