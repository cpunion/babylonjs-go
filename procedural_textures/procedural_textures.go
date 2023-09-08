package procedural_textures

import (
	"embed"
)

//go:embed babylonjs.proceduralTextures.min.js
var FS embed.FS
var JSData []byte

const JSFile = "babylonjs.proceduralTextures.min.js"

func init() {
	var err error
	JSData, err = FS.ReadFile(JSFile)
	if err != nil {
		panic(err)
	}
}
