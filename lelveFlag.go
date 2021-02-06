// Author:eddiezesoo zhuangyisheng@qq.com 2021-02
package ZLog

type Level int

const (
	DEBUG Level = iota
	INFO
	WARN
	ERROR
)

var LevelFlags = []string{"DEBUG", "INFO", "WARN", "ERROR"}
