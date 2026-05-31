//ff:func feature=gen-gogin type=test control=sequence
//ff:what buildCallRefreshCreateLines 단위 테스트 (auth.CreateRefresh + UserClaim 리터럴 방출)
package ssac

import (
	"strings"
	"testing"

	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestBuildCallRefreshCreateLines(t *testing.T) {
	g := &methodGen{FuncName: "Login"}
	seq := ssacparser.Sequence{
		Inputs: map[string]string{
			"UserID": "user.ID",
			"Role":   `"member"`,
		},
	}
	lines := g.buildCallRefreshCreateLines(seq, "refreshed")
	joined := strings.Join(lines, "\n")

	if !strings.Contains(joined, "auth.CreateRefresh(ctx, refreshed.RefreshToken,") {
		t.Errorf("missing CreateRefresh call:\n%s", joined)
	}
	// mapFields sorts keys: Role before UserID.
	if !strings.Contains(joined, `model.UserClaim{Role: "member", UserID: user.ID}`) {
		t.Errorf("unexpected UserClaim literal:\n%s", joined)
	}
	if !strings.Contains(joined, "api.Login500JSONResponse") {
		t.Errorf("missing 500 error response:\n%s", joined)
	}
	if lines[len(lines)-1] != "}" {
		t.Errorf("expected closing brace, got %q", lines[len(lines)-1])
	}
}
