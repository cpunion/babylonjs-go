package inspector

import (
	"embed"
)

//go:embed babylon.inspector.bundle.js
var FS embed.FS
var JSData []byte

const JSFile = "babylon.inspector.bundle.js"

func init() {
	var err error
	JSData, err = FS.ReadFile(JSFile)
	if err != nil {
		panic(err)
	}
}
