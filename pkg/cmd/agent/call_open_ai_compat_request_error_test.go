//ff:func feature=agent type=test control=sequence
//ff:what TestCallOpenAICompat — API 키 부재 조기반환 + httptest 로 성공/비200/파싱오류/빈선택/요청오류 분기 검증
package agent

import (
	"testing"
)

func TestCallOpenAICompatRequestError(t *testing.T) {
	t.Setenv("XAI_API_KEY", "fake-key")
	// A syntactically valid but unroutable scheme triggers http.DefaultClient.Do
	// to fail, exercising the "request: %w" branch.
	_, err := callOpenAICompat("http://127.0.0.1:0/x", "xai", "grok", "sys", "user")
	if err == nil {
		t.Fatal("expected request error for unreachable endpoint")
	}
}
