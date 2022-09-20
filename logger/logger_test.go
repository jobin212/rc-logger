package logger

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlush(t *testing.T) {
	b := bytes.Buffer{}
	testLogger := NewLogger(&b)
	msg := "hello!"
	testLogger.Log(msg)
	testLogger.Flush()
	assert.Equal(t, msg+"\n", b.String())
}
