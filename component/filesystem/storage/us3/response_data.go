package us3

// DataSetItem 文件数据项
type DataSetItem struct {
	// 文件所属Bucket名称
	BucketName string
	// 文件名称,utf-8编码
	FileName string
	// 文件hash值
	Hash string
	// 文件mimetype
	MimeType string
	// 文件大小
	Size int64
	// 文件创建时间
	CreateTime int64
	// 文件修改时间
	ModifyTime int64
	// 文件存储类型，分别是标准、低频、归档，对应有效值：STANDARD, IA, ARCHIVE
	StorageClass string
}

// PrefixFileList 前缀文件列表
type PrefixFileList struct {
	// Bucket的名称
	BucketName string
	// Bucket的ID
	BucketId string
	// 下一个标志字符串，utf-8编码
	NextMarker string
	DataSet    []DataSetItem
}

// InitiateMultipartUploadReply 分配上传返回值
type InitiateMultipartUploadReply struct {
	// 本次分片上传的上传Id
	UploadId string
	// 分片的块大小
	BlkSize int
	// 上传文件所属Bucket的名称
	Bucket string
	// 上传文件在Bucket中的Key名称
	Key string
}
