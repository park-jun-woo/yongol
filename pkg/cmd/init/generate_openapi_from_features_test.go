//ff:func feature=cli-init type=test control=iteration dimension=1
//ff:what TestGenerateOpenAPIFromFeatures — path 그룹핑 stub 생성 / 잘못된 path 에러 / write 에러 분기 검증

package cliinit

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func TestGenerateOpenAPIFromFeatures_Success(t *testing.T) {
	target := t.TempDir()
	if err := os.MkdirAll(filepath.Join(target, "specs", "api"), 0o755); err != nil {
		t.Fatal(err)
	}
	data := templateData{ProjectID: "App", Description: "desc"}
	feats := []features.Feature{
		{Op: "ListTasks", Path: "GET /tasks"},
		{Op: "CreateTask", Path: "POST /tasks"}, // same URI -> grouped
		{Op: "GetTask", Path: "GET /tasks/{id}"},
	}
	if err := generateOpenAPIFromFeatures(target, data, feats); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	got, err := os.ReadFile(filepath.Join(target, "specs", "api", "openapi.yaml"))
	if err != nil {
		t.Fatal(err)
	}
	s := string(got)
	if !strings.Contains(s, "openapi: 3.0.3") || !strings.Contains(s, "title: App") {
		t.Errorf("header missing: %q", s)
	}
	if !strings.Contains(s, "operationId: ListTasks") || !strings.Contains(s, "operationId: CreateTask") {
		t.Errorf("operationIds missing: %q", s)
	}
	// /tasks should appear once (grouped) with both get and post under it.
	if strings.Count(s, "  /tasks:\n") != 1 {
		t.Errorf("expected /tasks grouped once: %q", s)
	}
}

func TestGenerateOpenAPIFromFeatures_BadPath(t *testing.T) {
	target := t.TempDir()
	if err := os.MkdirAll(filepath.Join(target, "specs", "api"), 0o755); err != nil {
		t.Fatal(err)
	}
	feats := []features.Feature{{Op: "X", Path: "GARBAGE"}}
	if err := generateOpenAPIFromFeatures(target, templateData{}, feats); err == nil {
		t.Fatal("want error for invalid path")
	}
}

func TestGenerateOpenAPIFromFeatures_WriteError(t *testing.T) {
	target := t.TempDir() // no specs/api -> write fails
	feats := []features.Feature{{Op: "X", Path: "GET /x"}}
	if err := generateOpenAPIFromFeatures(target, templateData{}, feats); err == nil {
		t.Fatal("want write error when dest dir missing")
	}
}
