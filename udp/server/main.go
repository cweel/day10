package main

import (
	"day10/udp/common"
	"fmt"
	"net"
)

func main() {
	laddr, err := net.ResolveUDPAddr("udp", ":5658")
	common.CheckErr(err)
	conn, err := net.ListenUDP("udp", laddr) //立即返回一个“connection”(虚拟的)
	common.CheckErr(err)

	defer conn.Close()
	requestBytes := make([]byte, 1024)
	for {
		n, raddr, err := conn.ReadFromUDP(requestBytes)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("read %s from client %s\n", requestBytes[:n], raddr)

		_, err = conn.WriteToUDP(requestBytes[:n], raddr)
		common.CheckErr(err)
	}
}
