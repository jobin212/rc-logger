package logger

import (
	"fmt"
	"io"
	"sync"
	"time"
)

/*
Implement a buffered logger that can be used by multiple callers concurrently to log messages.

The logger will buffer logged messages in memory and flush out buffered content when the size reaches a configured threshold
or the time has passed a configured threshold since the first buffered message was logged. You can assume the flush function
is already provided, or use print to screen as implementation. The flush operation should not block the log method.

*/

var capacity int = 10
var timeThreshold time.Duration = 5 * time.Second

type Logger struct {
	buffer       chan string
	flushLock    sync.Mutex
	timerLock    sync.Mutex
	timerRunning bool
	writer       io.Writer
}

func NewLogger(writer io.Writer) (l *Logger) {
	logger := &Logger{
		buffer: make(chan string, capacity),
		writer: writer,
	}
	return logger
}

func (l *Logger) Log(msg string) {
	l.flushIn(timeThreshold)
	select {
	case l.buffer <- msg:
	default:
		go l.Flush()
		l.buffer <- msg
	}
}

func (l *Logger) Flush() {
	l.flushLock.Lock()
	defer l.flushLock.Unlock()
	for {
		select {
		case msgToPrint := <-l.buffer:
			fmt.Fprintln(l.writer, msgToPrint)
		default:
			return
		}
	}
}

func (l *Logger) flushIn(t time.Duration) {
	go func() {
		l.timerLock.Lock()
		defer l.timerLock.Unlock()
		if l.timerRunning {
			return
		}
		l.timerRunning = true
		go func() {
			time.Sleep(t)
			l.Flush()
			l.timerLock.Lock()
			defer l.timerLock.Unlock()
			l.timerRunning = false
		}()
	}()
}
