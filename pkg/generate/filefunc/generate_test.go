//ff:func feature=gen-filefunc type=test control=iteration dimension=1
//ff:what test: TestGenerate_Zenflow — zenflow fixture 로 codebook.yaml 생성 및 키 존재 검증
package filefunc

import (
	"os"
	"path/filepath"
	"testing"

	"gopkg.in/yaml.v3"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// TestGenerate_Zenflow runs the generator against the zenflow dummy specs,
// writes codebook.yaml into a temp artifacts tree, and asserts that the
// expected required/optional keys are present. Skipped when the dummy
// directory cannot be located.
func TestGenerate_Zenflow(t *testing.T) {
	specsDir := zenflowSpecsDir(t)
	if specsDir == "" {
		t.Skip("zenflow dummy specs not available")
	}
	detected, err := yongol.DetectSSOTs(specsDir)
	if err != nil {
		t.Fatalf("DetectSSOTs: %v", err)
	}
	fs := yongol.ParseAll(specsDir, detected)

	artifacts := t.TempDir()
	backendInternal := filepath.Join(artifacts, "backend", "internal")
	for _, sub := range []string{"api", "middleware", "model", "auth", "db", "workflow", "template", "webhook", "billing"} {
		if err := os.MkdirAll(filepath.Join(backendInternal, sub), 0o755); err != nil {
			t.Fatalf("mkdir %s: %v", sub, err)
		}
	}

	if err := Generate(fs, artifacts); err != nil {
		t.Fatalf("Generate: %v", err)
	}
	path := filepath.Join(artifacts, "backend", "codebook.yaml")
	raw, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read codebook: %v", err)
	}
	var book Codebook
	if err := yaml.Unmarshal(raw, &book); err != nil {
		t.Fatalf("unmarshal codebook: %v", err)
	}

	for _, want := range []string{"workflow", "template", "webhook", "billing", "api", "middleware", "model", "auth"} {
		if _, ok := book.Required.Feature[want]; !ok {
			t.Errorf("missing feature %q", want)
		}
	}
	for _, want := range []string{"handler", "service", "model"} {
		if _, ok := book.Required.Type[want]; !ok {
			t.Errorf("missing type %q", want)
		}
	}
	for _, want := range []string{"request-binding", "response-serialize"} {
		if _, ok := book.Optional.Topic[want]; !ok {
			t.Errorf("missing topic %q", want)
		}
	}
}
