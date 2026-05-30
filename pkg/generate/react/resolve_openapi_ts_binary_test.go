//ff:func feature=gen-react type=test control=sequence
//ff:what resolveOpenapiTsBinary env→PATH→npx 해결 순서·부재 에러 검증

package react

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

// isolatePATH replaces PATH with an empty directory so exec.LookPath cannot
// resolve openapi-typescript or npx, and clears the project-dir env vars.
func isolatePATH(t *testing.T) {
	t.Helper()
	empty := t.TempDir()
	t.Setenv("PATH", empty)
	t.Setenv("YONGOL_OPENAPI_TS_PROJECT_DIR", "")
	t.Setenv("YONGOL_SWC_PROJECT_DIR", "")
}

func TestResolveOpenapiTsBinaryFromProjectDir(t *testing.T) {
	isolatePATH(t)

	projDir := t.TempDir()
	binDir := filepath.Join(projDir, "node_modules", ".bin")
	if err := os.MkdirAll(binDir, 0o755); err != nil {
		t.Fatal(err)
	}
	binPath := filepath.Join(binDir, "openapi-typescript")
	if err := os.WriteFile(binPath, []byte("#!/bin/sh\n"), 0o755); err != nil {
		t.Fatal(err)
	}
	t.Setenv("YONGOL_OPENAPI_TS_PROJECT_DIR", projDir)

	argv, env, err := resolveOpenapiTsBinary()
	if err != nil {
		t.Fatal(err)
	}
	if len(argv) != 1 || argv[0] != binPath {
		t.Errorf("expected argv [%s], got %v", binPath, argv)
	}
	if env != nil {
		t.Errorf("expected nil env, got %v", env)
	}
}

func TestResolveOpenapiTsBinaryFromSwcDir(t *testing.T) {
	isolatePATH(t)

	swcDir := t.TempDir()
	binDir := filepath.Join(swcDir, "node_modules", ".bin")
	if err := os.MkdirAll(binDir, 0o755); err != nil {
		t.Fatal(err)
	}
	binPath := filepath.Join(binDir, "openapi-typescript")
	if err := os.WriteFile(binPath, []byte("#!/bin/sh\n"), 0o755); err != nil {
		t.Fatal(err)
	}
	t.Setenv("YONGOL_SWC_PROJECT_DIR", swcDir)

	argv, _, err := resolveOpenapiTsBinary()
	if err != nil {
		t.Fatal(err)
	}
	if len(argv) != 1 || argv[0] != binPath {
		t.Errorf("expected argv [%s], got %v", binPath, argv)
	}
}

func TestResolveOpenapiTsBinaryFromPATH(t *testing.T) {
	isolatePATH(t)

	// Place an executable named openapi-typescript on PATH.
	pathDir := t.TempDir()
	binPath := filepath.Join(pathDir, "openapi-typescript")
	if err := os.WriteFile(binPath, []byte("#!/bin/sh\n"), 0o755); err != nil {
		t.Fatal(err)
	}
	t.Setenv("PATH", pathDir)

	argv, _, err := resolveOpenapiTsBinary()
	if err != nil {
		t.Fatal(err)
	}
	if len(argv) != 1 || argv[0] != binPath {
		t.Errorf("expected resolved PATH binary [%s], got %v", binPath, argv)
	}
}

func TestResolveOpenapiTsBinaryFromNpx(t *testing.T) {
	isolatePATH(t)

	// Only npx is available on PATH.
	pathDir := t.TempDir()
	npxPath := filepath.Join(pathDir, "npx")
	if err := os.WriteFile(npxPath, []byte("#!/bin/sh\n"), 0o755); err != nil {
		t.Fatal(err)
	}
	t.Setenv("PATH", pathDir)

	argv, _, err := resolveOpenapiTsBinary()
	if err != nil {
		t.Fatal(err)
	}
	want := []string{"npx", "--yes", "openapi-typescript"}
	if len(argv) != len(want) {
		t.Fatalf("expected %v, got %v", want, argv)
	}
	for i := range want {
		if argv[i] != want[i] {
			t.Errorf("argv[%d] = %q, want %q", i, argv[i], want[i])
		}
	}
}

func TestResolveOpenapiTsBinaryNotFound(t *testing.T) {
	isolatePATH(t)

	// Guard: ensure neither binary is reachable in the isolated PATH.
	if _, err := exec.LookPath("npx"); err == nil {
		t.Skip("npx unexpectedly resolvable in isolated PATH")
	}

	argv, env, err := resolveOpenapiTsBinary()
	if err == nil {
		t.Fatalf("expected error when no binary available, got argv=%v env=%v", argv, env)
	}
}
