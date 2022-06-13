package util

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

// ReadFile 读取文件
func ReadFile(fileName string) ([]byte, error) {
	return ioutil.ReadFile(fileName)
}

// CreateDir 创建文件夹
func CreateDir(path string) error {
	// 判断一个文件是否存在
	_, err := os.Stat(path)
	if err != nil { // 文件读取错误
		if os.IsNotExist(err) { // 文件不存在
			err := os.MkdirAll(path, 0777)
			if err != nil {
				return fmt.Errorf("目录创建失败! : %s", path)
			}
			return nil // 创建成功
		} else { // 其他错误!
			return fmt.Errorf("目录操作失败! : %s", err)
		}
	}

	return nil
}

// WriteFile 写入文件(带缓冲的)
func WriteFile(fileName string, content []byte, flag int, perm os.FileMode) error {
	file, err := os.OpenFile(fileName, flag, perm)
	if err != nil {
		return fmt.Errorf("文件打开失败! : %s", err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	// 写入数据
	_, err = writer.Write(content)
	if err != nil {
		return fmt.Errorf("写入缓冲区失败! : %s", err)
	}
	writer.Flush()
	return nil
}

// FileUnit 文件的单位换算
var FileUnit = []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}

// TranslateFileSize 转换文件的大小
func TranslateFileSize(size uint64, bit uint) string {
	if size > 1024 {
		size /= 1024
		bit++
		return TranslateFileSize(size, bit)
	}
	unit := FileUnit[bit]
	return fmt.Sprintf("%d %s", size, unit)
}
