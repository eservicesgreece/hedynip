package main

import (
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var version string
var buildstamp string
var githash string

var (
	app = kingpin.New("hedynip", fullversion())

	listall         = kingpin.Command("all", "List All Tunnels").Alias("a")
	listallusername = listall.Flag("username", "Account Username").Short('u').Required().String()
	listallpassword = listall.Flag("password", "Account Password").Short('p').Required().String()

	listtunnel         = kingpin.Command("tunnel", "Lists Tunnel Details").Alias("t")
	listtunnelid       = listtunnel.Flag("id", "Tunnel ID").Short('i').Short('i').Required().String()
	listtunnelusername = listtunnel.Flag("username", "Account Username").Short('u').Required().String()
	listtunnelpassword = listtunnel.Flag("password", "Account Password or Update Key if set").Short('p').Required().String()

	updateIP               = kingpin.Command("update", "Update IP").Alias("d")
	updateIPtunnelid       = updateIP.Flag("id", "Tunnel ID").Short('i').Required().String()
	updateIPtunnelIP       = updateIP.Flag("ip", "Set IP Manually").Short('m').String()
	updateIPtunnelusername = updateIP.Flag("username", "Account Username").Short('u').Required().String()
	updateIPtunnelpassword = updateIP.Flag("password", "Account Password or Update Key if set").Short('p').Required().String()

	getmyip = kingpin.Command("getmyip", "Get my IPv4").Alias("g")
)
