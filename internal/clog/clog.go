package clog

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"
)

type LogLevel int

const (
	LogLevelInfo LogLevel = iota
	LogLevelWarn
	LogLevelError
)

func (ll LogLevel) String() string {
	switch ll {
	case 0:
		return "INFO"
	case 1:
		return "WARN"
	case 2:
		return "ERROR"
	default:
		return "UNKNOWN"
	}
}

type Message struct {
	logger    *Log
	Intf      interface{} `json:"any"`
	Message   string      `json:"message"`
	Timestamp int64       `json:"timestamp"`
}

func (m *Message) Msg(message string) {
	m.Message = message
	m.logger.write(m)
}

func (m *Message) Any(any interface{}) *Message {
	m.Intf = any
	return m
}

type Log struct {
	level   LogLevel
	name    string
	writers map[string]io.Writer
}

func (l *Log) Info() *Message {
	return &Message{logger: l, Timestamp: time.Now().UnixMilli()}
}

func (l *Log) write(msg *Message) {
	buf, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	buf = append(buf, '\n')
	for _, w := range l.writers {
		_, err := w.Write(buf)
		if err != nil {
			panic(err)
		}

	}
}

func findLogFile(name string) *os.File {
	f, err := os.Stat(name)
	if err != nil || f == nil {
		return nil
	}
	logFile, err := os.OpenFile(name, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return nil
	}
	return logFile
}

func NewLog(level int, name, appname string) *Log {
	stdoutWriter := NewClogFileWriter(os.Stdout)
	fn := fmt.Sprintf("/Users/daniel.dailey/tmp/%s.log", appname)
	logFile := findLogFile(fn)
	if logFile == nil {
		newLogFile, err := os.Create(fn)
		if err != nil {
			panic(err)
		}
		logFile = newLogFile
	}
	fileWriter := NewClogFileWriter(logFile)
	logger := &Log{
		level: LogLevel(level),
		name:  name,
		writers: map[string]io.Writer{
			"stdout": stdoutWriter,
			"file":   fileWriter,
		},
	}
	return logger
}
