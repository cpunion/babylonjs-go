package node_editor

import (
	"embed"
)

//go:embed babylon.nodeEditor.js
var FS embed.FS
var JSData []byte

const JSFile = "babylon.nodeEditor.js"

func init() {
	var err error
	JSData, err = FS.ReadFile(JSFile)
	if err != nil {
		panic(err)
	}
}
