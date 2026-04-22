//ff:func feature=gen-gogin type=generator control=sequence topic=observability
//ff:what GeneratePrometheus — internal/middleware/prometheus.go 기록

package middleware

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// GeneratePrometheus emits internal/middleware/prometheus.go carrying the
// PrometheusMiddleware + PrometheusHandler + 3 default metrics. Buckets are
// sourced from manifest.backend.observability.metrics.buckets when set, else
// fall back to prometheus.DefBuckets (rendered inline by renderPrometheusSource).
//
// When metrics are explicitly disabled (enabled=false) the file is still
// written — the middleware is only wired up in main.go by blockPrometheus,
// so a toggled-off project retains inert helpers without compilation errors.
func GeneratePrometheus(fs *yongol.Fullstack, artifactsDir string) error {
	if fs == nil || fs.Manifest == nil {
		return nil
	}
	if fs.Manifest.Backend.Module == "" {
		return nil
	}
	mwDir := filepath.Join(artifactsDir, "backend", "internal", "middleware")
	if err := os.MkdirAll(mwDir, 0o755); err != nil {
		return fmt.Errorf("mkdir middleware: %w", err)
	}

	buckets := prometheusBuckets(fs)
	path := filepath.Join(mwDir, "prometheus.go")
	if err := os.WriteFile(path, []byte(renderPrometheusSource(buckets)), 0o644); err != nil {
		return fmt.Errorf("write prometheus.go: %w", err)
	}
	return nil
}

// prometheusBuckets resolves the histogram buckets from manifest. Missing
// config returns nil so renderPrometheusSource emits prometheus.DefBuckets.
func prometheusBuckets(fs *yongol.Fullstack) []float64 {
	if fs == nil || fs.Manifest == nil {
		return nil
	}
	obs := fs.Manifest.Backend.Observability
	if obs == nil || obs.Metrics == nil {
		return nil
	}
	return obs.Metrics.Buckets
}
