package kodfs_core

type FileChunk struct {
	File_chunk_id int

	File_chunk_seq int

	fFile_chunk_size int

	File_chunk_postion_start_in_block int

	File_chunk_position_end_in_block int

	File_chunk_belong_to_block Block

	File_chunk_belong_to_file Kofile
}
