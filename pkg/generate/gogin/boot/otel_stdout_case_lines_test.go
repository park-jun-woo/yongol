//ff:func feature=gen-gogin type=test control=iteration dimension=1 topic=observability
//ff:what otelStdoutCaseLines — switch "stdout" case 의 본문 라인 생성
package boot

import (
	"strings"
	"testing"
)

func TestOtelStdoutCaseLines(t *testing.T) {
	body := strings.Join(otelStdoutCaseLines(), "\n")
	for _, must := range []string{
		`case "stdout":`,
		"stdouttrace.New(stdouttrace.WithPrettyPrint())",
		"spanExporter = exp",
	} {
		if !strings.Contains(body, must) {
			t.Errorf("stdout branch missing %q, got:\n%s", must, body)
		}
	}
}
