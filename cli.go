package main

import (
	"crypto/rand"
	b64 "encoding/base64"
	"flag"
	"fmt"
	"garcimore/database"
	"os"
)

func main() {

	var (
		config string
	)
	flag.Usage = func() {
		fmt.Printf("Usage of %s:\n", os.Args[0])
		fmt.Printf("  gandalf mode command [options]")
		fmt.Printf("  mode : cluster, aggragor, connector, agent\n")
		fmt.Printf("	cluster command : init, join\n")
		fmt.Printf("      arguments:\n")
		fmt.Printf("  	    logical name	  \n")
		fmt.Printf("  		bind address    \n")
		fmt.Printf("  		join address     \n")
	}

	flag.StringVar(&config, "c", "", "")
	flag.StringVar(&config, "config", "", "")
	flag.Parse()
	args := flag.Args()

	if len(args) >= 1 {
		mode := args[0]
		switch mode {
		case "cluster":
			if len(args) >= 2 {
				command := args[1]

				switch command {
				case "database":
					done := make(chan bool)

					go database4(args[2], args[3])
					database.List(database.DefaultCluster)

					/* if len(args) >= 4 {
						//addr := args[3]
						done := make(chan bool)
						toto := []string{"127.0.0.1:10000", "127.0.0.1:10001", "127.0.0.1:10002"}
						go NewDatabaseNodeCluster("/home/orness/db/", toto).Run()
						<-done
					} else {
						flag.Usage()
					} */
					<-done
					break
				case "database2":
					done := make(chan bool)

					go database4(args[2], args[3])
					database.List(database.DefaultCluster)
					fmt.Println("ADD")
					err := database.AddNodesToLeader(args[2], args[3], database.DefaultCluster)
					fmt.Println(err)

					database.List(database.DefaultCluster)
					/* if len(args) >= 4 {
						//addr := args[3]
						done := make(chan bool)
						toto := []string{"127.0.0.1:10000", "127.0.0.1:10001", "127.0.0.1:10002"}
						go NewDatabaseNodeCluster("/home/orness/db/", toto).Run()
						<-done
					} else {
						flag.Usage()
					} */
					<-done

					break
				case "list":
					database.List(database.DefaultCluster)
					break
				case "add":
					fmt.Println("ADD")
					err := database.AddNodesToLeader(args[2], args[3], database.DefaultCluster)
					fmt.Println(err)

					break
				case "init":
					if len(args) >= 4 {
						LogicalName := args[2]
						BindAdd := args[3]
						//CREATE CLUSTER
						fmt.Println("INTI")

						fmt.Println("Running Gandalf with:")
						fmt.Println("  Mode : " + mode)
						fmt.Println("  Logical Name : " + LogicalName)
						fmt.Println("  Bind Address : " + BindAdd)
						fmt.Println("  Config : " + config)
						clusterInit(LogicalName, BindAdd)

					} else {
						flag.Usage()
					}
					break
				case "join": //join
					if len(args) >= 5 {
						LogicalName := args[2]
						BindAdd := args[3]
						JoinAdd := args[4]
						//CREATE CLUSTER
						fmt.Println("JOIN")
						fmt.Println("Running Gandalf with:")
						fmt.Println("  Mode : " + mode)
						fmt.Println("  Logical Name : " + LogicalName)
						fmt.Println("  Bind Address : " + BindAdd)
						fmt.Println("  Join Address : " + JoinAdd)
						fmt.Println("  Config : " + config)
						clusterJoin(LogicalName, BindAdd, JoinAdd)

					} else {
						flag.Usage()
					}
					break
				case "genkey":
					key := make([]byte, 32)
					_, err := rand.Read(key)
					if err != nil {
						fmt.Println("ERROR")
					}
					fmt.Println("Key : " + string(b64.URLEncoding.EncodeToString(key)))
					break
				default:
					break
				}
			} else {
				flag.Usage()
			}
		default:
			break
		case "agent":
			if len(args) >= 1 {
				command := args[1] //CREATE TENANT //CREATE USER //LIST USER //LIST TENANT
				switch command {
				case "CREATE_TENANT":
					if len(args) >= 2 {
						tenant := args[2] //TENANT
						secret := args[3] //SECRET
						fmt.Println(tenant)
						fmt.Println(secret)
					}
					break
				default:
					break
				}
				server := args[2]
				key := args[3]
				fmt.Println(server)
				fmt.Println(key)
			} else {
				flag.Usage()
			}
		}

	} else {
		flag.Usage()
	}
}
