package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ipQueryResponse struct {
	Status string `json:"status"`
	IP     string `json:"query"`
}

type DnsRecord struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	IP   string `json:"content"`
}

type dnsQueryResponse struct {
	Rec []DnsRecord `json:"result"`
}

type DnsZoneInfo struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

type getZonesResponse struct {
	Rec []DnsZoneInfo `json:"result"`
}

func GetPublicIP() (string, error) {
	qryUrl := "http://ip-api.com/json?fields=status,query"
	resp, err := http.Get(qryUrl)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var ipQry ipQueryResponse
	if err := json.Unmarshal(body, &ipQry); err != nil {
		return "", err
	}

	return ipQry.IP, nil
}

func GetZoneId(zoneName string, apiToken string) (string, error) {
	reqUrl := fmt.Sprintf("https://api.cloudflare.com/client/v4/zones?name=%s&status=active&page=1&per_page=20&order=status&direction=desc&match=all", zoneName)

	client := http.Client{}
	req, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		return "", err
	}

	req.Header = http.Header{
		"Content-Type":  {"application/json"},
		"Authorization": {fmt.Sprintf("Bearer %s", apiToken)},
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var getZonesQry getZonesResponse
	if err := json.Unmarshal(body, &getZonesQry); err != nil {
		return "", err
	}
	if len(getZonesQry.Rec) == 0 {
		return "", errors.New("DNS zone not found")
	}

	return getZonesQry.Rec[0].Id, nil
}

func GetDnsRecords(zoneId string, apiToken string) ([]DnsRecord, error) {
	reqUrl := fmt.Sprintf("https://api.cloudflare.com/client/v4/zones/%s/dns_records?type=A",
		zoneId)

	client := http.Client{}
	req, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		return nil, err
	}

	req.Header = http.Header{
		"Content-Type":  {"application/json"},
		"Authorization": {fmt.Sprintf("Bearer %s", apiToken)},
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var dnsQry dnsQueryResponse
	if err := json.Unmarshal(body, &dnsQry); err != nil {
		return nil, err
	}

	return dnsQry.Rec, nil
}
