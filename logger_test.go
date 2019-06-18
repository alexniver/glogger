package logger

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func TestLogger(t *testing.T) {
	f, err := os.OpenFile("testlogfile.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer func() {
		err := f.Close()
		if nil != err {
			fmt.Println(err)
		}
	}()

	SetOutput(f)

	LogLevel = LevelDebug
	Debug("test debug, v : %+v", "value")
	LogLevel = LevelInfo
	Debug("test debug, v : %+v", "value")
	Info("test Info, v : %+v", "value")
	LogLevel = LevelWarn
	Debug("test debug, v : %+v", "value")
	Info("test Info, v : %+v", "value")
	Warn("test warn, v : %+v", "value")
	LogLevel = LevelError
	Debug("test debug, v : %+v", "value")
	Info("test Info, v : %+v", "value")
	Warn("test warn, v : %+v", "value")
	Error("test error, v : %+v", "value")
	LogLevel = LevelFatal
	Debug("test debug, v : %+v", "value")
	Info("test Info, v : %+v", "value")
	Warn("test warn, v : %+v", "value")
	Error("test error, v : %+v", "value")
	Fatal("test fatal, v : %+v", "value")

}
