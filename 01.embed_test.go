package golang_embed

import (
	_ "embed"
	"fmt"
	"testing"
)

//go:embed version.txt
var version string // variable version akan otomatis berisi string dari file version.txt

func TestString(t *testing.T) {
	fmt.Println(version)
}
