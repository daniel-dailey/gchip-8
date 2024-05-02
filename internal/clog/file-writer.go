package clog

import (
	"os"
)

type ClogFileWriter struct {
	file *os.File
}

func (cfw *ClogFileWriter) Write(p []byte) (n int, err error) {
	n, err = cfw.file.Write(p)
	if err != nil {
		return 0, err
	}
	return n, nil
}

func (cfw *ClogFileWriter) Close() error {
	err := cfw.file.Close()
	if err != nil {
		return err
	}
	return nil
}

func NewClogFileWriter(f *os.File) *ClogFileWriter {
	return &ClogFileWriter{
		file: f,
	}
}
