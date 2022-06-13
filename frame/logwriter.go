package frame

import (
	"acme/frame/frame-data"
	"acme/library/util"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

// 创建日志文件夹
func createLogDir(logPath, typeName string) string {
	// 年月文件夹
	year := time.Now().Format(util.Year)
	month := time.Now().Format(util.Month)
	// 创建基础文件夹
	baseDir := fmt.Sprintf("%s/%s/%s/%s", logPath, typeName, year, month)
	err := util.CreateDir(baseDir)
	if err != nil {
		panic(err.Error())
	}
	return baseDir
}

// 写入类型日志文件
func writeTypeLog(logPath, logType, content string) {
	// 创建日志文件夹
	baseDir := createLogDir(logPath, logType)
	logFile := fmt.Sprintf("%s/%s.log", baseDir, time.Now().Format(util.Day))
	// 写入文件
	err := util.WriteFile(logFile, []byte(content), os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0755)
	if err != nil {
		panic(err.Error())
	}
}

// WriteRequestLog 记录请求日志
func WriteRequestLog(logPath string, requestData frame_data.RequestData) {
	// json编码
	data, err := json.Marshal(requestData)
	if err != nil {
		panic(err.Error())
	}
	// 文件内容
	content := fmt.Sprintf("[ %s ] %s \r\n", requestData.Time, data)
	// 写入类型日志文件
	writeTypeLog(logPath, "request", content)
}

// WriteResponseLog 记录日志响应日志
func WriteResponseLog(logPath string, responseData frame_data.ResponseData) {
	// json编码
	data, err := json.Marshal(responseData)
	if err != nil {
		panic(err.Error())
	}
	// 写入文件
	content := fmt.Sprintf("[ %s ] %s \r\n", responseData.Time, data)
	// 写入类型日志文件
	writeTypeLog(logPath, "response", content)
}

// WritePanicLog 记录异常日志
func WritePanicLog(logPath string, panicData frame_data.PanicData) {
	// json编码
	data, err := json.Marshal(panicData)
	if err != nil {
		panic(err.Error())
	}
	// 文件内容
	content := fmt.Sprintf("[ %s ] %s \r\n", panicData.ReqData.Time, data)
	// 写入类型日志文件
	writeTypeLog(logPath, "panic", content)
}
