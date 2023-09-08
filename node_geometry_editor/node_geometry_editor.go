package node_geometry_editor

import (
	"embed"
)

//go:embed babylon.nodeGeometryEditor.js
var FS embed.FS
var JSData []byte

const JSFile = "babylon.nodeGeometryEditor.js"

func init() {
	var err error
	JSData, err = FS.ReadFile(JSFile)
	if err != nil {
		panic(err)
	}
}
