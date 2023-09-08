package babylonjs

import (
	"fmt"
	"testing"

	"github.com/cpunion/babylonjs-go/gui"
	"github.com/cpunion/babylonjs-go/gui_editor"
	"github.com/cpunion/babylonjs-go/inspector"
	"github.com/cpunion/babylonjs-go/loaders"
	"github.com/cpunion/babylonjs-go/materials"
	"github.com/cpunion/babylonjs-go/serializers"
	"github.com/cpunion/babylonjs-go/viewer"
	"github.com/cpunion/babylonjs-go/viewer_assets"
)

func TestCompile(t *testing.T) {
	fmt.Printf("%s, %d\n", JSFile, len(JSData))
	fmt.Printf("%s, %d\n", viewer.JSFile, len(viewer.JSData))
	fmt.Printf("%s, %d\n", loaders.JSFile, len(loaders.JSData))
	fmt.Printf("%s, %d\n", materials.JSFile, len(materials.JSData))
	fmt.Printf("%s, %d\n", serializers.JSFile, len(serializers.JSData))
	fmt.Printf("%s, %d\n", inspector.JSFile, len(inspector.JSData))
	fmt.Printf("%s, %d\n", gui.JSFile, len(gui.JSData))
	fmt.Printf("%s, %d\n", gui_editor.JSFile, len(gui_editor.JSData))
	fmt.Printf("%s, %d\n", viewer_assets.JSFile, len(viewer_assets.JSData))
}
