package main

import (
	"day10/udp/common"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/websocket"
)

func main() {
	dialer := &websocket.Dialer{}
	header := http.Header{
		"Name": []string{"chen"},
	}
	conn, resp, err := dialer.Dial("ws://localhost:5656", header)
	if err != nil {
		fmt.Printf("dial 失败 %s \n", err)
		return
	}
	fmt.Println(resp.StatusCode)
	msg, _ := io.ReadAll(resp.Body)
	fmt.Println(string(msg))
	resp.Body.Close()
	defer conn.Close()
	fmt.Println("response header:")
	for k, v := range resp.Header {
		fmt.Printf("key=%s  value=%s\n", k, v)
	}

	request := common.Request{A: 3, B: 4}
	var response common.Response
	for i := 0; i < 5; i++ {
		conn.WriteJSON(request)
		conn.ReadJSON(&response)
		fmt.Println(response.Sum)
	}
}
