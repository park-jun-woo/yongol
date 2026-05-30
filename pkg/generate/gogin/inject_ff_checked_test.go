//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestInjectFFChecked — backend Go 파일 //ff:checked 주입 + skip/error 경로 검증

package gogin

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

func TestInjectFFChecked(t *testing.T) {
	annotated := []byte("/" + "/ff:func feature=demo type=handler control=sequence\n" +
		"/" + "/ff:what Demo — demo function\n" +
		"package demo\n\n" +
		"func Demo(x int) int {\n\treturn x + 1\n}\n")

	t.Run("InjectsAndSkips", func(t *testing.T) {
		specs := t.TempDir()
		backend := t.TempDir()
		// func spec dir "billing" -> internal/billing is skipped.
		if err := os.MkdirAll(filepath.Join(specs, "func", "billing"), 0o755); err != nil {
			t.Fatalf("setup: %v", err)
		}
		// owned file under internal/api -> injected.
		apiDir := filepath.Join(backend, "internal", "api")
		if err := os.MkdirAll(apiDir, 0o755); err != nil {
			t.Fatalf("setup: %v", err)
		}
		owned := filepath.Join(apiDir, "server.go")
		mustWrite(t, owned, string(annotated))
		// skipped func-spec mirror file under internal/billing.
		billingDir := filepath.Join(backend, "internal", "billing")
		if err := os.MkdirAll(billingDir, 0o755); err != nil {
			t.Fatalf("setup: %v", err)
		}
		skipped := filepath.Join(billingDir, "deduct.go")
		mustWrite(t, skipped, string(annotated))

		if err := injectFFChecked(backend, specs); err != nil {
			t.Fatalf("injectFFChecked error: %v", err)
		}
		if got, _ := os.ReadFile(owned); !bytes.Contains(got, []byte("/"+"/ff:checked")) {
			t.Errorf("owned file should be injected:\n%s", got)
		}
		if got, _ := os.ReadFile(skipped); bytes.Contains(got, []byte("/"+"/ff:checked")) {
			t.Errorf("func-spec mirror should be skipped:\n%s", got)
		}
	})

	t.Run("FuncSpecRelPathsError", func(t *testing.T) {
		// specs/func is a regular file -> funcSpecRelPaths returns an error.
		specs := t.TempDir()
		if err := os.WriteFile(filepath.Join(specs, "func"), []byte("x"), 0o644); err != nil {
			t.Fatalf("setup: %v", err)
		}
		if err := injectFFChecked(t.TempDir(), specs); err == nil {
			t.Errorf("expected funcSpecRelPaths error, got nil")
		}
	})
}
