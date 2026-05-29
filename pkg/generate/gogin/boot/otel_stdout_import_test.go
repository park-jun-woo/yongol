//ff:func feature=gen-gogin type=test control=sequence topic=observability
//ff:what otelStdoutImport — stdouttrace exporter import 라인

package boot

import (
	"strings"
	"testing"
)

func TestOtelStdoutImport(t *testing.T) {
	imp := otelStdoutImport()
	if !strings.Contains(imp, "stdout/stdouttrace") {
		t.Errorf("stdout import should target stdouttrace, got %q", imp)
	}
}
