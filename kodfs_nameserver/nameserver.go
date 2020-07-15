package kodfs_nameserver

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/guoqingpeng/kodfs/kodfs_config"
	"github.com/guoqingpeng/kodfs/kodfs_dataserver"
	"net"
	"strconv"
	"time"
)

type NameServer struct {
	NameNode  *NameNode
	DataNodes []kodfs_dataserver.DataNode
}

func NewNameServer() *NameServer {

	return &NameServer{}
}

//开启监听，主要维护数据服务器的存活状态 dataserver info
func (ns *NameServer) NameServer_Start(cfg *kodfs_config.KodfsConfig) {

	ns.NameNode = NewNameNode("192.168.6.20", 52255)
	address := ns.NameNode.Nameserver_ip + ":" + strconv.Itoa(ns.NameNode.Nameserver_port)
	lner, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println("listener creat error", err)
	}
	fmt.Println("waiting for client")

	//开启一个线程定时轮训每个服务器的状态
	//每个dataserver的更新时间要求5秒已更新，未达到这个标准则视为掉线
	//需要发出警告
	go func() {
		// 放心这个线程是不会down 掉的
		for {
			t1 := time.NewTicker(2 * time.Second)
			select {
			case <-t1.C:
				NameServerSelfCheckDataServerStatus(ns)
			}
		}
	}()

	//处理心跳包
	for {
		conn, err := lner.Accept()
		if err != nil {
			fmt.Println("accept error", err)
		}
		go handleDataServerSocket(conn, ns)
	}

}

func handleDataServerSocket(conn net.Conn, ns *NameServer) {
	defer conn.Close()
	buffer := make([]byte, 10240)
	recvLen, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Read error", err)
	}
	strBuffer := string(buffer[:recvLen])
	fmt.Println("Message: ", strBuffer)
	dn := kodfs_dataserver.NewDataNode()
	err = json.Unmarshal(bytes.Trim(buffer, "\x00"), dn)

	if err != nil {
		fmt.Println("反序列化失败")
	}

	//开始更新数据服务器的状态
	if len(ns.DataNodes) == 0 {
		ns.DataNodes = make([]kodfs_dataserver.DataNode, 1)
		ns.DataNodes[0] = *dn
	} else {
		dnIn := false
		for i := 0; i < len(ns.DataNodes); i++ {

			if ns.DataNodes[i].Dataserver_name == dn.Dataserver_name {

				if ns.DataNodes[i].Server_status == 1 {
					fmt.Println("数据服务器重新上线！")
				} else {
					fmt.Println("数据服务器提交数据更新！")
				}

				ns.DataNodes[i] = *dn
				dnIn = true
				break
			}

		}
		if !dnIn {
			ns.DataNodes = append(ns.DataNodes, *dn)
			fmt.Println("新增加了一台数据服务器！")
		}

	}

	_, err = conn.Write([]byte(time.Now().String() + ":" + strBuffer))
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

//名称服务器自检数据服务器的健康状态
func NameServerSelfCheckDataServerStatus(ns *NameServer) {

	if len(ns.DataNodes) > 0 {
		for i := 0; i < len(ns.DataNodes); i++ {
			if ns.DataNodes[i].Server_status == 0 {
				if (time.Now().Unix() - ns.DataNodes[i].Timestmp) > 10 {
					fmt.Println("数据服务器长期没有联系上名称服务器，执行下线:" + ns.DataNodes[i].Dataserver_name)
					ns.DataNodes[i].Server_status = 1
				} else {
					fmt.Println("数据服务器正常运行:" + ns.DataNodes[i].Dataserver_name)
				}
			}
		}
	} else {
		fmt.Println("状态自检程序运行中，目前还没有数据服务器加入到集群中")

	}

}
