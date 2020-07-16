package kodfs_dataserver

import "github.com/guoqingpeng/kodfs/kodfs_metadata"

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

	//更新状态时间戳
	Timestmp int64

	Blocks []kodfs_metadata.Block
}

func NewDataNode() *DataNode {
	return &DataNode{}
}
