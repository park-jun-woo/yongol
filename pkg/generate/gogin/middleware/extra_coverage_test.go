//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestMiddlewareExtras — csrf/prometheus/rate-limit/request-id/writeFiles 유틸 검증

package middleware

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/prepared"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestCsrfActive(t *testing.T) {
	cases := []struct {
		name string
		auth prepared.Auth
		want bool
	}{
		{"not-required", prepared.Auth{CsrfRequired: false}, false},
		{"required-but-absent", prepared.Auth{CsrfRequired: true, Present: false}, false},
		{"required-nil-raw", prepared.Auth{CsrfRequired: true, Present: true, Raw: nil}, false},
		{"required-nil-csrf-defaults-on", prepared.Auth{CsrfRequired: true, Present: true, Raw: &pmanifest.Auth{}}, true},
		{"required-csrf-enabled", prepared.Auth{CsrfRequired: true, Present: true, Raw: &pmanifest.Auth{Csrf: &pmanifest.CsrfConfig{Enabled: true}}}, true},
		{"required-csrf-disabled", prepared.Auth{CsrfRequired: true, Present: true, Raw: &pmanifest.Auth{Csrf: &pmanifest.CsrfConfig{Enabled: false}}}, false},
	}
	for _, c := range cases {
		if got := csrfActive(c.auth); got != c.want {
			t.Errorf("%s: csrfActive = %v, want %v", c.name, got, c.want)
		}
	}
}

func TestGenerateCsrf(t *testing.T) {
	t.Run("SkipsWhenInactive", func(t *testing.T) {
		arts := t.TempDir()
		if err := GenerateCsrf(prepared.Auth{CsrfRequired: false}, arts); err != nil {
			t.Fatalf("error: %v", err)
		}
		if _, err := os.Stat(filepath.Join(arts, "backend")); !os.IsNotExist(err) {
			t.Errorf("expected no output when csrf inactive")
		}
	})

	t.Run("WritesWhenActive", func(t *testing.T) {
		arts := t.TempDir()
		a := prepared.Auth{CsrfRequired: true, Present: true, Raw: &pmanifest.Auth{}}
		if err := GenerateCsrf(a, arts); err != nil {
			t.Fatalf("error: %v", err)
		}
		if _, err := os.Stat(filepath.Join(arts, "backend", "internal", "middleware", "csrf.go")); err != nil {
			t.Errorf("expected csrf.go: %v", err)
		}
	})

	t.Run("MkdirError", func(t *testing.T) {
		arts := t.TempDir()
		internal := filepath.Join(arts, "backend", "internal")
		if err := os.MkdirAll(filepath.Dir(internal), 0o755); err != nil {
			t.Fatalf("setup: %v", err)
		}
		if err := os.WriteFile(internal, []byte("x"), 0o644); err != nil {
			t.Fatalf("setup: %v", err)
		}
		a := prepared.Auth{CsrfRequired: true, Present: true, Raw: &pmanifest.Auth{}}
		if err := GenerateCsrf(a, arts); err == nil {
			t.Errorf("expected mkdir error, got nil")
		}
	})

	t.Run("WriteError", func(t *testing.T) {
		arts := t.TempDir()
		// pre-create csrf.go as a directory so WriteFile fails after mkdir.
		mwDir := filepath.Join(arts, "backend", "internal", "middleware")
		if err := os.MkdirAll(filepath.Join(mwDir, "csrf.go"), 0o755); err != nil {
			t.Fatalf("setup: %v", err)
		}
		a := prepared.Auth{CsrfRequired: true, Present: true, Raw: &pmanifest.Auth{}}
		if err := GenerateCsrf(a, arts); err == nil {
			t.Errorf("expected write csrf.go error, got nil")
		}
	})
}

// makeMiddlewareDirCollision sets up arts so writeFiles' MkdirAll on the
// middleware directory fails (its parent "internal" is a regular file).
func makeMiddlewareDirCollision(t *testing.T) string {
	t.Helper()
	arts := t.TempDir()
	internal := filepath.Join(arts, "backend", "internal")
	if err := os.MkdirAll(filepath.Dir(internal), 0o755); err != nil {
		t.Fatalf("setup: %v", err)
	}
	if err := os.WriteFile(internal, []byte("x"), 0o644); err != nil {
		t.Fatalf("setup: %v", err)
	}
	return arts
}

func TestGeneratePrometheusWriteError(t *testing.T) {
	arts := makeMiddlewareDirCollision(t)
	fs := &yongol.Fullstack{Manifest: &pmanifest.ProjectConfig{Backend: pmanifest.Backend{Module: "example.com/app"}}}
	if err := GeneratePrometheus(fs, arts); err == nil {
		t.Errorf("expected write prometheus error, got nil")
	}
}

func TestGenerateRateLimitWriteError(t *testing.T) {
	arts := makeMiddlewareDirCollision(t)
	fs := &yongol.Fullstack{Manifest: &pmanifest.ProjectConfig{Backend: pmanifest.Backend{Module: "example.com/app"}}}
	if err := GenerateRateLimit(fs, arts); err == nil {
		t.Errorf("expected write rate_limit error, got nil")
	}
}

func TestPrometheusBuckets(t *testing.T) {
	if got := prometheusBuckets(nil); got != nil {
		t.Errorf("nil fs: want nil, got %v", got)
	}
	if got := prometheusBuckets(&yongol.Fullstack{Manifest: &pmanifest.ProjectConfig{}}); got != nil {
		t.Errorf("no observability: want nil, got %v", got)
	}
	fs := &yongol.Fullstack{Manifest: &pmanifest.ProjectConfig{
		Backend: pmanifest.Backend{Observability: &pmanifest.Observability{
			Metrics: &pmanifest.ObservabilityMetrics{Buckets: []float64{0.1, 1}},
		}},
	}}
	if got := prometheusBuckets(fs); len(got) != 2 || got[0] != 0.1 {
		t.Errorf("want [0.1 1], got %v", got)
	}
}

func TestGeneratePrometheus(t *testing.T) {
	t.Run("SkipsNilManifest", func(t *testing.T) {
		if err := GeneratePrometheus(&yongol.Fullstack{}, t.TempDir()); err != nil {
			t.Errorf("error: %v", err)
		}
	})
	t.Run("SkipsEmptyModule", func(t *testing.T) {
		fs := &yongol.Fullstack{Manifest: &pmanifest.ProjectConfig{}}
		if err := GeneratePrometheus(fs, t.TempDir()); err != nil {
			t.Errorf("error: %v", err)
		}
	})
	t.Run("WritesFiles", func(t *testing.T) {
		arts := t.TempDir()
		fs := &yongol.Fullstack{Manifest: &pmanifest.ProjectConfig{Backend: pmanifest.Backend{Module: "example.com/app"}}}
		if err := GeneratePrometheus(fs, arts); err != nil {
			t.Fatalf("error: %v", err)
		}
		if _, err := os.Stat(filepath.Join(arts, "backend", "internal", "middleware", "prometheus_middleware.go")); err != nil {
			t.Errorf("expected prometheus_middleware.go: %v", err)
		}
	})
}

func TestGenerateRateLimit(t *testing.T) {
	t.Run("SkipsNilManifest", func(t *testing.T) {
		if err := GenerateRateLimit(&yongol.Fullstack{}, t.TempDir()); err != nil {
			t.Errorf("error: %v", err)
		}
	})
	t.Run("SkipsEmptyModule", func(t *testing.T) {
		fs := &yongol.Fullstack{Manifest: &pmanifest.ProjectConfig{}}
		if err := GenerateRateLimit(fs, t.TempDir()); err != nil {
			t.Errorf("error: %v", err)
		}
	})
	t.Run("WritesFiles", func(t *testing.T) {
		arts := t.TempDir()
		fs := &yongol.Fullstack{Manifest: &pmanifest.ProjectConfig{Backend: pmanifest.Backend{Module: "example.com/app"}}}
		if err := GenerateRateLimit(fs, arts); err != nil {
			t.Fatalf("error: %v", err)
		}
		if _, err := os.Stat(filepath.Join(arts, "backend", "internal", "middleware", "fixed_rate_limit.go")); err != nil {
			t.Errorf("expected fixed_rate_limit.go: %v", err)
		}
	})
}

func TestGenerateRequestID(t *testing.T) {
	t.Run("WritesFiles", func(t *testing.T) {
		arts := t.TempDir()
		if err := GenerateRequestID(arts); err != nil {
			t.Fatalf("error: %v", err)
		}
		if _, err := os.Stat(filepath.Join(arts, "backend", "internal", "middleware", "request_id.go")); err != nil {
			t.Errorf("expected request_id.go: %v", err)
		}
	})
	t.Run("WriteError", func(t *testing.T) {
		arts := t.TempDir()
		internal := filepath.Join(arts, "backend", "internal")
		if err := os.MkdirAll(filepath.Dir(internal), 0o755); err != nil {
			t.Fatalf("setup: %v", err)
		}
		if err := os.WriteFile(internal, []byte("x"), 0o644); err != nil {
			t.Fatalf("setup: %v", err)
		}
		if err := GenerateRequestID(arts); err == nil {
			t.Errorf("expected write error, got nil")
		}
	})
}

func TestRenderRateLimitSources(t *testing.T) {
	m := renderRateLimitSources()
	for _, k := range []string{"fixed_rate_limit.go", "fixed_rate_limit_key.go", "route_rate_limit.go", "route_rate_limit_key.go"} {
		if _, ok := m[k]; !ok {
			t.Errorf("missing rate-limit source %q", k)
		}
	}
}

func TestRemoveStaleCombined(t *testing.T) {
	t.Run("RemovesAndIgnoresMissing", func(t *testing.T) {
		dir := t.TempDir()
		// create one stale file; the others are missing (ignored).
		if err := os.WriteFile(filepath.Join(dir, "prometheus.go"), []byte("x"), 0o644); err != nil {
			t.Fatalf("setup: %v", err)
		}
		if err := removeStaleCombined(dir); err != nil {
			t.Fatalf("error: %v", err)
		}
		if _, err := os.Stat(filepath.Join(dir, "prometheus.go")); !os.IsNotExist(err) {
			t.Errorf("expected prometheus.go removed")
		}
	})

	t.Run("RemoveError", func(t *testing.T) {
		dir := t.TempDir()
		if err := os.WriteFile(filepath.Join(dir, "prometheus.go"), []byte("x"), 0o644); err != nil {
			t.Fatalf("setup: %v", err)
		}
		if err := os.Chmod(dir, 0o555); err != nil {
			t.Fatalf("setup: %v", err)
		}
		defer os.Chmod(dir, 0o755)
		if err := removeStaleCombined(dir); err == nil {
			t.Skip("Remove did not fail (likely root)")
		}
	})
}

func TestWriteFiles(t *testing.T) {
	t.Run("WritesGoAndPlain", func(t *testing.T) {
		dir := filepath.Join(t.TempDir(), "out")
		files := map[string]string{
			"a.go":  "package x\n\nfunc A() {}\n",
			"b.txt": "plain",
		}
		if err := writeFiles(dir, files); err != nil {
			t.Fatalf("error: %v", err)
		}
		if _, err := os.Stat(filepath.Join(dir, "a.go")); err != nil {
			t.Errorf("expected a.go: %v", err)
		}
		if _, err := os.Stat(filepath.Join(dir, "b.txt")); err != nil {
			t.Errorf("expected b.txt: %v", err)
		}
	})

	t.Run("MkdirError", func(t *testing.T) {
		base := t.TempDir()
		fp := filepath.Join(base, "file")
		if err := os.WriteFile(fp, []byte("x"), 0o644); err != nil {
			t.Fatalf("setup: %v", err)
		}
		if err := writeFiles(filepath.Join(fp, "sub"), map[string]string{"a.go": "package x"}); err == nil {
			t.Errorf("expected mkdir error, got nil")
		}
	})

	t.Run("WriteFileError", func(t *testing.T) {
		dir := t.TempDir()
		// target file name pre-exists as a directory -> WriteIfNotPreserved fails.
		if err := os.MkdirAll(filepath.Join(dir, "a.go"), 0o755); err != nil {
			t.Fatalf("setup: %v", err)
		}
		if err := writeFiles(dir, map[string]string{"a.go": "package x\n\nfunc A() {}\n"}); err == nil {
			t.Errorf("expected write error, got nil")
		}
	})
}

func TestWriteValidator(t *testing.T) {
	t.Run("Writes", func(t *testing.T) {
		dir := t.TempDir()
		if err := writeValidator(dir); err != nil {
			t.Fatalf("error: %v", err)
		}
		if _, err := os.Stat(filepath.Join(dir, "request_validator.go")); err != nil {
			t.Errorf("expected request_validator.go: %v", err)
		}
	})
	t.Run("WriteError", func(t *testing.T) {
		dir := t.TempDir()
		// target is a directory -> WriteFile fails.
		if err := os.MkdirAll(filepath.Join(dir, "request_validator.go"), 0o755); err != nil {
			t.Fatalf("setup: %v", err)
		}
		if err := writeValidator(dir); err == nil {
			t.Errorf("expected write error, got nil")
		}
	})
}
