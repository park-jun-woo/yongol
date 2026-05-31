//ff:func feature=agent type=test control=sequence
//ff:what TestLoadAPIKey — 환경변수 우선, XDG credentials.yaml fallback, 미존재 시 에러 검증
package agent

import (
	"testing"
)

func TestLoadAPIKeyHomeError(t *testing.T) {
	// Both XDG_CONFIG_HOME and HOME empty makes os.UserHomeDir fail on unix,
	// exercising the "load API key" error branch.
	t.Setenv("XAI_API_KEY", "")
	t.Setenv("XDG_CONFIG_HOME", "")
	t.Setenv("HOME", "")
	if _, err := loadAPIKey("xai"); err == nil {
		t.Skip("UserHomeDir did not fail on this platform with empty HOME")
	}
}
