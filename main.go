package main

import (
	"fmt"
	"net"
	"sync"

	"http/request"
	"http/response"
)

const (
	HOST = "127.0.0.1"
	PORT = "8080"
)

var connMutex = sync.Mutex{}

var headers = map[string]string{
	"Server": "golang",
}

func main() {
	server, err := net.Listen("tcp", HOST+":"+PORT)
	if err != nil {
		panic(err)
	}
	defer server.Close()
	fmt.Println("listening on " + HOST + ":" + PORT)

	for {
		conn, err := server.Accept()
		if err != nil {
			panic(err)
		}
		connMutex.Lock()

		fmt.Printf("received connection from %s\n", conn.RemoteAddr())
		res := response.Response(200, headers, "OK")

		_, err = request.NewRequest(conn, nil)
		if err != nil {
			res = response.Response(500, nil, err.Error())
		}

		fmt.Printf("sending response\n%s\n", res)
		conn.Write([]byte(res))
		conn.Close()
		connMutex.Unlock()
	}
}
