package golang_embed

import (
	_ "embed"
	"io/fs"
	"io/ioutil"
	"testing"
)

//go:embed golang.png
var logo []byte

func TestByteArray(t *testing.T) {
	err := ioutil.WriteFile("golang_next.png", logo, fs.ModePerm) // write file
	if err != nil {
		panic(err)
	}
}
