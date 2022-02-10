package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main()  {

	baseURL := "https://127.0.0.1:3443/api/v1/me"

	//跳过证书验证
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	c := http.Client{Timeout: time.Duration(1)*time.Second,Transport: tr}

	req,err := http.NewRequest("GET", baseURL,nil)

	if err !=nil{
		fmt.Println("error %s",err)
		return
	}

	req.Header.Add("Content-Type","application/json")
	//AWVS API是通过Header中的X-Auth的值进行验证的
	req.Header.Add("X-Auth", `1986ad8c0a5b3df4d7028d5f3c06e936c7982812192b94279907c2e461589ffa3`)

    resp,err :=c.Do(req)
	if err !=nil{
		fmt.Println("error %s",err)
		return
	}

	defer resp.Body.Close()

    body,err :=ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Printf("Error %s", err)
		return
	}

	fmt.Printf("Body : %s", body)


}



