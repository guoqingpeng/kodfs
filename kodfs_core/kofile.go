package kodfs_core

type Kofile struct {
	file_id int

	file_name string

	file_type string

	fileSize int

	file_chunk_count int

	chunks []FileChunk
}
