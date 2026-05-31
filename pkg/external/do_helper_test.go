//ff:func feature=external type=test control=iteration dimension=1
//ff:what TestDoHelper — do() 헬퍼 코드 생성 검증 (수신자 타입·HTTP 흐름 포함)
package external

import (
	"strings"
	"testing"
)

func TestDoHelper(t *testing.T) {
	code := doHelper("EscrowClient")

	for _, want := range []string{
		"func (c *EscrowClient) do(ctx context.Context, method, path string, body any, result any) error {",
		"json.NewEncoder(&buf).Encode(body)",
		"http.NewRequestWithContext(ctx, method, c.baseURL+path, &buf)",
		`req.Header.Set("Content-Type", "application/json")`,
		"resp, err := c.client.Do(req)",
		"if resp.StatusCode >= 400 {",
	} {
		if !strings.Contains(code, want) {
			t.Errorf("doHelper output missing %q", want)
		}
	}
	// %%w / %%d escapes in the template must render as single % in output.
	if strings.Contains(code, "%%") {
		t.Error("doHelper output still contains unescaped %% sequences")
	}
}
