package kodfs_metadata

import "github.com/guoqingpeng/kodfs/kodfs_nameserver"

type Block struct {
	Block_id int

	Block_size int

	FileChunks []FileChunk

	NameNode *kodfs_nameserver.NameNode
}
