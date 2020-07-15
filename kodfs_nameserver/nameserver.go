package kodfs_nameserver

import (
	"encoding/json"
	"fmt"
	"github.com/guoqingpeng/kodfs/kodfs_config"
	"github.com/guoqingpeng/kodfs/kodfs_dataserver"
	"net"
	"strconv"
)

type NameServer struct {
	NameNode *NameNode
}

func NewNameServer() *NameServer {

	return &NameServer{}
}

//开启监听，主要维护数据服务器的存活状态 dataserver info
func (ns *NameServer) NameServer_Start(cfg *kodfs_config.KodfsConfig) {

	ns.NameNode = NewNameNode("127.0.0.1", 52255)
	address := ns.NameNode.Nameserver_ip + ":" + strconv.Itoa(ns.NameNode.Nameserver_port)
	lner, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println("listener creat error", err)
	}
	fmt.Println("waiting for client")
	for {
		conn, err := lner.Accept()
		if err != nil {
			fmt.Println("accept error", err)
		}
		go handleDataServerSocket(conn)
	}

}

func handleDataServerSocket(conn net.Conn) {
	defer conn.Close()
	buffer := make([]byte, 10240)
	recvLen, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Read error", err)
	}
	strBuffer := string(buffer[:recvLen])
	fmt.Println("Message: ", strBuffer)

	dn := kodfs_dataserver.NewDataNode()

	json.Unmarshal(buffer, dn)

	_, err = conn.Write([]byte("I am server, you message"))

	if err != nil {
		fmt.Println("send message error", err)
	}
}

//
func (ns *NameServer) WriteFileService() {

}

//返回一个文件的所有的chunk信息
func (ns *NameServer) ReadFileService() {

}
