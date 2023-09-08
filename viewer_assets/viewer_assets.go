package viewer_assets

import (
	"embed"
)

//go:embed babylon.viewer.assets.js
var FS embed.FS
var JSData []byte

const JSFile = "babylon.viewer.assets.js"

func init() {
	var err error
	JSData, err = FS.ReadFile(JSFile)
	if err != nil {
		panic(err)
	}
}
