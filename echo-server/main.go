package main

import (
	"bufio"
	"log"
	"net"
)

//这是个回显服务器

//回显收到的数据
func echo(conn net.Conn) {
	defer conn.Close()

	reader :=bufio.NewReader(conn)
	s,err :=reader.ReadString('\n')
	if err!=nil{
		log.Fatalln("Unable to read data")
	}
	log.Printf("Read %d bytes: %s",len(s),s)

	log.Println("Writing data")
	writer :=bufio.NewWriter(conn)
	if _,err :=writer.WriteString(s);err!=nil{
		log.Fatalln("Unable to write data")
	}

	writer.Flush()
}

func main() {

	//绑定TCP端口20080
	listener, err := net.Listen("tcp", ":20080")
	if err != nil {
		log.Fatalln("Unable to bind to port")
	}
	log.Printf("Listening On 0.0.0.0:20080")

	for {
		//等待连接，在已建立的连接上创建net.conn
		conn, err := listener.Accept()
		log.Println("Received connection")
		if err != nil {
			log.Fatalln("Unable to accept connection")
		}
		go echo(conn)
	}

}
