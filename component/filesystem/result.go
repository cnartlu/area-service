package filesystem

type ErrUpload map[string]error

type UploadResult struct {
	// Success 上传成功
	Success bool
	// ErrUpload 上传错误
	Errors ErrUpload
	// Filename 上传文件名
	Filename string
	// Filesize 文件大小
	Filesize int
	// OriginName 原始文件名称
	OriginName string
	// MimeType 文件的mime值
	MimeType string
}
