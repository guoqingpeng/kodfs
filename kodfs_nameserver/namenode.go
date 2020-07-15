package kodfs_nameserver

import "github.com/guoqingpeng/kodfs/kodfs_dataserver"

type NameNode struct {
	Nameserver_name string

	Nameserver_ip string

	Nameserver_port int

	Dataservers []kodfs_dataserver.DataServer
}

func NewNameNode(nameserver_ip string, nameserver_port int) *NameNode {
	return &NameNode{Nameserver_ip: nameserver_ip, Nameserver_port: nameserver_port}
}
