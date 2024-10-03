package common

import "github.com/gorilla/websocket"

type Request struct {
	A int
	B int
}

type Response struct {
	Sum int
}
type Iter interface {
	handleOneConnection(*websocket.Conn)
}
