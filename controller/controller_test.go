package controller

import (

	//"read-pdf-go/controller"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadFileController(t *testing.T) {
	t.Run("group", func(t *testing.T) {
		t.Parallel()
		t.Run("Test1", readerController01)

	})
}

func readerController01(t *testing.T) {
	count, err := ReadFileController("../test/base_teste.txt")
	assert.True(t, count == 49998)
	assert.Nil(t, err)
}
