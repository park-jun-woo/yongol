//ff:func feature=validate type=test control=sequence topic=design-manifest
//ff:what TestXnv01_Skip_NilManifest — manifest nil 시 skip

package design_manifest

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXnv01_Skip_NilManifest(t *testing.T) {
	fs := &yongol.Fullstack{}
	diags := xnv01PathExists(fs)
	if len(diags) != 0 {
		t.Fatalf("want 0 diags (nil manifest), got %+v", diags)
	}
}
