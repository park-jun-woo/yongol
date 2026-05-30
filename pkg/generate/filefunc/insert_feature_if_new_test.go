//ff:func feature=gen-filefunc type=test control=sequence
//ff:what TestInsertFeatureIfNew — 빈 이름/중복 키/신규 삽입 분기 검증

package filefunc

import "testing"

func TestInsertFeatureIfNew(t *testing.T) {
	dst := map[string]string{"existing": "old"}

	insertFeatureIfNew(dst, "", "ignored")     // empty name skipped
	insertFeatureIfNew(dst, "existing", "new") // existing key not overwritten
	insertFeatureIfNew(dst, "fresh", "desc")   // new key inserted

	if _, ok := dst[""]; ok {
		t.Errorf("empty name should not be inserted")
	}
	if got := dst["existing"]; got != "old" {
		t.Errorf("existing key should be preserved, got %q", got)
	}
	if got := dst["fresh"]; got != "desc" {
		t.Errorf("fresh key not inserted, got %q", got)
	}
}
