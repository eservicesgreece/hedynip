package main

import (
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var version string
var buildstamp string
var githash string

var (
	app = kingpin.New("hedynip", fullversion())

	cmdusername = kingpin.Flag("username", "Account Username").Short('u').String()
	cmdpassword = kingpin.Flag("password", "Account Password").Short('p').String()

	listall = kingpin.Command("all", "List All Tunnels").Alias("a")

	listtunnel   = kingpin.Command("tunnel", "Lists Tunnel Details").Alias("t")
	listtunnelid = listtunnel.Flag("id", "Tunnel ID").Short('i').Short('i').Required().String()

	updateIP         = kingpin.Command("update", "Update IP").Alias("d")
	updateIPtunnelid = updateIP.Flag("id", "Tunnel ID").Short('i').Required().String()
	updateIPtunnelIP = updateIP.Flag("ip", "Set IP Manually").Short('m').String()

	getmyip = kingpin.Command("getmyip", "Get my IPv4").Alias("g")

	configlocation = kingpin.Flag("config", "Configuration file location").Short('c').String()
	showconfig     = kingpin.Command("showconfig", "Show Current Configuration").Alias("s")
)
