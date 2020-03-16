package database

import (
	"context"
	"fmt"
	"time"

	"github.com/canonical/go-dqlite/client"
	"github.com/pkg/errors"
)

const defaultBaseAdd = "127.0.0.1:900"
const DefaultNodeDirectory = "/home/dev-ubuntu/db/"

var DefaultCluster = []string{"127.0.0.1:9000", "127.0.0.1:9001", "127.0.0.1:9002"}

func getLeader(cluster []string) (*client.Client, error) {

	store := getStore(cluster)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	return client.FindLeader(ctx, store)
}

func getStore(cluster []string) client.NodeStore {

	store := client.NewInmemNodeStore()
	if len(cluster) == 0 {
		cluster = DefaultCluster
	}
	infos := make([]client.NodeInfo, 3)
	for i, address := range cluster {
		infos[i].ID = uint64(i + 1)
		infos[i].Address = address
	}
	store.Set(context.Background(), infos)
	return store
}

func AddNodesToLeader(id int, nodeConnection string, defaultcluster []string) (err error) {
	var cluster *[]string
	cluster = &defaultcluster

	if err != nil {
		return errors.Wrapf(err, "%s is not a number", id)
	}
	if id == 0 {
		return fmt.Errorf("ID must be greater than zero")
	}
	if nodeConnection == "" {
		nodeConnection = fmt.Sprintf("%s%d", defaultBaseAdd, id)

	}
	info := client.NodeInfo{
		ID:      uint64(id),
		Address: nodeConnection,
	}

	client, err := getLeader(*cluster)
	if err != nil {
		return errors.Wrap(err, "can't connect to cluster leader")
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	if err := client.Add(ctx, info); err != nil {
		return errors.Wrap(err, "can't add node")
	}

	return nil
}
