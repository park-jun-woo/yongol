//ff:func feature=gen-gogin type=test control=sequence topic=observability
//ff:what TestBuildCall_WrapCallsOff — wrap_calls=false 일 때 otel import/Tracer 호출이 없어야 함

package ssac

import (
	"strings"
	"testing"

	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// TestBuildCall_WrapCallsOff confirms the default shape: no otel.Tracer
// invocation, no otel import. This matches the pre-Phase009 contract so
// projects that leave wrap_calls off have identical output.
func TestBuildCall_WrapCallsOff(t *testing.T) {
	g := &methodGen{
		FuncName:   "CreateUser",
		FileName:   "user_service.ssac",
		ModulePath: "example.com/zenflow",
		WrapCalls:  false,
	}
	seq := ssacparser.Sequence{
		Type:  "call",
		Model: "mail.SendEmail",
		Inputs: map[string]string{
			"To": "user.Email",
		},
	}
	lines, imports := g.buildCall(seq)
	body := strings.Join(lines, "\n")
	imp := strings.Join(imports, "\n")

	if strings.Contains(body, "otel.Tracer") {
		t.Fatalf("wrap_calls=false must NOT emit otel.Tracer, got:\n%s", body)
	}
	if strings.Contains(imp, "go.opentelemetry.io/otel") {
		t.Fatalf("wrap_calls=false must NOT import otel, got:\n%s", imp)
	}
	if !strings.Contains(body, "mail.SendEmail(ctx, mail.SendEmailRequest{") {
		t.Fatalf("expected plain ctx-first call, got:\n%s", body)
	}
}
