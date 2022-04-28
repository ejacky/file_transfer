package entity

type MultipartUploadInfo struct {
	FileHash   string
	FileSize   uint64
	UploadID   string
	ChunkSize  uint64
	ChunkCount uint64
	// 已经上传完成的分块索引列表
	ChunkExists []uint64
}
