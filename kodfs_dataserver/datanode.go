package kodfs_dataserver

import "github.com/guoqingpeng/kodfs/kodfs_core"

type DataNode struct {

	//数据服务器名称
	dataserver_name string

	//数据服务器ip
	dataserver_ip int

	//数据服务器端口
	dataserver_port int

	//数据服务器总容量
	total_capacity int

	//数据服务器剩余容量
	left_capacity int

	//数据服务器数据磁盘位置
	data_dir string

	blocks []kodfs_core.Block
}
