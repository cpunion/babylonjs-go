package babylonjs

import (
	"fmt"
	"testing"

	"github.com/cpunion/babylonjs-go/v6/accessibility"
	"github.com/cpunion/babylonjs-go/v6/gui"
	"github.com/cpunion/babylonjs-go/v6/gui_editor"
	"github.com/cpunion/babylonjs-go/v6/inspector"
	"github.com/cpunion/babylonjs-go/v6/ktx2decoder"
	"github.com/cpunion/babylonjs-go/v6/loaders"
	"github.com/cpunion/babylonjs-go/v6/materials"
	"github.com/cpunion/babylonjs-go/v6/node_editor"
	"github.com/cpunion/babylonjs-go/v6/node_geometry_editor"
	"github.com/cpunion/babylonjs-go/v6/post_process"
	"github.com/cpunion/babylonjs-go/v6/procedural_textures"
	"github.com/cpunion/babylonjs-go/v6/serializers"
	"github.com/cpunion/babylonjs-go/v6/viewer"
	"github.com/cpunion/babylonjs-go/v6/viewer_assets"
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
	fmt.Printf("%s, %d\n", accessibility.JSFile, len(accessibility.JSData))
	fmt.Printf("%s, %d\n", ktx2decoder.JSFile, len(ktx2decoder.JSData))
	fmt.Printf("%s, %d\n", node_editor.JSFile, len(node_editor.JSData))
	fmt.Printf("%s, %d\n", node_geometry_editor.JSFile, len(node_geometry_editor.JSData))
	fmt.Printf("%s, %d\n", post_process.JSFile, len(post_process.JSData))
	fmt.Printf("%s, %d\n", procedural_textures.JSFile, len(procedural_textures.JSData))
}
