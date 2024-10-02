package main

import (
	"day10/udp/common"
	"fmt"
	"net"
)

func main() {
	raddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:5658")
	common.CheckErr(err)
	conn, err := net.DialUDP("udp", nil, raddr) //立即成功返回“connection”(虚拟的)
	common.CheckErr(err)

	_, err = conn.Write([]byte{'h', 'e', 'l', 'l', 'o'})
	conn.Write([]byte("golang"))
	conn.Write([]byte("中国人"))
	common.CheckErr(err)

	response := make([]byte, 1024)
	n, err := conn.Read(response)
	common.CheckErr(err)
	fmt.Printf("get response %s\n", response[:n])

	conn.Close()
}
