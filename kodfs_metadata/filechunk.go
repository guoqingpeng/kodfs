package kodfs_metadata

type FileChunk struct {
	File_chunk_id int

	File_chunk_seq int

	File_chunk_size int

	File_chunk_postion_start_in_block int

	File_chunk_position_end_in_block int

	File_chunk_belong_to_block int

	File_chunk_is_empty int

	File_chunk_belong_to_file int
}

func NewFileChunk() *FileChunk {
	return &FileChunk{}
}
