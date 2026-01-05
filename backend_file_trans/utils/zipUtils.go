package utils

import (
	"archive/zip"
	"bytes"
	"io"
)

// CompressFiles 将多个文件压缩成zip包
func CompressFiles(files map[string]io.Reader) (io.Reader, int64, error) {
	buf := new(bytes.Buffer)
	zw := zip.NewWriter(buf)

	for fileName, reader := range files {
		// 创建zip文件中的文件条目
		f, err := zw.Create(fileName)
		if err != nil {
			return nil, 0, err
		}

		// 将文件内容写入zip
		if _, err := io.Copy(f, reader); err != nil {
			return nil, 0, err
		}
	}

	// 关闭zip writer
	if err := zw.Close(); err != nil {
		return nil, 0, err
	}

	return bytes.NewReader(buf.Bytes()), int64(buf.Len()), nil
}
