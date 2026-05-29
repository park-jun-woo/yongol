//ff:func feature=gen-gogin type=test control=sequence
//ff:what authInitConfigureLines — auth.Configure(auth.Config{...}) 블록 라인 생성

package boot

import (
	"strings"
	"testing"
)

func TestAuthInitConfigureLines(t *testing.T) {
	cfg := authInitConfig{
		SecretEnv:   "MY_SECRET",
		AccessName:  "ac",
		RefreshName: "rf",
	}
	body := strings.Join(authInitConfigureLines(cfg), "\n")
	for _, must := range []string{
		"auth.Configure(auth.Config{",
		`SecretEnv:  "MY_SECRET",`,
		"AccessTTL:  accessTTL,",
		"Mode:       authMode,",
		`AccessName:  "ac",`,
		`RefreshName: "rf",`,
		"SameSite:    sameSite,",
	} {
		if !strings.Contains(body, must) {
			t.Errorf("authInitConfigureLines missing %q, got:\n%s", must, body)
		}
	}
}
