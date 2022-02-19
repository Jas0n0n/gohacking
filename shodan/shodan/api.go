package shodan

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type APIInfo struct {
	ScanCredits int `json:"scan_credits"`
	UsageLimits struct {
		ScanCredits  int `json:"scan_credits"`
		QueryCredits int `json:"query_credits"`
		MonitoredIps int `json:"monitored_ips"`
	} `json:"usage_limits"`
	Plan         string `json:"plan"`
	Https        bool   `json:"https"`
	Unlocked     bool   `json:"unlocked"`
	QueryCredits int    `json:"query_credits"`
	MonitoredIps int    `json:"monitored_ips"`
	UnlockedLeft int    `json:"unlocked_left"`
	Telnet       bool   `json:"telnet"`
}

func (c *Client)APIInfo() (*APIInfo,error)  {

	res,err :=http.Get(fmt.Sprintf("%s/api-info?key=%s",BaseURL,c.apikey))

	if err !=nil{
		return nil, err
	}

	defer res.Body.Close()

	var ret APIInfo
	if err:=json.NewDecoder(res.Body).Decode(&ret);err!=nil{
		return nil, err
	}

	return &ret,nil
}