package main

import (
	"io"
	"log"
	"net"
	"sync"
	"time"
	"os"
)

type proxy struct {
	upstream
}

func (p *proxy) handle(conn net.Conn) {
	defer conn.Close()
	remote, err := net.Dial("tcp", p.Next())
	if err != nil {
		log.Printf("%s", err)
		return
	}
	defer remote.Close()
	ttl := time.Now().Add(2 * time.Second)
	conn.SetDeadline(ttl)
	remote.SetDeadline(ttl)
	go io.Copy(remote, conn)
	io.Copy(conn, remote)
}

type upstream struct {
	mu   sync.Mutex
	urls []string
	i    int
}

func (up *upstream) Next() string {
	up.mu.Lock()
	defer up.mu.Unlock()
	up.i++
	if up.i == len(up.urls) {
		up.i = 0
	}
	return up.urls[up.i]
}

func main() {
	port := os.Args[1]
	url := "localhost:"+port
	pxy := proxy{
		upstream{
			urls: []string{"localhost:3300", "localhost:3301"},
		},
	}
	srv, err := net.Listen("tcp", url)
	if err != nil {
		panic(err)
	}
	defer srv.Close()
	log.Printf("lb running on port %s", port)
	for {
		conn, _ := srv.Accept()
		go pxy.handle(conn)
	}
}
