//ff:func feature=stml-gen type=test control=iteration dimension=1
//ff:what zodChainFor — 정상 경로 통과 (지원 타입이 zodChain 과 동일 체인 반환) 검증

package stml

import (
	"testing"

	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
)

// A supported field type returns the same chain as zodChain, without panicking.
func TestZodChainFor_Supported(t *testing.T) {
	cases := []struct {
		name string
		fc   oapiparser.FieldConstraint
		want string
	}{
		{
			name: "required string",
			fc:   oapiparser.FieldConstraint{Type: "string", Required: true},
			want: "z.string().min(1)",
		},
		{
			name: "optional integer",
			fc:   oapiparser.FieldConstraint{Type: "integer"},
			want: "z.number().int().optional()",
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := zodChainFor("CreateThing", tc.name, tc.fc)
			if got != tc.want {
				t.Errorf("zodChainFor(...) = %q, want %q", got, tc.want)
			}
			// Must match the unwrapped zodChain result.
			if base := zodChain(tc.fc); got != base {
				t.Errorf("zodChainFor != zodChain: %q vs %q", got, base)
			}
		})
	}
}
