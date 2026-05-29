//ff:func feature=validate type=test control=sequence topic=ssac-ddl
//ff:what isFuncManagedTable — Flags["funcManaged.<table>"] true/false 및 nil Ground 조회

package ssac_ddl

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/rule"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestIsFuncManagedTable(t *testing.T) {
	fs := &yongol.Fullstack{}
	fs.SetGround(&rule.Ground{Flags: rule.StringSet{"funcManaged.bids": true}})

	if !isFuncManagedTable(fs, "bids") {
		t.Errorf("isFuncManagedTable(bids) = false, want true")
	}
	if isFuncManagedTable(fs, "courses") {
		t.Errorf("isFuncManagedTable(courses) = true, want false")
	}

	// nil Ground → false (no panic).
	nilFS := &yongol.Fullstack{}
	if isFuncManagedTable(nilFS, "bids") {
		t.Errorf("isFuncManagedTable with nil Ground = true, want false")
	}
}
