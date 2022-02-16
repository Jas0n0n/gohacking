package main

import (
	"fmt"
	"log"
	"webVirusScan/awvs/api"
)

func main() {
	apikey := "1986ad8c0a5b3df4d7028d5f3c06e936c7982812192b94279907c2e461589ffa3"
  	client := api.New(apikey)
  	res,err := client.GetUsrInfo()
	if err !=nil {
		log.Println("get failed")
	}
	fmt.Println(res.Email)


}
