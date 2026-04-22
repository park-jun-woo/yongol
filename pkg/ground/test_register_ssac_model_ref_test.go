//ff:func feature=rule type=test control=sequence dimension=1
//ff:what registerSSaCModelRef — Model 이름 + pluralized DDL table 등록

package ground

import "testing"

// TestRegisterSSaCModelRef_Plural verifies both the Go model name and the
// lowercase plural (= DDL table) are inserted.
func TestRegisterSSaCModelRef_Plural(t *testing.T) {
	set := map[string]bool{}
	registerSSaCModelRef("Course.FindByID", set)
	if !set["Course"] {
		t.Errorf("Course missing: %v", set)
	}
	if !set["courses"] {
		t.Errorf("plural 'courses' missing: %v", set)
	}
}

// TestRegisterSSaCModelRef_NoDot — no dot → no insert.
func TestRegisterSSaCModelRef_NoDot(t *testing.T) {
	set := map[string]bool{}
	registerSSaCModelRef("NoDot", set)
	if len(set) != 0 {
		t.Errorf("expected empty set for no-dot input, got %v", set)
	}
}
