package api

import (
	"crypto/tls"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type UsrInfo struct {
	AccessAllGroups bool     `json:"access_all_groups"`
	ChildAccount    bool     `json:"child_account"`
	Email           string   `json:"email"`
	Enabled         bool     `json:"enabled"`
	FirstName       string   `json:"first_name"`
	Lang            string   `json:"lang"`
	Role            string   `json:"role"`
	UserId          string   `json:"user_id"`
	TotpEnabled     bool     `json:"totp_enabled"`
	Su              bool     `json:"su"`
	AccessRights    []string `json:"access_rights"`
	Notifications   struct {
		Scans         bool `json:"scans"`
		Reports       bool `json:"reports"`
		Targets       bool `json:"targets"`
		Updates       bool `json:"updates"`
		Workers       bool `json:"workers"`
		Attachments   bool `json:"attachments"`
		MonthlyStatus bool `json:"monthly_status"`
	} `json:"notifications"`
}

func (c *Client)GetUsrInfo() (*UsrInfo,error) {
	//绕过安全认证
	tr := http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	httpclient :=http.Client{Timeout: time.Duration(1)*time.Second,Transport: &tr}

	req,err := http.NewRequest("GET",BaseURL,nil)
	if err !=nil {
		log.Panicf("error %s",err)
		return nil, err
	}

	req.Header.Add("Content-Type",`application/json`)
	req.Header.Add("X-Auth",c.apikey)

	resp,err := httpclient.Do(req)
	if err !=nil {
		log.Panicf("error %s",err)
		return nil, err
	}

	defer resp.Body.Close()

	var ret UsrInfo
	if err:=json.NewDecoder(resp.Body).Decode(&ret);err!=nil{
		return nil,err
	}
	return &ret,nil
}