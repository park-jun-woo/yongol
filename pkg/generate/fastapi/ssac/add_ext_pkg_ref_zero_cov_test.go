//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestFastapiSsacHelpers_ZeroCov — addExtPkgRef/appendSnakeRune 커버
package ssac

import (
	"testing"
)

func TestAddExtPkgRef_ZeroCov(t *testing.T) {
	d := &importData{ExtPkgs: map[string]map[string]bool{}}
	addExtPkgRef(d, "mail", "Send")
	addExtPkgRef(d, "mail", "Queue") // existing pkg map reused
	if !d.ExtPkgs["mail"]["Send"] || !d.ExtPkgs["mail"]["Queue"] {
		t.Fatalf("addExtPkgRef = %v", d.ExtPkgs)
	}
}
