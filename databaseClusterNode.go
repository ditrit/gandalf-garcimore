package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"

	dqlite "github.com/canonical/go-dqlite"
	"github.com/canonical/go-dqlite/client"
	"github.com/pkg/errors"
)

type DatabaseNodeCluster struct {
	databaseClusterDirectory  string
	databaseClusterConnection string
	databaseClusterId         uint64
	DatabaseClient            *DatabaseClient
	databaseClusterNodes      map[string]*dqlite.Node
}

func NewDatabaseNodeCluster(databaseClusterDirectory string, databaseClusterConnection string, databaseClusterId uint64) (databaseNodeCluster *DatabaseNodeCluster) {
	databaseNodeCluster = new(DatabaseNodeCluster)
	databaseNodeCluster.databaseClusterDirectory = databaseClusterDirectory
	databaseNodeCluster.databaseClusterConnection = databaseClusterConnection
	databaseNodeCluster.databaseClusterId = databaseClusterId
	databaseNodeCluster.databaseClusterNodes = make(map[string]*dqlite.Node)
	databaseNodeCluster.DatabaseClient = NewDatabaseClient()

	return
}

func (dc DatabaseNodeCluster) Run() {
	//RUN
	err := dc.startNode(dc.databaseClusterId, dc.databaseClusterDirectory, dc.databaseClusterConnection)
	fmt.Println(err)

}

func (dc DatabaseNodeCluster) startNode(id uint64, dir, address string) (err error) {

	nodeID := strconv.FormatUint(id, 10)
	nodeDir := filepath.Join(dir, nodeID)
	if err := os.MkdirAll(nodeDir, 0755); err != nil {
		return errors.Wrapf(err, "can't create %s", nodeDir)
	}
	node, err := dqlite.New(
		id, address, nodeDir,
		dqlite.WithBindAddress(address),
		dqlite.WithNetworkLatency(20*time.Second),
	)
	dc.databaseClusterNodes[nodeID] = node
	if err != nil {
		return errors.Wrap(err, "failed to create node")
	}
	if err := node.Start(); err != nil {
		return errors.Wrap(err, "failed to start node")
	}
	return
}

func (dc DatabaseNodeCluster) addNodesToLeader() (err error) {
	info := client.NodeInfo{
		ID:      dc.databaseClusterId,
		Address: dc.databaseClusterConnection,
	}

	client, err := dc.DatabaseClient.GetLeader()
	if err != nil {
		return errors.Wrap(err, "can't connect to cluster leader")
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	if err := client.Add(ctx, info); err != nil {
		return errors.Wrap(err, "can't add node")
	}
	return
}
