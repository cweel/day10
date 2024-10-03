package main

import (
	"day10/udp/common"
	"fmt"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/websocket"
)

type WsServer struct {
	//listener net.Listener
	addr    string
	upgrade *websocket.Upgrader
}

func NewWsServer(port int) *WsServer {
	ws := new(WsServer)
	ws.addr = "127.0.0.1" + ":" + strconv.Itoa(port)
	ws.upgrade = &websocket.Upgrader{
		HandshakeTimeout: 5 * time.Second,
		ReadBufferSize:   4096,
		WriteBufferSize:  4096,
	}
	return ws
}
func (ws *WsServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	conn, err := ws.upgrade.Upgrade(w, r, nil) //将http协议升级到websocket协议
	if err != nil {
		fmt.Printf("升级失败 %s\n", err)
		return
	}
	fmt.Printf("跟客户端%s建立好了websocket连接\n", r.RemoteAddr)

	go ws.handleOneConnection(conn)
}
func (ws *WsServer) handleOneConnection(conn *websocket.Conn) {
	defer func() {
		conn.Close()
	}()
	var request common.Request
	println("aa")
	{ //长链接
		println("cc")
		conn.SetReadDeadline(time.Now().Add(20 * time.Second))
		if err := conn.ReadJSON(&request); err != nil {
			if netError, ok := err.(net.Error); ok {
				if netError.Timeout() {
					fmt.Printf("发生了读超时")
					return
				}
			}
			fmt.Println(err)
			return
		}

		response := common.Response{Sum: request.A + request.B}
		if err := conn.WriteJSON(response); err != nil {
			fmt.Println(err)
		}
	}
	println("bb")
}
func (ws *WsServer) Start() error {
	// var err error
	// ws.listener, err = net.Listen("tcp", ws.addr)
	// if err != nil {
	// 	fmt.Printf("listen 失败 %s \n", err)
	// 	return err
	// }
	// if err = http.Serve(ws.listener, ws); err != nil {
	// 	fmt.Printf("serve 失败 %s \n", err)
	// 	return err
	// }
	if err := http.ListenAndServe(ws.addr, ws); err != nil {
		fmt.Printf("ListenAndServe 失败 %s \n", err)
		return err
	} else {
		return nil
	}
}
func main() {
	ws := NewWsServer(5656)
	ws.Start()
}
