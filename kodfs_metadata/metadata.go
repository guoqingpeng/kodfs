package kodfs_metadata

type KofileMetaData struct {
	File_id int

	File_name string

	File_type string

	FileSize int

	File_chunk_count int

	Chunks []FileChunk
}
