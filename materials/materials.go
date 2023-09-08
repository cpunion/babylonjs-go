package materials

import (
	"embed"
)

//go:embed babylonjs.materials.min.js
var FS embed.FS
var JSData []byte

const JSFile = "babylonjs.materials.min.js"

func init() {
	var err error
	JSData, err = FS.ReadFile(JSFile)
	if err != nil {
		panic(err)
	}
}
