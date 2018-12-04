package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func getDATA(url, username, password string) string {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	if username != "" && password != "" {
		req.SetBasicAuth(username, password)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	bodyText, err := ioutil.ReadAll(resp.Body)
	s := string(bodyText)
	return s
}
