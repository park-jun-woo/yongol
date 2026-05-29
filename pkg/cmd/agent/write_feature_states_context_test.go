//ff:func feature=agent type=test control=sequence
//ff:what TestWriteFeatureStatesContext — 테이블 states 기록, nil/빈 테이블/states 부재 시 무기록 검증

package agent

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func TestWriteFeatureStatesContext(t *testing.T) {
	ff := &features.FeaturesFile{
		Tables: map[string]features.TableDef{
			"order": {States: []string{"draft", "active"}},
			"empty": {States: nil},
		},
	}

	var b strings.Builder
	writeFeatureStatesContext(&b, ff, "order")
	if got := b.String(); !strings.Contains(got, "States for order: draft, active") {
		t.Errorf("states → %q", got)
	}

	// nil features file: nothing written.
	var b2 strings.Builder
	writeFeatureStatesContext(&b2, nil, "order")
	if b2.Len() != 0 {
		t.Errorf("nil ff wrote %q", b2.String())
	}

	// Empty table name: nothing.
	var b3 strings.Builder
	writeFeatureStatesContext(&b3, ff, "")
	if b3.Len() != 0 {
		t.Errorf("empty table wrote %q", b3.String())
	}

	// Table with no states: nothing.
	var b4 strings.Builder
	writeFeatureStatesContext(&b4, ff, "empty")
	if b4.Len() != 0 {
		t.Errorf("no-states table wrote %q", b4.String())
	}
}
