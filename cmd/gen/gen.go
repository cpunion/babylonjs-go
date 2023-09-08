package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

/* Example:
package babylonjs

import (
	"embed"
)

//go:embed babylon.js
var FS embed.FS
var JS []byte

const JSFile = "babylon.js"

func init() {
	var err error
	JS, err = FS.ReadFile(jsFile)
	if err != nil {
		panic(err)
	}
}
*/

func makePkg(pkg, dir, jsFile string) error {
	println("makePkg:", pkg, dir, jsFile)
	jsPath := filepath.Join("node_modules", dir, jsFile)
	println(jsPath)

	code := fmt.Sprintf("package %s\n\n", pkg)
	code += `import (
	"embed"
)

`
	code += fmt.Sprintf("//go:embed %s\n", jsFile)
	code += "var FS embed.FS\n"
	code += "var JSData []byte\n\n"
	code += fmt.Sprintf("const JSFile = \"%s\"\n", jsFile)
	code += `
func init() {
	var err error
	JSData, err = FS.ReadFile(JSFile)
	if err != nil {
		panic(err)
	}
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

func genTest(pkgs []string) {
	imports := []string{}
	tests := []string{}
	for _, pkg := range pkgs {
		if pkg == "babylonjs" {
			continue
		}
		imports = append(imports, fmt.Sprintf("\t\"github.com/cpunion/babylonjs-go/%s\"", pkg))
		tests = append(tests, fmt.Sprintf("\tfmt.Printf(\"%%s, %%d\\n\", %s.JSFile, len(%s.JSData))", pkg, pkg))
	}
	sort.Strings(imports)
	code := fmt.Sprintf(`package babylonjs

import (
	"fmt"
	"testing"

%s
)

func TestCompile(t *testing.T) {
	fmt.Printf("%%s, %%d\n", JSFile, len(JSData))
%s
}
`, strings.Join(imports, "\n"), strings.Join(tests, "\n"))

	err := os.WriteFile("compile_test.go", []byte(code), 0644)
	if err != nil {
		panic(err)
	}
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
	pkgs := make([]string, 0)
	lines := strings.Split(string(content), "\n")
	for _, path := range lines {
		if len(path) == 0 {
			continue
		}
		jsFile := filepath.Base(path)
		dir := filepath.Dir(path)
		pkg := strings.Split(dir, "/")[0]
		pkg = strings.ReplaceAll(strings.TrimPrefix(pkg, "babylonjs-"), "-", "_")
		pkgs = append(pkgs, pkg)
		err := makePkg(pkg, dir, jsFile)
		if err != nil {
			panic(err)
		}
	}
	genTest(pkgs)
}
