//ff:func feature=validate type=test control=iteration dimension=1 topic=query-structural
//ff:what Q-08 positive — 선언되었으나 본문에 없는 파라미터는 ERROR

package query

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// TestQ08UnusedParamFires covers Q-08: declared param missing from body → ERROR.
func TestQ08UnusedParamFires(t *testing.T) {
	specs := loadSpecs(t, "q08_unused_param.sql")
	// Inject a synthetic unused parameter. The parser only records params
	// actually present in the SQL, so we augment Params here to exercise Q-08.
	for i := range specs {
		specs[i].Params = append(specs[i].Params, "GhostParam")
	}
	fs := &yongol.Fullstack{SQLcQueries: specs}
	diags := q08UnusedParam(fs)
	if len(diags) == 0 {
		t.Fatalf("want Q-08 diagnostic, got none")
	}
	if !strings.Contains(diags[0].Message, "[Q-08]") {
		t.Errorf("rule id missing: %q", diags[0].Message)
	}
}
