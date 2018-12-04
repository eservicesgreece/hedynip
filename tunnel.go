package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

//Tunnels is the struct holding all found Hurricane Electric Tunnels in the account
type Tunnels struct {
	XMLName    xml.Name `xml:"tunnels"`
	TunnelList []Tunnel `xml:"tunnel"`
}

//Tunnel holds individual tunnel data
type Tunnel struct {
	XMLName     xml.Name `xml:"tunnel" json:"-"`
	ID          string   `xml:"id,attr"`
	Description string   `xml:"description"`
	ServerV4    string   `xml:"serverv4"`
	ClientV4    string   `xml:"clientv4"`
	ServerV6    string   `xml:"serverv6"`
	ClientV6    string   `xml:"clientv6"`
	Routed64    string   `xml:"routed64"`
	Routed48    string   `xml:"routed48"`
}

func listTunnels(username, password string) {
	rawXML := getDATA("https://tunnelbroker.net/tunnelInfo.php", username, password)
	var tunnels Tunnels
	xml.Unmarshal([]byte(rawXML), &tunnels)

	var printTunnels [][]string
	for k := range tunnels.TunnelList {
		printTunnels = append(printTunnels, []string{tunnels.TunnelList[k].ID, tunnels.TunnelList[k].Description, tunnels.TunnelList[k].ServerV4, tunnels.TunnelList[k].ClientV4, tunnels.TunnelList[k].ServerV6, tunnels.TunnelList[k].ClientV6, tunnels.TunnelList[k].Routed64, tunnels.TunnelList[k].Routed48})
	}

	printTunnelInfo(printTunnels)
}

func listTunnelbyID(tunnelID, username, password string) {
	rawXML := getDATA("https://tunnelbroker.net/tunnelInfo.php?tid="+tunnelID, username, password)
	var tunnels Tunnels
	err := xml.Unmarshal([]byte(rawXML), &tunnels)
	if err != nil {
		fmt.Println(err)
		fmt.Println("No Data Returned - Check Username & Password")
		os.Exit(1)
	}

	var TunnelbyID [][]string
	TunnelbyID = append(TunnelbyID, []string{tunnels.TunnelList[0].ID, tunnels.TunnelList[0].Description, tunnels.TunnelList[0].ServerV4, tunnels.TunnelList[0].ClientV4, tunnels.TunnelList[0].ServerV6, tunnels.TunnelList[0].ClientV6, tunnels.TunnelList[0].Routed64, tunnels.TunnelList[0].Routed48})

	printTunnelInfo(TunnelbyID)
}

func updateTunnelIPv4(TunnelID, currentIPv4, username, password string) {
	rawReply := getDATA("https://ipv4.tunnelbroker.net/nic/update?hostname="+TunnelID+"&myip="+currentIPv4, username, password)
	fmt.Println(rawReply)
}
