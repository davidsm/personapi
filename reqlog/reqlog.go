package reqlog

import (
	"log"
	"os"
)

const logFlags = log.LstdFlags

func newLogger(file *os.File) *log.Logger {
	return log.New(file, "", logFlags)
}

func Stdout() *log.Logger {
	return newLogger(os.Stdout)
}

func File(path string) (*log.Logger, error) {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0660)
	if err != nil {
		return nil, err
	}
	return newLogger(file), nil
}
