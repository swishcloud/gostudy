package logger

import (
	"io"
	"log"
	"os"
	"sync"
)

func NewFileLogger(filepath string, prefix string) *log.Logger {
	f, err := os.OpenFile("log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	logger := log.New(f, prefix+" ", log.Ldate|log.Ltime|log.Lshortfile|log.LUTC)
	return logger
}
func NewLogger(out io.Writer, prefix string) *log.Logger {
	logger := log.New(out, prefix+" ", log.Ldate|log.Ltime|log.Lshortfile|log.LUTC)
	return logger
}

type FileConcurrentWriter struct {
	f *os.File
	m *sync.Mutex
}

func NewFileConcurrentWriter(filename string) *FileConcurrentWriter {
	fw := &FileConcurrentWriter{}
	f, err := os.OpenFile(filename,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	fw.f = f
	fw.m = &sync.Mutex{}
	return fw
}
func (w *FileConcurrentWriter) Write(p []byte) (n int, err error) {
	w.m.Lock()
	defer w.m.Unlock()
	return w.f.Write(p)
}
