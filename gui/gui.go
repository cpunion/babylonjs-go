package gui

import (
	"embed"
	"sync"
)

//go:embed babylon.gui.min.js
var FS embed.FS

const JSFile = "babylon.gui.min.js"

var js []byte
var once sync.Once

func JS() []byte {
	once.Do(func() {
		var err error
		js, err = FS.ReadFile(JSFile)
		if err != nil {
			panic(err)
		}
	})
	return js
}
