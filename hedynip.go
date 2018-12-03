package main

import (
	"fmt"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

func main() {
	kingpin.Version(fullversion())
	kingpin.CommandLine.HelpFlag.Short('h')
	kingpin.CommandLine.VersionFlag.Short('v')

	var appFlags = kingpin.Parse()

	switch appFlags {
	case listall.FullCommand():
		listTunnels(*listallusername, *listallpassword)

	case listtunnel.FullCommand():
		listTunnelbyID(*listtunnelid, *listtunnelusername, *listtunnelpassword)

	case updateIP.FullCommand():
		if *updateIPtunnelIP != "" {
			updateTunnelIPv4(*updateIPtunnelid, *updateIPtunnelIP, *updateIPtunnelusername, *updateIPtunnelpassword)
		} else {
			updateTunnelIPv4(*updateIPtunnelid, getIPv4(), *updateIPtunnelusername, *updateIPtunnelpassword)
		}

	case getmyip.FullCommand():
		fmt.Println(getIPv4())

	}
}
