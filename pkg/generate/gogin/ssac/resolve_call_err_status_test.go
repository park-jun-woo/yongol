//ff:func feature=gen-gogin type=test control=iteration dimension=2
//ff:what resolveCallErrStatus 단위 테스트 (seq explicit > project > builtin > 500)

package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/funcspec"
)

func TestResolveCallErrStatus(t *testing.T) {
	project := []funcspec.FuncSpec{
		{Package: "dashboard", Name: "summarize", ErrStatus: 422},
	}
	builtin := []funcspec.FuncSpec{
		{Package: "mail", Name: "send", ErrStatus: 502},
	}
	cases := []struct {
		name      string
		seqStatus int
		pkg       string
		fn        string
		want      int
	}{
		{"seq explicit wins", 409, "dashboard", "Summarize", 409},
		{"project funcspec match", 0, "dashboard", "Summarize", 422},
		{"builtin funcspec match", 0, "mail", "Send", 502},
		{"no match falls to 500", 0, "unknown", "Nope", 500},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := resolveCallErrStatus(tc.seqStatus, tc.pkg, tc.fn, project, builtin)
			if got != tc.want {
				t.Errorf("resolveCallErrStatus = %d, want %d", got, tc.want)
			}
		})
	}
}
