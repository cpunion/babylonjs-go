package accessibility

import (
	"embed"
)

//go:embed babylon.accessibility.js
var FS embed.FS
var JSData []byte

const JSFile = "babylon.accessibility.js"

func init() {
	var err error
	JSData, err = FS.ReadFile(JSFile)
	if err != nil {
		panic(err)
	}
}
