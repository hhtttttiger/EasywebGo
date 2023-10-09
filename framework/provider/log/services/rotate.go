package services

// HadeRotateLog 代表会进行切割的日志文件存储
type HadeRotateLog struct {
	HadeLog

	// 日志文件存储目录
	folder string
	// 日志文件名
	file string
}

// NewHadeRotateLog 实例化HadeRotateLog
func NewHadeRotateLog(params ...interface{}) (interface{}, error) {
	log := &HadeRotateLog{}
	return log, nil
}
