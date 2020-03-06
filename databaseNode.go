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

//DatabaseNode :
type DatabaseNode struct {
	clusterID                 uint64
	clusterDatabaseDirectory  string
	clusterDatabaseConnection string
	node                      *dqlite.Node
	clusterDatabaseClient     *DatabaseClient
}

//NewDatabaseNode :
func NewDatabaseNode(clusterID uint64, clusterDatabaseConnection string) *DatabaseNode {
	databaseNode := new(DatabaseNode)
	databaseNode.clusterID = clusterID
	databaseNode.clusterDatabaseConnection = clusterDatabaseConnection
	databaseNode.clusterDatabaseDirectory = "/tmp/"
	databaseNode.clusterDatabaseClient = NewDatabaseClient()

	return databaseNode
}

//run :
func (dn DatabaseNode) run() {
	err := dn.startNode(dn.clusterID, dn.clusterDatabaseDirectory, dn.clusterDatabaseConnection)
	fmt.Println(err)
	for {
		time.Sleep(time.Millisecond * time.Duration(100))
	}
}

//startNode :
func (dn DatabaseNode) startNode(id uint64, dir, address string) (err error) {
	nodeID := strconv.FormatUint(id, 10)
	nodeDir := filepath.Join(dir, nodeID)

	if errOs := os.MkdirAll(nodeDir, 0750); errOs != nil {
		return errors.Wrapf(errOs, "can't create %s", nodeDir)
	}

	node, err := dqlite.New(
		id, address, nodeDir,
		dqlite.WithBindAddress(address),
		dqlite.WithNetworkLatency(20*time.Millisecond),
	)

	if err != nil {
		return errors.Wrap(err, "failed to create node")
	}

	if err := node.Start(); err != nil {
		return errors.Wrap(err, "failed to start node")
	}

	return
}

//addNodesToLeader :
func (dn DatabaseNode) addNodesToLeader() (err error) {
	info := client.NodeInfo{
		ID:      uint64(dn.clusterID),
		Address: dn.clusterDatabaseConnection,
	}

	client, err := dn.clusterDatabaseClient.GetLeader()
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
