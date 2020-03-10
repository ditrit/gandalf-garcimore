package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	dqlite "github.com/canonical/go-dqlite"
	"github.com/pkg/errors"
)

type DatabaseNodeCluster struct {
	databaseClusterDirectory  string
	databaseClusterConnection string
	databaseClusterId         uint64
	databaseClusterNodes      map[string]*dqlite.Node
}

func NewDatabaseNodeCluster(databaseClusterDirectory string, databaseClusterConnection string, databaseClusterId uint64) (databaseNodeCluster *DatabaseNodeCluster) {
	databaseNodeCluster = new(DatabaseNodeCluster)
	databaseNodeCluster.databaseClusterDirectory = databaseClusterDirectory
	databaseNodeCluster.databaseClusterConnection = databaseClusterConnection
	databaseNodeCluster.databaseClusterId = databaseClusterId
	databaseNodeCluster.databaseClusterNodes = make(map[string]*dqlite.Node)

	return
}

func (dc DatabaseNodeCluster) Run() {
	//RUN
	fmt.Println("START")
	err := dc.startNode(dc.databaseClusterId, dc.databaseClusterDirectory, dc.databaseClusterConnection)
	fmt.Println("ERR")
	fmt.Println(err)

	//INIT DB
	time.Sleep(time.Second * time.Duration(5))
}

func (dc DatabaseNodeCluster) startNode(id uint64, dir, address string) (err error) {

	if id == 0 {
		return fmt.Errorf("ID must be greater than zero")
	}
	if address == "" {
		address = fmt.Sprintf("127.0.0.1:918%d", id)
	}
	dir = filepath.Join(dir, string(id))
	if err := os.MkdirAll(dir, 0755); err != nil {
		return errors.Wrapf(err, "can't create %s", dir)
	}
	node, err := dqlite.New(
		uint64(id), address, dir,
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

/*
func (dc DatabaseNodeCluster) initDatabaseCluster(databaseClient *DatabaseClient) error {
	driver, err := driver.New(getStore())
	if err != nil {
		return errors.Wrapf(err, "failed to create dqlite driver")
	}

	sql.Register("dqlite", driver)

	db, err := sql.Open("dqlite", "context.db")
	if err != nil {
		return errors.Wrap(err, "can't open demo database")
	}
	defer db.Close()

	//TENANT
	if _, err := db.Exec("CREATE TABLE IF NOT EXISTS tenant (id INTEGER NOT NULL PRIMARY KEY, name TEXT NOT NULL)"); err != nil {
		fmt.Println(err)
		return errors.Wrap(err, "can't create tenant table")
	}

	if _, err := db.Exec("INSERT INTO tenant (name) values (?)", "test"); err != nil {
		return errors.Wrap(err, "can't update key")
	}

	//CONNECTORTYPE
	if _, err := db.Exec("CREATE TABLE IF NOT EXISTS connector_type (id INTEGER NOT NULL PRIMARY KEY, name TEXT NOT NULL)"); err != nil {
		return errors.Wrap(err, "can't create connector_type table")
	}

	if _, err := db.Exec("INSERT INTO connector_type (name) values (?)", "test"); err != nil {
		fmt.Println(err)
		return errors.Wrap(err, "can't update key")
	}

	//COMMAND TYPE
	if _, err := db.Exec("CREATE TABLE IF NOT EXISTS command_type (id INTEGER NOT NULL PRIMARY KEY, name TEXT NOT NULL)"); err != nil {
		return errors.Wrap(err, "can't create command_type table")
	}

	if _, err := db.Exec("INSERT INTO command_type (name) values (?)", "test"); err != nil {
		return errors.Wrap(err, "can't update key")
	}

	//AGGREGATOR
	if _, err := db.Exec("CREATE TABLE IF NOT EXISTS aggregator (id INTEGER NOT NULL PRIMARY KEY, name TEXT NOT NULL)"); err != nil {
		fmt.Println(err)
		return errors.Wrap(err, "can't create aggregator table")
	}

	if _, err := db.Exec("INSERT INTO aggregator (name) values (?)", "aggregator1"); err != nil {
		return errors.Wrap(err, "can't update key")
	}

	if _, err := db.Exec("INSERT INTO aggregator (name) values (?)", "aggregator2"); err != nil {
		return errors.Wrap(err, "can't update key")
	}

	//CONNECTOR
	if _, err := db.Exec("CREATE TABLE IF NOT EXISTS connector (id INTEGER NOT NULL PRIMARY KEY, name TEXT NOT NULL)"); err != nil {
		fmt.Println(err)
		return errors.Wrap(err, "can't create connector table")
	}

	if _, err := db.Exec("INSERT INTO connector (name) values (?)", "connector1"); err != nil {
		return errors.Wrap(err, "can't update key")
	}

	if _, err := db.Exec("INSERT INTO connector (name) values (?)", "connector2"); err != nil {
		return errors.Wrap(err, "can't update key")
	}

	//APPLICAION CONTEXT
	if _, err := db.Exec("CREATE TABLE IF NOT EXISTS application_context (id INTEGER NOT NULL PRIMARY KEY, name TEXT NOT NULL, tenant INTEGER NOT NULL, connector_type INTEGER NOT NULL, command_type INTEGER NOT NULL, aggregator_destination INTEGER NOT NULL, connector_destination INTEGER NOT NULL, FOREIGN KEY(tenant) REFERENCES tenant(id), FOREIGN KEY(connector_type) REFERENCES connector_type(id), FOREIGN KEY(command_type) REFERENCES command_type(id), FOREIGN KEY(aggregator_destination) REFERENCES aggregator(id), FOREIGN KEY(connector_destination) REFERENCES connector(id))"); err != nil {
		fmt.Println(err)
		return errors.Wrap(err, "can't create application_context table")
	}

	if _, err := db.Exec("INSERT INTO application_context (name, tenant, connector_type, command_type, aggregator_destination, connector_destination) values (?, ?, ?, ?, ?, ?)",
		"test", 1, 1, 1, 1, 1); err != nil {
		return errors.Wrap(err, "can't update key")
	}

	return nil
}
*/
