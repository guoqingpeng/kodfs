package kodfs_nameserver

import "github.com/guoqingpeng/kodfs/kodfs_dataserver"

type NameServer struct {
	nameserver_name string

	nameserver_ip int

	nameserver_port int

	dataservers []kodfs_dataserver.DataServer
}
