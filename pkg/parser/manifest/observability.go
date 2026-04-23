//ff:type feature=projectconfig type=model
//ff:what Observability — backend.observability 섹션 모델 (metrics + tracing 집계점)

package manifest

// Observability mirrors the backend.observability: section of manifest.yaml.
// Phase008 introduced metrics; Phase009 adds tracing under the same namespace
// so the block stays a single logical observability root.
type Observability struct {
	Metrics *ObservabilityMetrics `yaml:"metrics,omitempty"`
	Tracing *ObservabilityTracing `yaml:"tracing,omitempty"`
}
