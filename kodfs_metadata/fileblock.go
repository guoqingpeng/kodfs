package kodfs_metadata

type Block struct {
	Block_id int

	Block_size int

	FileChunks []FileChunk
}

func NewBlock() *Block {
	return &Block{}
}
