package gui

import (
	"embed"
)

//go:embed babylon.gui.min.js
var FS embed.FS
var JSData []byte

const JSFile = "babylon.gui.min.js"

func init() {
	var err error
	JSData, err = FS.ReadFile(JSFile)
	if err != nil {
		panic(err)
	}
}
