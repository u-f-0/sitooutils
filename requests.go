// Package sitooutils contains utility functions for working with Sitoo.
package sitooutils

import (
	"bytes"
	"encoding/base64"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

//BasicAuth - Functon to convert string base64
func BasicAuth(user, pass string) string {
	auth := user + ":" + pass
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

//GetSitoo - Function to GET data from Sitoo
func GetSitoo(endpoint string, account string, password string) []byte {
	accountNo := strings.Split(account, "-")
	req, err := http.NewRequest("GET", "https://api130.mysitoo.com/v2/accounts/"+accountNo[0]+endpoint, nil)
	req.Header.Add("content-type", "application/json")
	req.Header.Add("Authorization", "Basic "+BasicAuth(account, password))
	resp, err := http.DefaultClient.Do(req)

	log.WithFields(log.Fields{
		"requesttype": "GET",
		"account":     account,
		"endpoint":    endpoint,
		"body":        nil,
	}).Info("Request sent")

	if err != nil {
		log.WithFields(log.Fields{
			"requesttype": "GET/Error",
			"account":     account,
			"endpoint":    endpoint,
			"body":        nil,
			"response":    err,
		}).Fatal("ERROR")
	}
	defer resp.Body.Close()

	response, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		log.WithFields(log.Fields{
			"requesttype": "GET/Response",
			"account":     account,
			"endpoint":    endpoint,
			"statuscode":  resp.StatusCode,
			"response":    string(response),
		}).Fatal("ERROR")

		os.Exit(1)
	} else {
		log.WithFields(log.Fields{
			"requesttype": "GET/Response",
			"account":     account,
			"endpoint":    endpoint,
			"statuscode":  resp.StatusCode,
			"response":    string(response),
		}).Info("OK")
		return response
	}
	return response
}

//PostSitoo - Function to POST data to Sitoo
func PostSitoo(endpoint string, account string, password string, json []byte) []byte {
	accountNo := strings.Split(account, "-")
	req, err := http.NewRequest("POST", "https://api-sandbox.mysitoo.com/v2/accounts/"+accountNo[0]+endpoint, bytes.NewBuffer(json))
	req.Header.Add("content-type", "application/json")
	req.Header.Add("Authorization", "Basic "+BasicAuth(account, password))
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	response, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		fmt.Println("StatusCode:", resp.StatusCode)
		log.Fatal(string(response))
		os.Exit(1)
	} else {
		return response
	}
	return response
}

//PutSitoo - Function to PUT data to Sitoo
func PutSitoo(endpoint string, account string, password string, json []byte) []byte {
	accountNo := strings.Split(account, "-")
	req, err := http.NewRequest("PUT", "https://api-sandbox.mysitoo.com/v2/accounts/"+accountNo[0]+endpoint, bytes.NewBuffer(json))
	req.Header.Add("content-type", "application/json")
	req.Header.Add("Authorization", "Basic "+BasicAuth(account, password))
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	response, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		fmt.Println("StatusCode:", resp.StatusCode)
		log.Fatal(string(response))
		os.Exit(1)
	} else {
		return response
	}
	return response
}
