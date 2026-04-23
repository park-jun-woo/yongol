//ff:func feature=validate type=test control=sequence topic=query-structural
//ff:what Q-04 positive — :many 쿼리에 LIMIT 없으면 ERROR

package query

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// TestQ04ManyLimitFires covers Q-04: :many query missing LIMIT fires.
func TestQ04ManyLimitFires(t *testing.T) {
	fs := &yongol.Fullstack{SQLcQueries: loadSpecs(t, "q04_many_no_limit.sql")}
	diags := q04ManyLimit(fs)
	if len(diags) == 0 {
		t.Fatalf("want Q-04 diagnostic, got none")
	}
	if !strings.Contains(diags[0].Message, "[Q-04]") {
		t.Errorf("rule id missing: %q", diags[0].Message)
	}
}
