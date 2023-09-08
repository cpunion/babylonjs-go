package main

import (
	"embed"
	"net/http"
	"time"

	"github.com/cpunion/babylonjs-go/v6/viewer"
	xfs "github.com/qiniu/x/http/fs"
	"github.com/qiniu/x/shell"
)

//go:embed index.html
var HTML embed.FS

func main() {
	babylonFS := xfs.Union(http.FS(HTML), http.FS(viewer.FS))
	go func() {
		time.Sleep(100 * time.Millisecond)
		shell.Open("http://localhost:8080")
	}()
	http.Handle("/", http.FileServer(babylonFS))
	http.ListenAndServe(":8080", nil)
}
