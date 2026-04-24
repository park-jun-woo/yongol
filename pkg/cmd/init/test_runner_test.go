//ff:type feature=cli-init type=test
//ff:what test — Run writes the expected skeleton tree and refuses to clobber

package cliinit

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRunCreatesSkeleton(t *testing.T) {
	tmp := t.TempDir()
	target := filepath.Join(tmp, "myapp")
	var outBuf, errBuf bytes.Buffer
	err := Run(&outBuf, &errBuf, Options{
		ProjectID:   "Myapp",
		Description: "Test description",
		Dir:         target,
		Module:      "github.com/test/myapp",
	})
	if err != nil {
		t.Fatalf("Run returned error: %v", err)
	}

	expectFiles := []string{
		"specs/manifest.yaml",
		"specs/api/openapi.yaml",
		"specs/db/sqlc.yaml",
		"specs/policy/authz.rego",
		"README.md",
		".gitignore",
	}
	for _, rel := range expectFiles {
		path := filepath.Join(target, rel)
		info, err := os.Stat(path)
		if err != nil {
			t.Errorf("expected file %s missing: %v", rel, err)
			continue
		}
		if info.Size() == 0 {
			t.Errorf("expected file %s is empty", rel)
		}
	}

	expectDirs := []string{
		"specs/db/queries",
		"specs/service",
		"specs/states",
		"specs/frontend/pages",
		"specs/frontend/components",
		"specs/tests",
	}
	for _, rel := range expectDirs {
		info, err := os.Stat(filepath.Join(target, rel))
		if err != nil || !info.IsDir() {
			t.Errorf("expected directory %s missing", rel)
		}
	}

	manifest, err := os.ReadFile(filepath.Join(target, "specs/manifest.yaml"))
	if err != nil {
		t.Fatalf("read manifest: %v", err)
	}
	body := string(manifest)
	if !strings.Contains(body, "apiVersion: yongol/v1") {
		t.Errorf("manifest missing yongol/v1 apiVersion:\n%s", body)
	}
	if !strings.Contains(body, "name: myapp") {
		t.Errorf("manifest missing normalized name=myapp:\n%s", body)
	}
	if !strings.Contains(body, "module: github.com/test/myapp") {
		t.Errorf("manifest missing module override:\n%s", body)
	}

	openapi, err := os.ReadFile(filepath.Join(target, "specs/api/openapi.yaml"))
	if err != nil {
		t.Fatalf("read openapi: %v", err)
	}
	if !strings.Contains(string(openapi), "paths: {}") {
		t.Errorf("openapi missing empty paths placeholder")
	}

	sqlc, err := os.ReadFile(filepath.Join(target, "specs/db/sqlc.yaml"))
	if err != nil {
		t.Fatalf("read sqlc.yaml: %v", err)
	}
	if !strings.Contains(string(sqlc), "sql_package: pgx/v5") {
		t.Errorf("sqlc.yaml missing pgx/v5 directive")
	}
}

func TestRunRefusesNonEmptyDir(t *testing.T) {
	tmp := t.TempDir()
	// Pre-populate the target so Run must refuse.
	if err := os.WriteFile(filepath.Join(tmp, "existing.txt"), []byte("x"), 0o644); err != nil {
		t.Fatal(err)
	}
	var outBuf, errBuf bytes.Buffer
	err := Run(&outBuf, &errBuf, Options{
		ProjectID:   "Myapp",
		Description: "Test",
		Dir:         tmp,
		Module:      "github.com/test/myapp",
	})
	if err == nil {
		t.Fatalf("Run should refuse non-empty dir without --force")
	}
	if !strings.Contains(err.Error(), "not empty") {
		t.Errorf("error message does not mention emptiness: %v", err)
	}
}

func TestRunForceOverridesNonEmptyDir(t *testing.T) {
	tmp := t.TempDir()
	if err := os.WriteFile(filepath.Join(tmp, "existing.txt"), []byte("x"), 0o644); err != nil {
		t.Fatal(err)
	}
	var outBuf, errBuf bytes.Buffer
	err := Run(&outBuf, &errBuf, Options{
		ProjectID:   "Myapp",
		Description: "Test",
		Dir:         tmp,
		Module:      "github.com/test/myapp",
		Force:       true,
	})
	if err != nil {
		t.Fatalf("Run with --force unexpectedly failed: %v", err)
	}
	if _, err := os.Stat(filepath.Join(tmp, "specs", "manifest.yaml")); err != nil {
		t.Errorf("manifest missing after --force run: %v", err)
	}
}

func TestRunRejectsBadProjectID(t *testing.T) {
	tmp := t.TempDir()
	var outBuf, errBuf bytes.Buffer
	err := Run(&outBuf, &errBuf, Options{
		ProjectID:   "my-app", // hyphen disallowed
		Description: "Test",
		Dir:         tmp,
		Module:      "github.com/test/myapp",
	})
	if err == nil {
		t.Fatalf("Run should reject ProjectID with hyphen")
	}
}
