package kodfs_core

type FileChunk struct {
	file_chunk_id int

	file_chunk_seq int

	file_chunk_size int

	file_chunk_belong_to_block Block

	file_chunk_belong_to_file Kofile
}