package main

import (
	"fmt"
	"log"
	"os"
	"shodan/shodan"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Usage: shodan searchterm")
	}

	apiKey := os.Getenv("SHODAN_API_KEY")

	s := shodan.New(apiKey)

	info, err := s.APIInfo()
	if err != nil {
		log.Panicln(err)
	}

	fmt.Printf(
		"Query Credits:%d\nScan Credits:%d\n\n",
		info.QueryCredits,
		info.ScanCredits)

	hotSearch, err := s.HotSearch(os.Args[1])
	if err != nil {
		log.Panicln(err)
	}

	for _, host := range hotSearch.Matches {
		fmt.Printf("%18s%8d\n", host.IPString, host.Port)
	}

}
