package babylonjs

import (
	"embed"
)

//go:embed babylon.js
var FS embed.FS
var JSData []byte

const JSFile = "babylon.js"

func init() {
	var err error
	JSData, err = FS.ReadFile(JSFile)
	if err != nil {
		panic(err)
	}
}
