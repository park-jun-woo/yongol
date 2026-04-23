//ff:func feature=gen-gogin type=util control=sequence topic=observability
//ff:what otelExporter — tracing.exporter 값 결정 (미지정 시 "noop")

package boot

import "github.com/park-jun-woo/yongol/pkg/yongol"

// otelExporter resolves the exporter kind with "noop" as the safe default
// when the field is unset. Validation (OBS-003) restricts explicit values
// to otlp/stdout/noop so unknown strings never reach here.
func otelExporter(fs *yongol.Fullstack) string {
	if !hasOtel(fs) {
		return "noop"
	}
	v := fs.Manifest.Backend.Observability.Tracing.Exporter
	if v == "" {
		return "noop"
	}
	return v
}
