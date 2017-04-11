package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	 "crypto/tls"
)

func main() {
	fmt.Println("start")
	// ignore ssl ca authrority(certificate signed by myself)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	httpRequest, err := http.NewRequest("GET", "https://nagonago.dip.jp/index.html", nil)

	resp, err := client.Do(httpRequest)
	if err != nil {
		fmt.Println("https error: %s", err.Error())
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("io error: %s", err.Error())
	}
	str := string(responseBody)
	fmt.Println(str)
	defer resp.Body.Close()
	fmt.Println("end")
}
