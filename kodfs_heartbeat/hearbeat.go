package kodfs_heartbeat

import (
	"encoding/json"
	"fmt"
	"github.com/guoqingpeng/kodfs/kodfs_dataserver"
	"github.com/guoqingpeng/kodfs/kodfs_nameserver"
	"net"
	"time"
)

//数据服务器发送健康心跳包给名称节点
func DataServerSendHeathStatus() {

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

	ds := kodfs_dataserver.NewDataNode()
	ds.Dataserver_ip = "127.0.0.1"
	ds.Dataserver_port = 55255
	ds.Dataserver_name = "kaku1"
	ds.Data_dir = "./data/"
	ds.Left_capacity = 34 * 1025
	ds.Total_capacity = 2024 * 1024
	ds.Server_status = 0
	var dsBytes, _ = json.Marshal(ds)
	fmt.Println(string((dsBytes)))

	conn.Write(dsBytes)
	dataServerRecv(conn)
}

//接受名称节点反馈回来的数据 主要是其他数据节点的存活信息
func dataServerRecv(conn net.Conn) {
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err == nil {
		fmt.Println("read message from server:" + string(buffer[:n]))
		fmt.Println("Message len:", n)
	}
}
