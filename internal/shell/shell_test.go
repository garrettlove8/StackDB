package shell_test

import (
	"StackDB/internal/setup"
	"StackDB/internal/shell"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStart(t *testing.T) {
	setup.Setup()
	err := shell.Start()
	assert.ErrorIs(t, io.EOF, err)
}
