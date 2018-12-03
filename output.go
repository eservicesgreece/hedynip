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
