//ff:func feature=agent type=test control=sequence
//ff:what TestScaffoldByName_ZeroCov — scaffoldOpenAPI / scaffoldSSaCFeature 직접 호출 (LLM 미사용 분기)
package agent

import (
	"bytes"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func TestScaffoldSSaCFeature_LLMError_ZeroCov(t *testing.T) {
	// New file + unsupported backend → LLM call errors out.
	feat := features.Feature{Op: "Login", Path: "/auth/login"}
	var out bytes.Buffer
	cfg := Config{Backend: "unsupported-backend", Model: "m"}
	if _, err := scaffoldSSaCFeature(t.TempDir(), feat, "", "", cfg, &out); err == nil {
		t.Fatal("expected LLM error from scaffoldSSaCFeature")
	}
}
