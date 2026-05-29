//ff:func feature=gen-gogin type=test control=selection topic=observability
//ff:what blockOtelInit — OpenTelemetry TracerProvider 초기화 (exporter/sampler/shutdown)

package boot

import (
	"strings"
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

func TestBlockOtelInit_OtlpExporter(t *testing.T) {
	fs := fsTracing(&pmanifest.ObservabilityTracing{Enabled: true, Exporter: "otlp", OtlpEndpoint: "collector:4317"})
	block := blockOtelInit(fs, "example.com/zenflow")
	body := strings.Join(block.Lines, "\n")
	imp := strings.Join(block.Imports, "\n")
	for _, must := range []string{
		`envString("BACKEND_OBSERVABILITY_TRACING_EXPORTER", "otlp")`,
		`case "otlp":`,
		`envString("BACKEND_OBSERVABILITY_TRACING_OTLP_ENDPOINT", "collector:4317")`,
		"sdktrace.NewTracerProvider(tpOpts...)",
	} {
		if !strings.Contains(body, must) {
			t.Errorf("otlp otel block missing %q, got:\n%s", must, body)
		}
	}
	if !strings.Contains(imp, "otlptracegrpc") {
		t.Errorf("otlp block must import otlp exporter, got:\n%s", imp)
	}
	if !strings.Contains(imp, `"time"`) {
		t.Errorf("otel block must import time for shutdown timeout, got:\n%s", imp)
	}
}

func TestBlockOtelInit_StdoutExporter(t *testing.T) {
	fs := fsTracing(&pmanifest.ObservabilityTracing{Enabled: true, Exporter: "stdout"})
	block := blockOtelInit(fs, "example.com/zenflow")
	if !strings.Contains(strings.Join(block.Lines, "\n"), `case "stdout":`) {
		t.Errorf("stdout exporter branch missing, got:\n%v", block.Lines)
	}
	if !strings.Contains(strings.Join(block.Imports, "\n"), "stdouttrace") {
		t.Errorf("stdout block must import stdouttrace, got:\n%v", block.Imports)
	}
}
