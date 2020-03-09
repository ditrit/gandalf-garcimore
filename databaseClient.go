package main

import (
	"context"
	"fmt"
	"shoset/net"
	"time"

	"github.com/canonical/go-dqlite/client"
)

type DatabaseClient struct {
	DatabaseClientCluster []string
}

func NewDatabaseClient(cluster []string) (databaseClient *DatabaseClient) {
	databaseClient = new(DatabaseClient)
	//databaseClient.databaseClientCluster = cluster
	//databaseClient.DatabaseClientCluster = []string{"127.0.0.1:9000", "127.0.0.1:9001", "127.0.0.1:9002"}
	return
}

func NewDatabaseClient2() (databaseClient *DatabaseClient) {
	databaseClient = new(DatabaseClient)
	//databaseClient.databaseClientCluster = cluster
	databaseClient.DatabaseClientCluster = []string{"127.0.0.1:9000", "127.0.0.1:9001", "127.0.0.1:9002"}
	return
}

func NewDatabaseClient3() (databaseClient *DatabaseClient) {
	databaseClient = new(DatabaseClient)
	//databaseClient.databaseClientCluster = cluster
	//databaseClient.databaseClientCluster = []string{"127.0.0.1:9000", "127.0.0.1:9001", "127.0.0.1:9002"}
	return
}

func (dc DatabaseClient) GetLeader() (*client.Client, error) {
	store := dc.GetStore2()
	//store := dc.GetStore()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	return client.FindLeader(ctx, store, client.WithLogFunc(logFunc))
}

func (dc DatabaseClient) GetStore() client.NodeStore {

	store := client.NewInmemNodeStore()
	if len(dc.DatabaseClientCluster) == 0 {
	}
	infos := make([]client.NodeInfo, len(dc.DatabaseClientCluster))
	for i, address := range dc.DatabaseClientCluster {
		infos[i].ID, _ = net.IP2ID(address)
		infos[i].Address = address
	}
	fmt.Println("INFOS")
	fmt.Println(infos)
	store.Set(context.Background(), infos)
	return store
}

func (dc DatabaseClient) GetStore2() client.NodeStore {

	store := client.NewInmemNodeStore()
	if len(dc.DatabaseClientCluster) == 0 {
	}
	infos := make([]client.NodeInfo, len(dc.DatabaseClientCluster))
	for i, address := range dc.DatabaseClientCluster {
		//infos[i].ID, _ = net.IP2ID(address)
		infos[i].Address = address
		if infos[i].Address == "127.0.0.1:9000" {
			infos[i].ID = uint64(1)
		} else if infos[i].Address == "127.0.0.1:9001" {
			infos[i].ID = uint64(2)
		} else {
			infos[i].ID = uint64(3)
		}
	}
	fmt.Println("INFOS")
	fmt.Println(infos)
	store.Set(context.Background(), infos)
	return store
}

func logFunc(l client.LogLevel, format string, a ...interface{}) {}
