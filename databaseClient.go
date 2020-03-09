package main

import (
	"context"
	"shoset/net"
	"time"

	"github.com/canonical/go-dqlite/client"
)

type DatabaseClient struct {
	DatabaseClientCluster []string
}

func NewDatabaseClient() (databaseClient *DatabaseClient) {
	databaseClient = new(DatabaseClient)
	//databaseClient.databaseClientCluster = cluster
	//databaseClient.DatabaseClientCluster = []string{"127.0.0.1:9000", "127.0.0.1:9001", "127.0.0.1:9002"}
	return
}

func (dc DatabaseClient) GetLeader() (*client.Client, error) {
	store := dc.GetStore()
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
	store.Set(context.Background(), infos)
	return store
}

func logFunc(l client.LogLevel, format string, a ...interface{}) {}
