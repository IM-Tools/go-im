/**
  @author:panliang
  @data:2021/8/16
  @note
**/
package zaplog

import (
	"encoding/json"
	"fmt"
	"im_app/core/utils"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

type ErrorLog struct {
	errors string
}

type ErrorInfo struct {
	Time         string `json:"time"`
	FileName     string `json:"file_name"`
	Function     string `json:"function"`
	ErrorMessage string `json:"error_message"`
	Line         int    `json:"line"`
}

/**
自定义日志方法
*/
func Warning(str string) {
	timeString := time.Unix(time.Now().Unix(), 0).Format("2006-01-02")
	fileName, line, functionName := "?", 0, "?"
	pc, fileName, line, ok := runtime.Caller(2)
	if ok {
		functionName = runtime.FuncForPC(pc).Name()
		functionName = filepath.Ext(functionName)
		functionName = strings.TrimPrefix(functionName, ".")
	}
	var msg = ErrorInfo{
		Time:         timeString,
		FileName:     fileName,
		ErrorMessage: str,
		Function:     functionName,
		Line:         line,
	}
	jsons, err := json.Marshal(msg)
	if err != nil {
		fmt.Println("json日志存储异常--", err)
	}
	errorJson := string(jsons) + "\n"
	path := utils.GetCurrentDirectory() + "/zaplog"
	logFile := path + "/" + timeString + "-error.zaplog"
	_, exist := os.Stat(path)
	if os.IsNotExist(exist) {
		os.Mkdir(path, os.ModePerm)
	}
	file, err := os.Open(logFile)

	if err != nil {
		files, err := os.Create(logFile)
		defer files.Close()
		if err != nil {
			fmt.Println(err)
		}
		files.Write([]byte(errorJson))
	} else {
		defer file.Close()
		file.Write([]byte(errorJson))
	}
}
func LogError(err error) {
	if err != nil {
		Warning(err.Error())
	}
}
