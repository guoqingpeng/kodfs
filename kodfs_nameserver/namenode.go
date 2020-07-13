package kodfs_nameserver

import "github.com/guoqingpeng/kodfs/kodfs_dataserver"

type NameNode struct {
	nameserver_name string

	nameserver_ip int

	nameserver_port int

	dataservers []kodfs_dataserver.DataServer
}

func NewNameNode() *NameNode {
	return &NameNode{}
}
