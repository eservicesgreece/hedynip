package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func getTunnelXML(url, username, password string) string {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.SetBasicAuth(username, password)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	bodyText, err := ioutil.ReadAll(resp.Body)
	s := string(bodyText)
	return s
}

func getIPv4Json(url string) string {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	bodyText, err := ioutil.ReadAll(resp.Body)
	s := string(bodyText)

	return s
}
