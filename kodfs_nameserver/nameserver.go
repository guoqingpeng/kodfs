package kodfs_nameserver

import (
	"fmt"
	"github.com/guoqingpeng/kodfs/kodfs_config"
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
	address := ns.NameNode.nameserver_ip + ":" + strconv.Itoa(ns.NameNode.nameserver_port)
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
	_, err = conn.Write([]byte("I am server, you message :" + strBuffer))
	if err != nil {
		fmt.Println("send message error", err)
	}
}

//
func (ns *NameServer) WriteFileService() {

}

//
func (ns *NameServer) ReadFileService() {

}
