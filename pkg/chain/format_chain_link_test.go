//ff:func feature=chain type=test control=sequence
//ff:what formatChainLink 가 SSOT/artifact 분기 및 line/summary 유무에 따라 포맷하는지 검증
package chain

import (
	"strings"
	"testing"
)

func TestFormatChainLink(t *testing.T) {
	// SSOT link with a line number: location should be File:Line.
	ssot := Link{Kind: "OpenAPI", File: "api/openapi.yaml", Line: 12, Summary: "GET /x"}
	got := formatChainLink(ssot, false)
	if !strings.Contains(got, "OpenAPI") || !strings.Contains(got, "api/openapi.yaml:12") || !strings.Contains(got, "GET /x") {
		t.Errorf("ssot with line: %q", got)
	}

	// SSOT link without a line number: location is just File.
	ssotNoLine := Link{Kind: "SSaC", File: "service/x.ssac", Line: 0, Summary: "@get"}
	got = formatChainLink(ssotNoLine, false)
	if strings.Contains(got, ":0") {
		t.Errorf("ssot no line should not include :0: %q", got)
	}
	if !strings.Contains(got, "service/x.ssac") {
		t.Errorf("ssot no line missing file: %q", got)
	}

	// Artifact link with summary: location should be File:Summary, icon from ownership.
	art := Link{Kind: "Handler", File: "internal/h.go", Summary: "DoThing", Ownership: "preserve"}
	got = formatChainLink(art, true)
	if !strings.Contains(got, "internal/h.go:DoThing") || !strings.Contains(got, "preserve") {
		t.Errorf("artifact with summary: %q", got)
	}

	// Artifact link with "(file)" summary: location stays just the file.
	artFile := Link{Kind: "Model", File: "internal/m.go", Summary: "(file)", Ownership: "gen"}
	got = formatChainLink(artFile, true)
	if strings.Contains(got, "(file)") {
		t.Errorf("artifact (file) summary should be omitted from location: %q", got)
	}
	if !strings.Contains(got, "gen") {
		t.Errorf("artifact missing gen icon: %q", got)
	}
}
