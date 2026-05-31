//ff:func feature=gen-gogin type=test control=sequence
//ff:what buildCallSetAuthCookiesLines 단위 테스트 (AccessTokenVar + refresh var로 SetAuthCookies)
package ssac

import (
	"strings"
	"testing"
)

func TestBuildCallSetAuthCookiesLines(t *testing.T) {
	g := &methodGen{AccessTokenVar: "issued"}
	lines := g.buildCallSetAuthCookiesLines("refreshed")
	joined := strings.Join(lines, "\n")
	if !strings.Contains(joined, "if ginCtx, ok := ctx.(*gin.Context); ok {") {
		t.Errorf("missing gin.Context assertion:\n%s", joined)
	}
	if !strings.Contains(joined, "auth.SetAuthCookies(ginCtx, issued.AccessToken, refreshed.RefreshToken)") {
		t.Errorf("missing SetAuthCookies with access+refresh vars:\n%s", joined)
	}
	if lines[len(lines)-1] != "}" {
		t.Errorf("expected closing brace, got %q", lines[len(lines)-1])
	}
}
