package main

import (
	"encoding/json"
	"fmt"
	"github.com/guoqingpeng/kodfs/kodfs_dataserver"
	"net"
	"os"
	"strconv"
	"time"
)

func main() {

	//模拟生成block 每个磁盘生成10个block，每个block64MB，每个block4MB，共16个chunk

	for i := 0; i < 16; i++ {

		block, _ := os.Create("./testBlocks/" + strconv.Itoa(i))

		blockBytes := make([]byte, 1024*1024*64)

		block.Write(blockBytes)

	}

	return

	//模拟不停的报告自身的状态
	for {

		for i := 0; i < 4000; i++ {

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
		}

		time.Sleep(time.Millisecond * 5000)
	}
}

func recv(conn net.Conn) {
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)

	if err == nil {
		fmt.Println(string(buffer[:n]))
	}

}
