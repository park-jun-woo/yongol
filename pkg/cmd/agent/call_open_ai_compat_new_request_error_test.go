//ff:func feature=agent type=test control=sequence
//ff:what TestCallOpenAICompat — API 키 부재 조기반환 + httptest 로 성공/비200/파싱오류/빈선택/요청오류 분기 검증
package agent

import (
	"strings"
	"testing"
)

func TestCallOpenAICompatNewRequestError(t *testing.T) {
	t.Setenv("XAI_API_KEY", "fake-key")
	// A control character in the URL makes http.NewRequest fail, exercising the
	// "create request: %w" branch before any network activity.
	_, err := callOpenAICompat("http://invalid\x7f.example/", "xai", "grok", "sys", "user")
	if err == nil {
		t.Fatal("expected create-request error for malformed URL")
	}
	if !strings.Contains(err.Error(), "create") {
		t.Errorf("expected create-request error, got: %v", err)
	}
}
