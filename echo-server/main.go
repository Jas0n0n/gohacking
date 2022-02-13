package main

import (
	"io"
	"log"
	"net"
)

//这是个回显服务器

//回显收到的数据
func echo(conn net.Conn) {
	defer conn.Close()

	//创建一个缓冲区来存储接收到的数据
	b := make([]byte, 512)

	for {

		//通过conn.Read接收数据到一个缓冲区
		size, err := conn.Read(b[0:])

		//接收到关闭信号
		if err == io.EOF {
			log.Println("Client disconnected")
			break
		}

		if err != nil {
			log.Println("Unexpected error")
			break
		}

		log.Printf("Received %d bytes:%s\n", size, string(b))

		//通过conn.Write发送数据
		log.Println("Writing data")
		if _, err := conn.Write(b[0:size]); err != nil {
			log.Fatalln("Unable to write data")
		}
	}
}

func main() {

	//绑定TCP端口20080
	lister, err := net.Listen("tcp", ":20080")
	if err != nil {
		log.Fatalln("Unable to bind to port")
	}
	log.Printf("Listening On 0.0.0.0:20080")

	for {
		//等待连接，在已建立的连接上创建net.conn
		conn, err := lister.Accept()
		log.Println("Received connection")
		if err != nil {
			log.Fatalln("Unable to accept connection")
		}
		go echo(conn)
	}

}
