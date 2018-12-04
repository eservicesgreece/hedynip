package main

import (
	"fmt"
	"os"

	"github.com/gookit/config"
	"github.com/gookit/config/json"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

func main() {
	kingpin.Version(fullversion())
	kingpin.CommandLine.HelpFlag.Short('h')
	kingpin.CommandLine.VersionFlag.Short('v')
	var appFlags = kingpin.Parse()

	config.AddDriver(json.Driver)

	configFile := *configlocation

	if configFile != "" {
		err := config.LoadFiles(configFile)
		if err != nil {
			fmt.Printf("Config %s not found.\n", configFile)
		}
	} else {
		configFile = "hedynip.json"
		err := config.LoadFiles(configFile)
		if err != nil {
			fmt.Println("Config hedynip.json not found.")
		}
	}

	var username, password string
	credentials, _ := config.StringMap("credentials")

	if *cmdusername != "" {
		username = *cmdusername
	} else {
		username = credentials["username"]
	}

	if *cmdpassword != "" {
		password = *cmdpassword
	} else {
		password = credentials["password"]
	}

	switch appFlags {

	case listall.FullCommand():
		if username == "" || password == "" {
			fmt.Println("Error: Username & Password required")
			os.Exit(1)
		}

		listTunnels(username, password)

	case listtunnel.FullCommand():
		updatekey, _ := config.StringMap(*listtunnelid)
		password = updatekey["updatekey"]

		if username == "" || password == "" {
			fmt.Println("Error: Username & Password required")
			os.Exit(1)
		}

		listTunnelbyID(*listtunnelid, username, password)

	case updateIP.FullCommand():
		updatekey, _ := config.StringMap(*listtunnelid)
		password = updatekey["updatekey"]

		if username == "" || password == "" {
			fmt.Println("Error: Username & Password required")
			os.Exit(1)
		}

		if *updateIPtunnelIP != "" {
			updateTunnelIPv4(*updateIPtunnelid, *updateIPtunnelIP, username, password)
		} else {
			updateTunnelIPv4(*updateIPtunnelid, getIPv4(), username, password)
		}

	case getmyip.FullCommand():
		fmt.Println(getIPv4())

	case showconfig.FullCommand():
		err := config.LoadFiles(configFile)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("config data: \n %#v\n", config.Data())
		}

	}
}
