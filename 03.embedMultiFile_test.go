package golang_embed

import (
	"embed"
	"fmt"
	"testing"
)

//go:embed files/a.txt
//go:embed files/b.txt
//go:embed files/c.txt
var files embed.FS

func TestMultipleFiles(t *testing.T) {
	a, _ := files.ReadFile("files/a.txt")
	fmt.Println(string(a))

	b, _ := files.ReadFile("files/b.txt")
	fmt.Println(string(b))

	c, _ := files.ReadFile("files/c.txt")
	fmt.Println(string(c))
}

//go:embed files/*.txt
var match embed.FS

func TestPatchMatcher(t *testing.T) {
	dir, _ := match.ReadDir("files")

	for _, entry := range dir {
		if !entry.IsDir() {
			fmt.Println(entry.Name())
			content, _ := files.ReadFile("files/" + entry.Name())
			fmt.Println("Content:", string(content))
		}
	}
}
