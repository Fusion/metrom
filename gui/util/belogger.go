package util

import (
	"fmt"
	"time"
)

type BELogger struct {
	logChan chan string
}

func (l *BELogger) Log(message string) {
	l.logChan <- fmt.Sprintf("%s: %s", time.Now().Format("2006-01-02 15:04"), message)
}

func (l *BELogger) GetLogs() *[]string {
	// I wonder how fast  would have to log to create starvation...
	fifo := []string{}
	for {
		select {
		case msg, ok := <-l.logChan:
			if !ok {
				return &fifo // so you closed the channel, you must be exiting...
			}
			fifo = append(fifo, msg)
		default:
			return &fifo // empty
		}
	}
}

var Logger = &BELogger{
	logChan: make(chan string, 1000),
}
