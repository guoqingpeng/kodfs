package kodfs_dataserver

import "github.com/guoqingpeng/kodfs/kodfs_core"

type DataNode struct {

	//数据服务器名称
	Dataserver_name string

	//数据服务器ip
	Dataserver_ip string

	//数据服务器端口
	Dataserver_port int

	//数据服务器总容量
	Total_capacity int

	//数据服务器剩余容量
	Left_capacity int

	//数据服务器数据磁盘位置
	Data_dir string

	//状态
	Server_status int

	Blocks []kodfs_core.Block
}

func NewDataNode(dataserver_name string) *DataNode {
	return &DataNode{Dataserver_name: dataserver_name}
}
