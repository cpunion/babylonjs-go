package gui_editor

import (
	"embed"
)

//go:embed babylon.guiEditor.js
var FS embed.FS
var JSData []byte

const JSFile = "babylon.guiEditor.js"

func init() {
	var err error
	JSData, err = FS.ReadFile(JSFile)
	if err != nil {
		panic(err)
	}
}
