package xyq

import (
	"fmt"
	"io"
	"log"
	"os"
)

var loggerCache map[string]*log.Logger = make(map[string](*log.Logger))

func NewLogger(fileName string) *log.Logger {
	if logger, ok := loggerCache[fileName]; ok {
		return logger
	}
	stdWriter := os.Stdout
	fileWriter, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		panic(fmt.Sprintf("File %s open error for reason {%s}", fileName, err))
	}
	logger := log.New(io.MultiWriter(stdWriter, fileWriter), "", log.Lshortfile|log.LstdFlags)
	loggerCache[fileName] = logger
	return logger
}
