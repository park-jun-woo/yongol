//ff:func feature=agent type=test control=sequence
//ff:what TestWriteFeatureTableContext — 테이블 관련 feature 목록 기록, nil/무관련 시 무기록 검증
package agent

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func TestWriteFeatureTableContext(t *testing.T) {
	ff := &features.FeaturesFile{
		Features: []features.Feature{
			{Op: "CreateUser", Path: "/users", Desc: "create", Table: "users"},
			{Op: "ListOrgs", Path: "/orgs", Desc: "list", Table: "orgs"},
		},
	}

	var b strings.Builder
	writeFeatureTableContext(&b, ff, "users")
	out := b.String()
	if !strings.Contains(out, "Related features:") || !strings.Contains(out, "CreateUser /users: create") {
		t.Errorf("related → %q", out)
	}
	if strings.Contains(out, "ListOrgs") {
		t.Errorf("unrelated feature leaked: %q", out)
	}

	// No matching table: nothing written.
	var b2 strings.Builder
	writeFeatureTableContext(&b2, ff, "nomatch")
	if b2.Len() != 0 {
		t.Errorf("no-match wrote %q", b2.String())
	}

	// nil ff: nothing.
	var b3 strings.Builder
	writeFeatureTableContext(&b3, nil, "users")
	if b3.Len() != 0 {
		t.Errorf("nil ff wrote %q", b3.String())
	}
}
