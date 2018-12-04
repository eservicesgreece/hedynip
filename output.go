package main

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

func printTunnelInfo(tunnelList [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Tunnel ID", "Description", "Server IPv4", "Client IPv4", "Server IPv6", "Client IPv6", "Routed /64", "Routed /48"})
	table.SetBorder(false)
	table.AppendBulk(tunnelList)
	table.Render()
}

func printTunnelrDNS(rDNSList [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"rDNS 1", "rDNS 2", "rDNS 3", "rDNS 4", "rDNS 5"})
	table.SetBorder(false)
	table.AppendBulk(rDNSList)
	table.Render()
}
