//ff:func feature=agent type=test control=sequence
//ff:what TestWriteFeaturePublicContext — feature public 속성을 컨텍스트에 기록 검증

package agent

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func TestWriteFeaturePublicContext(t *testing.T) {
	var b strings.Builder
	writeFeaturePublicContext(&b, features.Feature{Public: true})
	if got := b.String(); !strings.Contains(got, "Feature public: true") {
		t.Errorf("public=true → %q", got)
	}

	var b2 strings.Builder
	writeFeaturePublicContext(&b2, features.Feature{Public: false})
	if got := b2.String(); !strings.Contains(got, "Feature public: false") {
		t.Errorf("public=false → %q", got)
	}
}
