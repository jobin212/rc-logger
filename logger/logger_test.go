package logger

import (
	"bytes"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFlush(t *testing.T) {
	var b bytes.Buffer
	testLogger := NewLogger(&b)
	msg := "hello!"
	testLogger.Log(msg)
	testLogger.Flush()
	assert.Equal(t, msg+"\n", b.String())
}

func TestTimeout(t *testing.T) {
	var b bytes.Buffer
	testLogger := NewLogger(&b)

	msg := "hello!"
	testLogger.buffer <- msg
	assert.Equal(t, 0, len(testLogger.timerChan))
	assert.Equal(t, "", b.String())

	ctx, cancel := context.WithCancel(context.Background())
	testLogger.flushIn(ctx)
	assert.Equal(t, 1, len(testLogger.timerChan))
	assert.Equal(t, "", b.String())
	cancel()
	for count := 0; count < 100 && len(testLogger.timerChan) > 0; count++ {
		time.Sleep(1 * time.Millisecond)
	}
	assert.Equal(t, 0, len(testLogger.timerChan))
	assert.Equal(t, msg+"\n", b.String())
}
