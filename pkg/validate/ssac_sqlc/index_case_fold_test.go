//ff:func feature=validate type=test control=sequence topic=ssac-sqlc
//ff:what TestSSaCSqlcHelpers — unit tests for the pure ssac_sqlc helper functions
package ssac_sqlc

import (
	"testing"
)

func TestIndexCaseFold(t *testing.T) {
	if got := indexCaseFold("id as user_id", " AS "); got != 2 {
		t.Errorf("indexCaseFold lowercase AS = %d, want 2", got)
	}
	if got := indexCaseFold("id AS x", " as "); got != 2 {
		t.Errorf("indexCaseFold mixed = %d, want 2", got)
	}
	if got := indexCaseFold("nothing", "zzz"); got != -1 {
		t.Errorf("indexCaseFold no match = %d, want -1", got)
	}
}
