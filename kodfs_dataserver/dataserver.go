package kodfs_dataserver

import "github.com/guoqingpeng/kodfs/kodfs_core"

type DataServer struct {
	dataserver_name string

	dataserver_ip int

	dataserver_port int

	blocks []kodfs_core.Block
}
