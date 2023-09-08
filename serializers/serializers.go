package serializers

import (
	"embed"
)

//go:embed babylonjs.serializers.min.js
var FS embed.FS
var JSData []byte

const JSFile = "babylonjs.serializers.min.js"

func init() {
	var err error
	JSData, err = FS.ReadFile(JSFile)
	if err != nil {
		panic(err)
	}
}
