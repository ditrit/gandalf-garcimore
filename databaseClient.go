//Package database :
//File DatabaseClient.go
package main

import (
	"context"
	"database/sql"
	"time"
	"fmt"
	"github.com/canonical/go-dqlite/client"
	"github.com/canonical/go-dqlite/driver"
	"github.com/pkg/errors"
)

//DatabaseClient :
type DatabaseClient struct {
	Cluster    []string
	databaseDB map[string]*sql.DB
}

//NewDatabaseClient :
func NewDatabaseClient() (databaseClient *DatabaseClient) {
	databaseClient = new(DatabaseClient)
	databaseClient.databaseDB = make(map[string]*sql.DB)
	return
}

//GetLeader :
func (dc DatabaseClient) GetLeader() (*client.Client, error) {
	store := dc.getStore()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	return client.FindLeader(ctx, store, nil)
}

//getStore :
func (dc DatabaseClient) getStore() client.NodeStore {
	store := client.NewInmemNodeStore()
	fmt.Println("STORE CLIENT " )
	fmt.Println(dc.Cluster)
	if len(dc.Cluster) == 0 {
		// TODO handle this case
	}
	infos := make([]client.NodeInfo, len(dc.Cluster))
	for i, address := range dc.Cluster {
		infos[i].ID = uint64(i + 1)
		infos[i].Address = address
	}

	_ = store.Set(context.Background(), infos)

	return store
}

//open :
func (dc DatabaseClient) open(tenant string) (*sql.DB, error) {
	driver, err := driver.New(dc.getStore())
	if err != nil {
		return nil, errors.Wrapf(err, "failed to create dqlite driver")
	}
	sql.Register(tenant, driver)

	db, err := sql.Open(tenant, tenant+".db")
	if err != nil {
		return nil, errors.Wrap(err, "can't open database")
	}
	dc.databaseDB[tenant] = db
	defer db.Close()

	return db, nil
}

//getDatabase :
func (dc DatabaseClient) GetDatabase(tenant string) *sql.DB {
	if database, ok := dc.databaseDB[tenant]; ok {
		return database
	}
	database, _ := dc.open(tenant)
	return database
}
