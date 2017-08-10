package zhlog

import (
	"fmt"
	"io"
	"log"
)

const (
	LevelDebug = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
)

// LogLevel 日志级别
type LogLevel int

type zlog struct {
	level LogLevel
}

var logger zlog

func init() {
	log.SetFlags(log.Lshortfile | log.Ltime)
}

func (z *zlog) setLevel(lev LogLevel) error {
	if lev < LevelDebug || lev > LevelFatal {
		return fmt.Errorf("wrong loglevel:%d", lev)
	}
	z.level = lev
	return nil
}

func (z *zlog) debug(format string, data ...interface{}) {
	if z.level <= LevelDebug {
		log.SetPrefix("[D]")
		log.Output(3, fmt.Sprintf(format, data...))
	}
}

func (z *zlog) info(format string, data ...interface{}) {
	if z.level <= LevelInfo {
		log.SetPrefix("[I]")
		log.Output(3, fmt.Sprintf(format, data...))
	}
}

func (z *zlog) warn(format string, data ...interface{}) {
	if z.level <= LevelWarn {
		log.SetPrefix("[W]")
		log.Output(3, fmt.Sprintf(format, data...))
	}
}

func (z *zlog) error(format string, data ...interface{}) {
	if z.level <= LevelError {
		log.SetPrefix("[E]")
		log.Output(3, fmt.Sprintf(format, data...))
	}
}

func (z *zlog) fatal(format string, data ...interface{}) {
	if z.level <= LevelFatal {
		log.SetPrefix("[F]")
		log.Output(3, fmt.Sprintf(format, data...))
	}
}

func SetOutPut(writer io.Writer) {
	log.SetOutput(writer)
}

//======================================================
// Wrapers
//======================================================

func SetLevel(lev LogLevel) error {
	return logger.setLevel(lev)
}

func Debug(format string, data ...interface{}) {
	logger.debug(format, data...)
}

func Info(format string, data ...interface{}) {
	logger.info(format, data...)
}

func Warn(format string, data ...interface{}) {
	logger.warn(format, data...)
}

func Error(format string, data ...interface{}) {
	logger.error(format, data...)
}

func Fatal(format string, data ...interface{}) {
	logger.fatal(format, data...)
}
