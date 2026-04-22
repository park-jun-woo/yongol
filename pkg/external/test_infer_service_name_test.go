//ff:func feature=external type=generator control=iteration dimension=1
//ff:what TestInferServiceName: 파일명·타이틀에서 서비스명 추론 테이블 드리븐 검증
package external

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestInferServiceName(t *testing.T) {
	tests := []struct {
		source string
		title  string
		want   string
	}{
		{"escrow.openapi.yaml", "", "Escrow"},
		{"my-service.openapi.yaml", "", "MyService"},
		{"stripe.yaml", "", "Stripe"},
		{"anything.yaml", "Payment Gateway", "PaymentGateway"},
		{"https://api.example.com/openapi.yaml", "", "Openapi"}, // fallback to filename
	}

	for _, tt := range tests {
		doc := &openapi3.T{}
		if tt.title != "" {
			doc.Info = &openapi3.Info{Title: tt.title}
		}
		got := inferServiceName(tt.source, doc)
		if got != tt.want {
			t.Errorf("inferServiceName(%q, title=%q) = %q, want %q", tt.source, tt.title, got, tt.want)
		}
	}
}
