// Author:eddiezesoo zhuangyisheng@qq.com 2021-02
package ZLog

type Rate int

const (
	Day Rate = iota
	Hour
	Minute
)

var TimeFormat = []string{"2006-01-02", "2006-01-02_15", "2006-01-02_1504"}
