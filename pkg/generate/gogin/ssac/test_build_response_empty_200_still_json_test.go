//ff:func feature=gen-gogin type=test control=sequence topic=response
//ff:what TestBuildResponse_Empty200StillUsesJSONResponse — non-204 빈 응답은 JSONResponse suffix 사용 (regression guard)

package ssac

import (
	"strings"
	"testing"

	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// TestBuildResponse_Empty200StillUsesJSONResponse verifies that non-204
// empty responses still use JSONResponse suffix (regression guard).
func TestBuildResponse_Empty200StillUsesJSONResponse(t *testing.T) {
	g := &methodGen{
		FuncName:      "GetHealth",
		SuccessStatus: 200,
		RespFields:    make(map[string]responseField),
		VarTypes:      map[string]string{},
	}
	seq := ssacparser.Sequence{
		Type: "response",
	}
	lines := g.buildResponse(seq)
	body := strings.Join(lines, "\n")

	want := "return api.GetHealth200JSONResponse{}, nil"
	if !strings.Contains(body, want) {
		t.Fatalf("expected %q, got:\n%s", want, body)
	}
}
