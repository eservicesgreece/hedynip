package main

import (
	"encoding/json"
)

//IPv4 will hold the json for the public IP
type IPv4 struct {
	Publicip string
}

func getIPv4() string {
	rawJSON := getDATA("https://api4.eservices-greece.net/dns/myip.json", "", "")
	var publicIPv4 IPv4

	json.Unmarshal([]byte(rawJSON), &publicIPv4)
	return publicIPv4.Publicip
}
