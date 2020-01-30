// credit: https://github.com/apsdehal/go-logger/blob/master/logger.go

package logutil

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

var (
	// Map for te various codes of colors
	colors map[Level]string
	lvls   map[Level]string
)

// Color numbers for stdout
const (
	Black = (iota + 30)
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
)

type Level int

const (
	DEBUG Level = iota
	INFO
	NOTICE
	WARNING
	ERROR
	CRITICAL
)

func init() {
	colors = map[Level]string{
		CRITICAL: colorString(Magenta),
		ERROR:    colorString(Red),
		WARNING:  colorString(Yellow),
		NOTICE:   colorString(Green),
		INFO:     colorString(White),
		DEBUG:    colorString(Cyan),
	}
	lvls = map[Level]string{
		CRITICAL: "CRIT",
		ERROR:    "EROR",
		WARNING:  "WARN",
		NOTICE:   "NOTC",
		INFO:     "INFO",
		DEBUG:    "DEBG",
	}

	defaultLogger = NewLogger(DEBUG, "", os.Stdout)
}

func colorString(color int) string {
	return fmt.Sprintf("\033[%dm", int(color))
}

var defaultLogger *Logger

type Logger struct {
	lvl    Level
	format string
	output io.Writer
}

func SetDefaultLogger(logger *Logger) {
	defaultLogger = logger
}

func NewLogger(lvl Level, format string, output io.Writer) *Logger {
	return &Logger{
		lvl:    lvl,
		format: format,
		output: output,
	}
}

func (l *Logger) log(lvl Level, format string, args ...interface{}) {
	if l.lvl > lvl {
		return
	}

	msg := fmt.Sprintf(format, args...)
	formatedMsg := fmt.Sprintf("%s %s goroutine:%s/%d %s ▶ %s",
		time.Now().Format("2006-01-02 15:04:05"),
		lvls[lvl],
		l.getGoroutineId(),
		runtime.NumGoroutine(),
		l.callerInfo(3),
		msg)

	buf := &bytes.Buffer{}
	buf.Write([]byte(colors[lvl]))
	buf.Write([]byte(formatedMsg))
	buf.Write([]byte("\033[0m\n"))

	l.output.Write(buf.Bytes())
}

func (l *Logger) callerInfo(depth int) string {
	pc, file, lineno, ok := runtime.Caller(depth)
	if !ok {
		return "no caller info found"
	}

	return fmt.Sprintf("%s %s:%d", filepath.Base(file), filepath.Base(runtime.FuncForPC(pc).Name()), lineno)
}

func (l *Logger) getGoroutineId() string {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	return string(b)
}

// Critical logs a message at a Critical Level
func (l *Logger) Critical(format string, args ...interface{}) {
	l.log(CRITICAL, format, args...)
}

// Error logs a message at Error level
func (l *Logger) Error(format string, args ...interface{}) {
	l.log(ERROR, format, args...)
}

// Warning logs a message at Warning level
func (l *Logger) Warning(format string, args ...interface{}) {
	l.log(WARNING, format, args...)
}

// Notice logs a message at Notice level
func (l *Logger) Notice(format string, args ...interface{}) {
	l.log(NOTICE, format, args...)
}

// Info logs a message at Info level
func (l *Logger) Info(format string, args ...interface{}) {
	l.log(INFO, format, args...)
}

// Debug logs a message at Debug level
func (l *Logger) Debug(format string, args ...interface{}) {
	l.log(DEBUG, format, args...)
}

// Critical logs a message at a Critical Level
func Critical(format string, args ...interface{}) {
	defaultLogger.log(CRITICAL, format, args...)
}

// Error logs a message at Error level
func Error(format string, args ...interface{}) {
	defaultLogger.log(ERROR, format, args...)
}

// Warning logs a message at Warning level
func Warning(format string, args ...interface{}) {
	defaultLogger.log(WARNING, format, args...)
}

// Notice logs a message at Notice level
func Notice(format string, args ...interface{}) {
	defaultLogger.log(NOTICE, format, args...)
}

// Info logs a message at Info level
func Info(format string, args ...interface{}) {
	defaultLogger.log(INFO, format, args...)
}

// Debug logs a message at Debug level
func Debug(format string, args ...interface{}) {
	defaultLogger.log(DEBUG, format, args...)
}
