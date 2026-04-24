//ff:func feature=gen-gogin type=util control=selection
//ff:what resolveGoModDeps — coreDeps + (manifest.tracing 활성 시) OTel + exporter 의존성 병합

package gogin

import "github.com/park-jun-woo/yongol/pkg/yongol"

// resolveGoModDeps returns the final list of `go get` arguments for the
// generated backend. OTel dependencies are appended only when
// manifest.backend.observability.tracing.enabled is true, and the
// exporter-specific module is selected based on the configured exporter.
func resolveGoModDeps(fs *yongol.Fullstack) []string {
	deps := append([]string(nil), coreDeps...)
	tracing := tracingEnabled(fs)
	if tracing == nil {
		return deps
	}
	// Phase005 pgx/v5 refit removed the otelsql DB wrapper; tracing over
	// pgx is deferred to a follow-up via otelpgx. Until then the otelsql
	// module is omitted from deps to keep go.sum clean.
	deps = append(deps,
		"go.opentelemetry.io/otel@"+otelVersion,
		"go.opentelemetry.io/otel/sdk@"+otelVersion,
		"go.opentelemetry.io/otel/trace@"+otelVersion,
		"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin@"+otelContribGinVer,
	)
	switch tracing.Exporter {
	case "otlp", "":
		deps = append(deps, "go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc@"+otelVersion)
	case "stdout":
		deps = append(deps, "go.opentelemetry.io/otel/exporters/stdout/stdouttrace@"+otelVersion)
	}
	return deps
}
