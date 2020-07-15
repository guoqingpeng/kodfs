package main

import (
	"encoding/json"
	"fmt"
	"github.com/guoqingpeng/kodfs/kodfs_dataserver"
	"net"
	"time"
)

func main() {
	for {

		serverAddr := "192.168.6.20:52255"
		tcpAddr, err := net.ResolveTCPAddr("tcp", serverAddr)

		if err != nil {
			fmt.Println("Resolve TCPAddr error", err)
		}
		conn, err := net.DialTCP("tcp", nil, tcpAddr)

		if err != nil {
			fmt.Println("connect server error", err)
		}

		ds := kodfs_dataserver.NewDataNode()
		ds.Dataserver_ip = "192.168.6.20"
		ds.Dataserver_port = 55255
		ds.Dataserver_name = "kaku1"
		ds.Data_dir = "./data/"
		ds.Left_capacity = 34 * 1025
		ds.Total_capacity = 2024 * 1024
		ds.Server_status = 0
		var dsBytes, _ = json.Marshal(ds)
		conn.Write(dsBytes)
		recv(conn)
		conn.Close()
		time.Sleep(time.Millisecond * 20)
	}
}

func recv(conn net.Conn) {
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)

	if err == nil {
		fmt.Println("read message from server:" + string(buffer[:n]))
	}

}
