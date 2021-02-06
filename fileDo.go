// Author:eddiezesoo zhuangyisheng@qq.com 2021-02
package ZLog

import (
	"log"
	"os"
)

var logSavePath = "logs/"

func openLogFile(logFile string) (file *os.File) {
	_, err := os.Stat(logFile)
	if os.IsNotExist(err) {
		mkLogDir()
	}
	if os.IsPermission(err) {
		log.Panic(err)
	}
	file, err = os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Fail to OpenFile :%v", err)
	}
	return
}

func mkLogDir() {
	dir, _ := os.Getwd()
	err := os.MkdirAll(dir+"/"+logSavePath, os.ModePerm)
	if err != nil {
		panic(err)
	}
}
