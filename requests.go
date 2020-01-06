// Package sitooutils contains utility functions for working with Sitoo.
package sitooutils

import (
	"bytes"
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

//BasicAuth - Functon to convert string base64
func BasicAuth(user, pass string) string {
	auth := user + ":" + pass
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func Node(accountNo string) string {
	var node string
	if string(accountNo[0]) == "9" {
		node = "-sandbox"
	} else if string(accountNo[0:3]) == "130" {
		node = "130"
	} else if string(accountNo[0:3]) == "100" {
		node = ""
	} else if string(accountNo[0:3]) == "201" {
		node = "201"
	} else if string(accountNo[0:3]) == "202" {
		node = "202"
	}
	return node
}

//GetSitoo - Function to GET data from Sitoo
func GetSitoo(endpoint string, account string, password string) []byte {
	accountSplit := strings.Split(account, "-")
	accountNo := accountSplit[0]

	req, err := http.NewRequest("GET", "https://api"+Node(accountNo)+".mysitoo.com/v2/accounts/"+accountNo+endpoint, nil)
	req.Header.Add("content-type", "application/json")
	req.Header.Add("Authorization", "Basic "+BasicAuth(account, password))
	resp, err := http.DefaultClient.Do(req)

	log.WithFields(log.Fields{
		"requesttype": "GET",
		"account":     account,
		"endpoint":    endpoint,
		"body":        nil,
	}).Debug("Request sent")

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
		}).Debug("OK")
		return response
	}
	return response
}

//PostSitoo - Function to POST data to Sitoo
func PostSitoo(endpoint string, account string, password string, json []byte) []byte {
	accountSplit := strings.Split(account, "-")
	accountNo := accountSplit[0]

	req, err := http.NewRequest("POST", "https://api"+Node(accountNo)+".mysitoo.com/v2/accounts/"+accountNo+endpoint, bytes.NewBuffer(json))
	req.Header.Add("content-type", "application/json")
	req.Header.Add("Authorization", "Basic "+BasicAuth(account, password))
	resp, err := http.DefaultClient.Do(req)

	log.WithFields(log.Fields{
		"requesttype": "POST",
		"account":     account,
		"endpoint":    endpoint,
		"body":        json,
	}).Info("Request sent")

	if err != nil {
		log.WithFields(log.Fields{
			"requesttype": "POST/Error",
			"account":     account,
			"endpoint":    endpoint,
			"body":        json,
			"response":    err,
		}).Fatal("ERROR")
	}
	defer resp.Body.Close()

	response, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		log.WithFields(log.Fields{
			"requesttype": "POST/Response",
			"account":     account,
			"endpoint":    endpoint,
			"statuscode":  resp.StatusCode,
			"response":    string(response),
		}).Fatal("ERROR")

		os.Exit(1)
	} else {
		log.WithFields(log.Fields{
			"requesttype": "POST/Response",
			"account":     account,
			"endpoint":    endpoint,
			"statuscode":  resp.StatusCode,
			"response":    string(response),
		}).Info("OK")
		return response
	}
	return response
}

//PutSitoo - Function to PUT data to Sitoo
func PutSitoo(endpoint string, account string, password string, json []byte) []byte {
	accountSplit := strings.Split(account, "-")
	accountNo := accountSplit[0]

	req, err := http.NewRequest("PUT", "https://api"+Node(accountNo)+".mysitoo.com/v2/accounts/"+accountNo+endpoint, bytes.NewBuffer(json))
	req.Header.Add("content-type", "application/json")
	req.Header.Add("Authorization", "Basic "+BasicAuth(account, password))
	resp, err := http.DefaultClient.Do(req)

	log.WithFields(log.Fields{
		"requesttype": "PUT",
		"account":     account,
		"endpoint":    endpoint,
		"body":        json,
	}).Info("Request sent")

	if err != nil {
		log.WithFields(log.Fields{
			"requesttype": "PUT/Error",
			"account":     account,
			"endpoint":    endpoint,
			"body":        json,
			"response":    err,
		}).Fatal("ERROR")
	}
	defer resp.Body.Close()

	response, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		log.WithFields(log.Fields{
			"requesttype": "PUT/Response",
			"account":     account,
			"endpoint":    endpoint,
			"statuscode":  resp.StatusCode,
			"response":    string(response),
		}).Fatal("ERROR")
	} else {
		log.WithFields(log.Fields{
			"requesttype": "PUT/Response",
			"account":     account,
			"endpoint":    endpoint,
			"statuscode":  resp.StatusCode,
			"response":    string(response),
		}).Info("OK")
		return response
	}
	return response
}
