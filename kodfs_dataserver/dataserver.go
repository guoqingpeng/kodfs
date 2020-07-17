package kodfs_dataserver

import (
	"encoding/json"
	"fmt"
	"github.com/guoqingpeng/kodfs/kodfs_metadata"
	"net"
	"strconv"
	"time"
)

type DataServer struct {
	DataNode *DataNode
}

//模拟生成block 每个磁盘生成10个block，每个block512 MB，每个block4MB，共16个chunk
func (ds *DataServer) InitDataServer() {

	blocks := make([]*kodfs_metadata.Block, 2000*100) // 每台数据服务器存储单位为T
	for i := 0; i < 2; i++ {

		//blockfile, _ := os.Create("./testBlocks/" + strconv.Itoa(i))

		//blockBytes := make([]byte, 1024*1024*64)

		//blockfile.Write(blockBytes)

		blockInfo := kodfs_metadata.NewBlock()
		blockInfo.Block_id = i
		blockInfo.Block_size = 1024 * 1024 * 64

		blockInfo.FileChunks = make([]*kodfs_metadata.FileChunk, 400)

		for j := 0; j < 4; j++ {
			chunkInfo := kodfs_metadata.NewFileChunk()
			chunkInfo.File_chunk_belong_to_block = blockInfo.Block_id
			chunkInfo.File_chunk_id = j
			chunkInfo.File_chunk_is_empty = 0
			chunkInfo.File_chunk_size = 1024 * 64
			chunkInfo.File_chunk_postion_start_in_block = j * chunkInfo.File_chunk_size
			chunkInfo.File_chunk_position_end_in_block = (j + 1) * chunkInfo.File_chunk_size
			blockInfo.FileChunks[j] = chunkInfo

		}
		blocks[i] = blockInfo
		ds.DataNode.Blocks = blocks
		fmt.Println("数据服务器初始化完成")
	}
}

//报告数据服务器自身的存活状态
func (ds *DataServer) SendDadaServerInfo() {
	dn := ds.DataNode
	//模拟不停的报告自身的状态
	serverAddr := "192.168.6.20:52255"
	tcpAddr, err := net.ResolveTCPAddr("tcp", serverAddr)

	if err != nil {
		fmt.Println("Resolve TCPAddr error", err)
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)

	if err != nil {
		fmt.Println("connect server error:", err)
		//continue
	}

	dn.Dataserver_ip = "192.168.6.20"
	dn.Dataserver_port = 55255
	dn.Dataserver_name = "kaku" + strconv.Itoa(8)
	dn.Data_dir = "./data/"
	dn.Left_capacity = 34 * 1024
	dn.Total_capacity = 2024 * 1024
	dn.Server_status = 0
	dn.Timestmp = time.Now().Unix()
	var dsBytes, _ = json.Marshal(ds)
	conn.Write(dsBytes)
	conn.Write([]byte("$"))
	recv(conn)
	conn.Close()

}

//接受名称服务器返回的数据
func recv(conn net.Conn) {

	for {
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err == nil {
			fmt.Println(string(buffer[:n]))
		}
		if n <= 0 {
			break
		}
	}
}
