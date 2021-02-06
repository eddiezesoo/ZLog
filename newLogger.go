// Author:eddiezesoo zhuangyisheng@qq.com 2021-02
package ZLog

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

type Logger struct {
	logger      *log.Logger
	rate        Rate
	filePrefix  string
	timeCompare string
	logFile     *os.File
}

func NewLogger(logFilePrefix string, logRate Rate) *Logger {
	return &Logger{
		rate:       logRate,
		filePrefix: logFilePrefix,
	}
}

func (l *Logger) DEBUG(v ...interface{}) {
	l.checkTime()
	l.linePrefix(DEBUG)
	l.logger.Println(v)
}

func (l *Logger) Info(v ...interface{}) {
	l.checkTime()
	l.linePrefix(INFO)
	l.logger.Println(v)
}

func (l *Logger) WARN(v ...interface{}) {
	l.checkTime()
	l.linePrefix(WARN)
	l.logger.Println(v)
}

func (l *Logger) ERROR(v ...interface{}) {
	l.checkTime()
	l.linePrefix(ERROR)
	l.logger.Println(v)
}

func (l *Logger) checkTime() {
	if l.timeCompare != time.Now().Format(TimeFormat[l.rate]) {
		_ = l.logFile.Close()
		l.logFile = openLogFile(fmt.Sprintf("%s%s_%s.log",
			logSavePath,
			l.filePrefix,
			time.Now().Format(TimeFormat[l.rate]),
		))
		l.logger = log.New(l.logFile, "", log.LstdFlags)
		l.timeCompare = time.Now().Format(TimeFormat[l.rate])
	}
}

func (l *Logger) linePrefix(level Level) {
	_, file, line, ok := runtime.Caller(2)
	if ok {
		l.logger.SetPrefix(fmt.Sprintf("[%s][%s:%d]", LevelFlags[level], filepath.Base(file), line))
	} else {
		l.logger.SetPrefix(fmt.Sprintf("[%s]", LevelFlags[level]))
	}
}
