package loaders

import (
	"embed"
)

//go:embed babylonjs.loaders.min.js
var FS embed.FS
var JSData []byte

const JSFile = "babylonjs.loaders.min.js"

func init() {
	var err error
	JSData, err = FS.ReadFile(JSFile)
	if err != nil {
		panic(err)
	}
}
