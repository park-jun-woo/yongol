//ff:func feature=gen-gogin type=util control=sequence topic=observability
//ff:what prometheusBuckets — manifest.backend.observability.metrics.buckets 를 반환

package middleware

import "github.com/park-jun-woo/yongol/pkg/yongol"

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
