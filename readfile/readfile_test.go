package readfile

import (

	//"read-pdf-go/controller"

	"testing"

	"github.com/dougadriquei/desafioneoway/utils"
	"github.com/stretchr/testify/assert"
)

func TestReadFile(t *testing.T) {
	t.Run("group", func(t *testing.T) {
		t.Parallel()
		t.Run("Test1", reader01)

	})
}

func reader01(t *testing.T) {
	csvfile := utils.OpenFile("test/base_teste.txt")
	defer csvfile.Close()
	result, err := ReadFile(csvfile)
	assert.True(t, len(result) == 49998)
	assert.True(t, result[49997].CpfCnpj == "042.098.288-40")
	assert.Nil(t, err)
}
