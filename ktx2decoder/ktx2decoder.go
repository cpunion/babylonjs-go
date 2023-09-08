package ktx2decoder

import (
	"embed"
)

//go:embed babylon.ktx2Decoder.js
var FS embed.FS
var JSData []byte

const JSFile = "babylon.ktx2Decoder.js"

func init() {
	var err error
	JSData, err = FS.ReadFile(JSFile)
	if err != nil {
		panic(err)
	}
}
