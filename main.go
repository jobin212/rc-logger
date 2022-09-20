package main

import (
	"fmt"
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
	buffer chan string
}

func (l *Logger) log(msg string) {
	select {
	case l.buffer <- msg:
	default:
		go l.flush()
		l.buffer <- msg
	}
}

func (l *Logger) flush() {
	for {
		select {
		case msgToPrint := <-l.buffer:
			fmt.Println(msgToPrint)
		default:
			return
		}
	}
}

var wg sync.WaitGroup

func main() {
	logger := &Logger{
		buffer: make(chan string, capacity),
	}

	for worker := 0; worker < 5; worker++ {
		wg.Add(1)
		go func(workerId int) {
			defer wg.Done()
			for i := 0; i < 10; i++ {
				logger.log(fmt.Sprintf("Worker: %d Msg %d", workerId, i))
			}
		}(worker)
	}

	wg.Wait()
	logger.flush()
}
