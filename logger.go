package logger

import (
	"io"
	"log"
	"runtime"
	"time"
)

const (
	LevelFatal = 1 << iota
	LevelError
	LevelWarn
	LevelInfo
	LevelDebug
)

var LogLevel int

func init() {
	// reset LogLevel
	LogLevel = LevelDebug
	log.SetFlags(0)
}

func SetOutput(w io.Writer) {
	log.SetOutput(w)
}

func Fatal(format string, v ...interface{}) {
	printf(LevelFatal, format, v...)
}

func Error(format string, v ...interface{}) {
	printf(LevelError, format, v...)
}

func Warn(format string, v ...interface{}) {
	printf(LevelWarn, format, v...)
}

func Info(format string, v ...interface{}) {
	printf(LevelInfo, format, v...)
}

func Debug(format string, v ...interface{}) {
	printf(LevelDebug, format, v...)
}

// Cheap integer to fixed-width decimal ASCII. Give a negative width to avoid zero-padding.
func itoa(buf *[]byte, i int, wid int) {
	// Assemble decimal in reverse order.
	var b [20]byte
	bp := len(b) - 1
	for i >= 10 || wid > 1 {
		wid--
		q := i / 10
		b[bp] = byte('0' + i - q*10)
		bp--
		i = q
	}
	// i < 10
	b[bp] = byte('0' + i)
	*buf = append(*buf, b[bp:]...)
}

func formatHeader(logLevel int) string {
	var buf []byte
	switch logLevel {
	case LevelFatal:
		buf = []byte("[Fatal]")
	case LevelError:
		buf = []byte("[Error]")
	case LevelWarn:
		buf = []byte("[Warn]")
	case LevelInfo:
		buf = []byte("[Info]")
	case LevelDebug:
		buf = []byte("[Debug]")
	}

	// date
	now := time.Now()
	year, month, day := now.Date()
	itoa(&buf, year, 4)
	buf = append(buf, '/')
	itoa(&buf, int(month), 2)
	buf = append(buf, '/')
	itoa(&buf, day, 2)
	buf = append(buf, ' ')

	// time
	hour, min, sec := now.Clock()
	itoa(&buf, hour, 2)
	buf = append(buf, ':')
	itoa(&buf, min, 2)
	buf = append(buf, ':')
	itoa(&buf, sec, 2)
	buf = append(buf, '.')
	itoa(&buf, now.Nanosecond()/1e3, 6)
	buf = append(buf, ' ')

	// short file name, runtime.Caller is expensive, info level don't call
	if LogLevel != LevelInfo {
		_, file, line, ok := runtime.Caller(3)
		if !ok {
			file = "???"
			line = 0
		}
		short := file
		for i := len(file) - 1; i > 0; i-- {
			if file[i] == '/' {
				short = file[i+1:]
				break
			}
		}
		file = short

		buf = append(buf, file...)
		buf = append(buf, ':')
		itoa(&buf, line, -1)
		buf = append(buf, ": "...)
	}

	return string(buf)
}

func printf(logLevel int, format string, v ...interface{}) {
	log.SetPrefix(formatHeader(logLevel))
	if LogLevel >= logLevel {
		log.Printf(format, v...)
	}
}
