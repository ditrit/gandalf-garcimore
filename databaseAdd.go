package main

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/canonical/go-dqlite/client"
	"github.com/pkg/errors"
)

func addNodesToLeader(idNode string, databaseClusterConnection string, defaultcluster []string) (err error) {
	var cluster *[]string
	cluster = &defaultcluster

	fmt.Println("ADDOUP")

	id, _ := strconv.Atoi(idNode)
	if err != nil {
		return errors.Wrapf(err, "%s is not a number", idNode)
	}
	if id == 0 {
		return fmt.Errorf("ID must be greater than zero")
	}
	if databaseClusterConnection == "" {
		databaseClusterConnection = fmt.Sprintf("127.0.0.1:918%d", id)
	}
	info := client.NodeInfo{
		ID:      uint64(id),
		Address: databaseClusterConnection,
	}
	fmt.Println(*cluster)
	client, err := getLeader(*cluster)
	if err != nil {
		return errors.Wrap(err, "can't connect to cluster leader")
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	if err := client.Add(ctx, info); err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		return errors.Wrap(err, "can't add node")
	}

	return nil
}
