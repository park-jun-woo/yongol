//ff:func feature=gen-gogin type=test control=sequence topic=observability
//ff:what blockOtelInit — OpenTelemetry TracerProvider 초기화 (exporter/sampler/shutdown)
package boot

import (
	"strings"
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

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
