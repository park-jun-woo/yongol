//ff:func feature=cli-init type=test control=sequence
//ff:what TestRenderTemplate — embed 읽기 에러 / 정상 렌더(데이터 치환) 분기 검증
package cliinit

import (
	"strings"
	"testing"
)

func TestRenderTemplate_Success(t *testing.T) {
	data := templateData{ProjectID: "MyApp", ProjectIDNormalized: "myapp", Description: "desc"}
	out, err := renderTemplate("templates/manifest.yaml.tmpl", data)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(out) == 0 {
		t.Fatal("expected rendered output")
	}
	// The manifest template substitutes ProjectID somewhere in the output.
	if !strings.Contains(string(out), "MyApp") && !strings.Contains(string(out), "myapp") {
		t.Errorf("expected ProjectID substitution in output: %q", out)
	}
}
