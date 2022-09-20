package main

import (
	"fmt"
	"os"
	"time"

	"github.com/jobin212/rc-logger/logger"
)

var capacity int = 10
var timeThreshold time.Duration = 5 * time.Second

func main() {
	logger := logger.NewLogger(os.Stdout)
	defer logger.Flush()

	for i := 0; i < 15; i++ {
		logger.Log(fmt.Sprintf("Msg %d", i))
	}

	// var wg sync.WaitGroup
	// for worker := 0; worker < 5; worker++ {
	// 	wg.Add(1)
	// 	go func(workerId int) {
	// 		defer wg.Done()
	// 		for i := 0; i < 10; i++ {
	// 			logger.log(fmt.Sprintf("Worker: %d Msg %d", workerId, i))
	// 		}
	// 	}(worker)
	// }

	// logger.log("Hello, world")
	// time.Sleep(10 * time.Second)

	// wg.Wait()
}
