package viewer

import (
	"embed"
)

//go:embed babylon.viewer.js
var FS embed.FS
var JSData []byte

const JSFile = "babylon.viewer.js"

func init() {
	var err error
	JSData, err = FS.ReadFile(JSFile)
	if err != nil {
		panic(err)
	}
}
