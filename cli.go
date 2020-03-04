package main

import (
	"crypto/rand"
	b64 "encoding/base64"
	"flag"
	"fmt"
	"os"
)

const path = "/etc/gandalf/"

func main() {

	var (
		config string
	)
	flag.Usage = func() {
		fmt.Printf("Usage of %s:\n", os.Args[0])
		fmt.Printf("  	gandalf mode command [options]")
		fmt.Printf("  	mode : cluster, aggragor, connector, agent\n")
		fmt.Printf("		cluster command : init, join\n")
		fmt.Printf("  			arguments:\n")
		fmt.Printf("  				logical name	  \n")
		fmt.Printf("  				bind address    \n")
		fmt.Printf("  				join address     \n")
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
				case "init":
					if len(args) >= 4 {
						LogicalName := args[2]
						BindAdd := args[3]
						//CREATE CLUSTER

						clusterInit(LogicalName, BindAdd)
						fmt.Println("Running Gandalf with:")
						fmt.Println("  Mode : " + mode)
						fmt.Println("  Logical Name : " + LogicalName)
						fmt.Println("  Bind Address : " + BindAdd)
						fmt.Println("  Config : " + config)
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
						clusterJoin(LogicalName, BindAdd, JoinAdd)
						fmt.Println("Running Gandalf with:")
						fmt.Println("  Mode : " + mode)
						fmt.Println("  Bind Address : " + BindAdd)
						fmt.Println("  Join Address : " + JoinAdd)
						fmt.Println("  Config : " + config)
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
