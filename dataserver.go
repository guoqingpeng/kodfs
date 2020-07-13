package main

import (
	"encoding/json"
	"fmt"
	"github.com/guoqingpeng/kodfs/kodfs_dataserver"
	"net"
)

func main() {

	serverAddr := "127.0.0.1:52255"
	tcpAddr, err := net.ResolveTCPAddr("tcp", serverAddr)
	if err != nil {
		fmt.Println("Resolve TCPAddr error", err)
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	defer conn.Close()
	if err != nil {
		fmt.Println("connect server error", err)
	}

	ds := kodfs_dataserver.NewDataNode("kaku1")
	var dsBytes, _ = json.Marshal(ds)
	fmt.Println("ss" + string((dsBytes)))

	conn.Write(dsBytes)
	recv(conn)
}

func recv(conn net.Conn) {
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err == nil {
		fmt.Println("read message from server:" + string(buffer[:n]))
		fmt.Println("Message len:", n)
	}
}
