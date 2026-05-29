//ff:func feature=chain type=test control=iteration dimension=2
//ff:what Print 가 빈 링크 / SSOT+artifact 분리 출력을 올바르게 수행하는지 검증
package chain

import (
	"bytes"
	"strings"
	"testing"
)

func TestPrint(t *testing.T) {
	// Empty links → "No SSOT links found." message.
	var buf bytes.Buffer
	Print(&buf, "Op", nil)
	out := buf.String()
	if !strings.Contains(out, "Feature Chain: Op") {
		t.Errorf("missing header: %q", out)
	}
	if !strings.Contains(out, "No SSOT links found.") {
		t.Errorf("empty links missing message: %q", out)
	}

	// Mixed SSOT and artifact links → artifacts under an "Artifacts" section.
	buf.Reset()
	links := []Link{
		{Kind: "OpenAPI", File: "api/openapi.yaml", Line: 3, Summary: "GET /x"},
		{Kind: "Handler", File: "internal/h.go", Summary: "DoThing", Ownership: "preserve"},
	}
	Print(&buf, "Op", links)
	out = buf.String()
	if !strings.Contains(out, "api/openapi.yaml") {
		t.Errorf("missing ssot link: %q", out)
	}
	if !strings.Contains(out, "Artifacts") {
		t.Errorf("missing artifacts header: %q", out)
	}
	if !strings.Contains(out, "internal/h.go") {
		t.Errorf("missing artifact link: %q", out)
	}

	// SSOT-only links → no Artifacts section.
	buf.Reset()
	Print(&buf, "Op", []Link{{Kind: "SSaC", File: "service/x.ssac", Line: 1, Summary: "@get"}})
	if strings.Contains(buf.String(), "Artifacts") {
		t.Errorf("ssot-only should not print Artifacts: %q", buf.String())
	}
}
