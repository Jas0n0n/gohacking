package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"sort"
	"time"
)

var scanTarget string
var cmdLine = flag.NewFlagSet("tcpPortScan", flag.ExitOnError)

func init() {
	fmt.Println("starting")
	cmdLine.StringVar(&scanTarget, "scanTarget", "127.0.0.1", "go run main.go -a 10.200.122.111 ,Default Scan Address is 127.0.01")
	fmt.Println(scanTarget)
	fmt.Println("inited")
}

func worker(ports, results chan int) {
	for p := range ports {
		address := fmt.Sprintf("%s:%d", scanTarget, p)

		//fmt.Printf("Scanning:%s\n",address)

		conn, err := net.DialTimeout("tcp", address, 1*time.Second)

		if err != nil {
			results <- 0
			continue
		}
		conn.Close()
		results <- p
		//fmt.Printf("resultes:%v\n",&results)

	}
}

func main() {

	cmdLine.Parse(os.Args[1:])

	ports := make(chan int, 100)
	results := make(chan int)
	var openPorts []int

	for i := 0; i < cap(ports); i++ {
		go worker(ports, results)
	}

	go func() {
		for i := 1; i < 65535; i++ {
			ports <- i
		}
	}()

	for i := 1; i < 65535; i++ {
		port := <-results
		if port != 0 {
			openPorts = append(openPorts, port)
		}
	}
	close(ports)
	close(results)

	sort.Ints(openPorts)
	for _, port := range openPorts {
		fmt.Printf("%d open\n", port)
	}
}
