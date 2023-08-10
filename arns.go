package goarns

import (
	"fmt"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"time"
)

type ArNS struct {
	DreUrl      string
	ArNSAddress string
	HttpClient  *http.Client
}

func NewArNS(dreUrl string, arNSAddr string, timout time.Duration) *ArNS {

	// default timeout is 5s
	if timout == 0 {
		timout = 5 * time.Second
	}

	httpClient := &http.Client{
		Timeout: timout, // Set the timeout for HTTP requests
	}
	return &ArNS{
		DreUrl:      dreUrl,
		ArNSAddress: arNSAddr,
		HttpClient:  httpClient,
	}
}

func (a *ArNS) QueryLatestRecord(domain string) (txId string, err error) {
	// step1 query NameCA address
	caAddress, err := a.QueryNameCa(domain)
	if err != nil {
		return "", err
	}
	// step2 query latest txId
	//Currently, only level-1 domain name resolution is queried
	txId, err = a.GetArNSTxID(caAddress, "@")

	return

}

func (a *ArNS) QueryNameCa(domain string) (caAddress string, err error) {
	baseURL := a.DreUrl + "/contract/"

	// Construct the complete URL
	url := baseURL + "?id=" + a.ArNSAddress

	// Make the HTTP request using the custom HTTP client
	response, err := a.HttpClient.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	// Check the response status code
	if response.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected response status: %s", response.Status)
	}

	// Read the response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	value := gjson.Get(string(body), "state.records."+domain+".contractTxId")

	if !value.Exists() {
		return "", fmt.Errorf("domain %s not exist", domain)
	}

	return value.String(), nil

}

func (a *ArNS) GetArNSTxID(caAddress string, domain string) (txId string, err error) {

	baseURL := a.DreUrl + "/contract/"

	// Construct the complete URL
	url := baseURL + "?id=" + caAddress

	// Make the HTTP request using the custom HTTP client
	response, err := a.HttpClient.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	// Check the response status code
	if response.StatusCode != http.StatusOK {
		return "", fmt.Errorf("GetArNSTxID: unexpected response status: %s", response.Status)
	}

	// Read the response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	value := gjson.Get(string(body), "state.records."+domain+".transactionId")

	if !value.Exists() {
		return "", fmt.Errorf("GetArNSTxID: domain %s not exist", domain)
	}

	return value.String(), nil

}
