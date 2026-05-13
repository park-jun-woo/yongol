//ff:func feature=gen-gogin type=generator control=sequence topic=observability
//ff:what GeneratePrometheus — internal/middleware/prometheus_*.go 기록 (1 file 1 func)

package middleware

import (
	"fmt"
	"path/filepath"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// GeneratePrometheus emits internal/middleware/prometheus_middleware.go and
// prometheus_handler.go — split so each file carries one func (filefunc F1).
// Buckets are sourced from manifest.backend.observability.metrics.buckets
// when set, else fall back to prometheus.DefBuckets.
func GeneratePrometheus(fs *yongol.Fullstack, artifactsDir string) error {
	if fs == nil || fs.Manifest == nil {
		return nil
	}
	if fs.Manifest.Backend.Module == "" {
		return nil
	}
	mwDir := filepath.Join(artifactsDir, "backend", "internal", "middleware")

	buckets := prometheusBuckets(fs)
	files := renderPrometheusSources(buckets)
	if err := writeFiles(mwDir, files); err != nil {
		return fmt.Errorf("write prometheus: %w", err)
	}
	return nil
}
