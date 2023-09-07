package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

/* Example:
package babylonjs

import (
	"embed"
	"sync"
)

//go:embed babylon.js
var FS embed.FS

const jsFile = "babylon.js"

var js []byte
var once sync.Once

func JS() []byte {
	once.Do(func() {
		var err error
		js, err = FS.ReadFile(jsFile)
		if err != nil {
			panic(err)
		}
	})
	return js
}
*/

func makePkg(pkg, dir, jsFile string) error {
	println("makePkg:", pkg, dir, jsFile)
	jsPath := filepath.Join("node_modules", dir, jsFile)
	println(jsPath)

	code := fmt.Sprintf("package %s\n\n", pkg)
	code += `import (
	"embed"
	"sync"
)

`
	code += fmt.Sprintf("//go:embed %s\n", jsFile)
	code += "var FS embed.FS\n\n"
	code += fmt.Sprintf("const JSFile = \"%s\"\n\n", jsFile)
	code += `var js []byte
var once sync.Once

func JS() []byte {
	once.Do(func() {
		var err error
		js, err = FS.ReadFile(JSFile)
		if err != nil {
			panic(err)
		}
	})
	return js
}
`
	folder := pkg
	if pkg == "babylonjs" {
		folder = "."
		pkg = "babylon"
	}
	err := os.MkdirAll(folder, 0755)
	if err != nil {
		return err
	}
	err = os.WriteFile(filepath.Join(folder, pkg+".go"), []byte(code), 0644)
	if err != nil {
		return err
	}
	return copyFile(jsPath, filepath.Join(folder, jsFile))
}

func copyFile(src, dst string) error {
	println(src, "->", dst)
	fsrc, err := os.Open(src)
	if err != nil {
		return err
	}
	defer fsrc.Close()
	fdst, err := os.Create(dst)
	if err != nil {
		return err
	}
	_, err = io.Copy(fdst, fsrc)
	return err
}

func main() {
	file := "pkgs.txt"
	if len(os.Args) > 1 {
		file = os.Args[1]
	}
	content, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(content), "\n")
	for _, path := range lines {
		if len(path) == 0 {
			continue
		}
		jsFile := filepath.Base(path)
		dir := filepath.Dir(path)
		pkg := strings.ReplaceAll(strings.TrimPrefix(dir, "babylonjs-"), "-", "_")
		pkg = strings.Split(pkg, "/")[0]
		err := makePkg(pkg, dir, jsFile)
		if err != nil {
			panic(err)
		}
	}
}
