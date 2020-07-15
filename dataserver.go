package main

import (
	"encoding/json"
	"fmt"
	"github.com/guoqingpeng/kodfs/kodfs_dataserver"
	"net"
	"strconv"
	"time"
)

func main() {
	for i := 0; i < 100; i++ {

		serverAddr := "192.168.6.20:52255"
		tcpAddr, err := net.ResolveTCPAddr("tcp", serverAddr)

		if err != nil {
			fmt.Println("Resolve TCPAddr error", err)
		}
		conn, err := net.DialTCP("tcp", nil, tcpAddr)

		if err != nil {
			fmt.Println("connect server error:", err)
			continue
		}

		ds := kodfs_dataserver.NewDataNode()
		ds.Dataserver_ip = "192.168.6.20"
		ds.Dataserver_port = 55255
		ds.Dataserver_name = "kaku" + strconv.Itoa(i)
		ds.Data_dir = "./data/"
		ds.Left_capacity = 34 * 1025
		ds.Total_capacity = 2024 * 1024
		ds.Server_status = 0
		ds.Timestmp = time.Now().Unix()
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
		fmt.Println(string(buffer[:n]))
	}

}
