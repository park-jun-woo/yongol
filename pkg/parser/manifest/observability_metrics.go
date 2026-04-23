//ff:type feature=projectconfig type=model
//ff:what ObservabilityMetrics — backend.observability.metrics 모델 (Prometheus 스크랩 설정)

package manifest

// ObservabilityMetrics controls Prometheus metrics emission for the generated
// Go+Gin backend. When the block is absent yongol applies sensible defaults
// (Enabled: true, Path: "/metrics", Buckets: prometheus.DefBuckets).
//
// Enabled is a *bool so the generator can distinguish "unset" (apply default
// true) from "explicitly false" (disable /metrics and the middleware).
//
//   - Enabled: opt-out flag. When nil → true. When set, respected verbatim.
//   - Path:    HTTP route for the scrape endpoint. Must start with "/".
//   - Buckets: histogram buckets (seconds) for http_request_duration_seconds.
//              Empty → prometheus.DefBuckets at codegen time.
//
// Env overrides (resolved in generated main.go):
//   BACKEND_OBSERVABILITY_METRICS_ENABLED
//   BACKEND_OBSERVABILITY_METRICS_PATH
type ObservabilityMetrics struct {
	Enabled *bool     `yaml:"enabled,omitempty"`
	Path    string    `yaml:"path,omitempty"`
	Buckets []float64 `yaml:"buckets,omitempty"`
}
