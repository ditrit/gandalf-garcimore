package database

import (
	"context"
	"fmt"
	"time"

	"github.com/canonical/go-dqlite/client"
	"github.com/pkg/errors"
)

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
