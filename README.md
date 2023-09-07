# babylonjs-go

Provides embed files for [BabylonJS](https://babylonjs.com).

## Usage

See [example](example/main.go).

```go
package main

import (
	"embed"
	"net/http"
	"time"

	"github.com/cpunion/babylonjs-go/viewer"
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
```

index.html:

```html
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <title>Example</title>
    <style>
      html,
      body {
        width: 100%;
        height: 100%;
        margin: 0;
        padding: 0;
        overflow: hidden;
      }
    </style>
    <script src="babylon.viewer.js"></script>
  </head>
  <body>
    <babylon model="https://playground.babylonjs.com/scenes/dummy3.babylon" />
  </body>
</html>
```
