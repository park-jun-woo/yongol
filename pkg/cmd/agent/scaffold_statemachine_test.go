//ff:func feature=agent type=test control=iteration dimension=1
//ff:what TestScaffoldStateMachine — states 없음 0 / 기존파일 skip / mkdir 에러 / 미지원 backend LLM 에러 분기 검증

package agent

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func TestScaffoldStateMachineNoTargets(t *testing.T) {
	var out bytes.Buffer
	// Tables without States produce no targets -> count 0, nil error.
	ff := &features.FeaturesFile{Tables: map[string]features.TableDef{"users": {}}}
	n, err := scaffoldStateMachine(t.TempDir(), ff, nil, Config{}, &out)
	if err != nil {
		t.Fatalf("no targets: unexpected error: %v", err)
	}
	if n != 0 {
		t.Fatalf("no targets: expected 0, got %d", n)
	}
}

func TestScaffoldStateMachineSkipExisting(t *testing.T) {
	dir := t.TempDir()
	statesDir := filepath.Join(dir, "states")
	if err := os.MkdirAll(statesDir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(statesDir, "orders.md"), []byte("stateDiagram\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	ff := &features.FeaturesFile{Tables: map[string]features.TableDef{
		"orders": {States: []string{"pending", "shipped"}},
	}}
	var out bytes.Buffer
	n, err := scaffoldStateMachine(dir, ff, nil, Config{}, &out)
	if err != nil {
		t.Fatalf("skip-existing: unexpected error: %v", err)
	}
	if n != 0 {
		t.Fatalf("skip-existing: expected count 0, got %d", n)
	}
	if !strings.Contains(out.String(), "skipped") {
		t.Errorf("expected skip message, got: %q", out.String())
	}
}

func TestScaffoldStateMachineMkdirError(t *testing.T) {
	// Creating "states" as a regular file makes os.MkdirAll(statesDir) fail.
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "states"), []byte("x"), 0o644); err != nil {
		t.Fatal(err)
	}
	ff := &features.FeaturesFile{Tables: map[string]features.TableDef{
		"orders": {States: []string{"pending"}},
	}}
	var out bytes.Buffer
	if _, err := scaffoldStateMachine(dir, ff, nil, Config{}, &out); err == nil {
		t.Fatal("expected mkdir error when states is a file")
	}
}

func TestScaffoldStateMachineLLMError(t *testing.T) {
	// Missing states file + unsupported backend -> target's LLM call fails, propagated.
	ff := &features.FeaturesFile{Tables: map[string]features.TableDef{
		"orders": {States: []string{"pending"}},
	}}
	var out bytes.Buffer
	cfg := Config{Backend: "unsupported-backend", Model: "m"}
	if _, err := scaffoldStateMachine(t.TempDir(), ff, nil, cfg, &out); err == nil {
		t.Fatal("expected LLM error from scaffoldStateMachineTarget")
	}
}

func TestScaffoldStateMachineTargetSkipExisting(t *testing.T) {
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "orders.md"), []byte("x"), 0o644); err != nil {
		t.Fatal(err)
	}
	var out bytes.Buffer
	created, err := scaffoldStateMachineTarget(dir, "orders", []string{"pending"}, nil, "sys", Config{}, &out)
	if err != nil {
		t.Fatalf("skip-existing: unexpected error: %v", err)
	}
	if created {
		t.Fatal("skip-existing: expected created=false")
	}
	if !strings.Contains(out.String(), "skipped") {
		t.Errorf("expected skip message, got: %q", out.String())
	}
}

func TestScaffoldStateMachineTargetLLMError(t *testing.T) {
	var out bytes.Buffer
	cfg := Config{Backend: "unsupported-backend", Model: "m"}
	created, err := scaffoldStateMachineTarget(t.TempDir(), "orders", []string{"pending"}, nil, "sys", cfg, &out)
	if err == nil {
		t.Fatal("expected LLM error")
	}
	if created {
		t.Fatal("expected created=false on error")
	}
}
