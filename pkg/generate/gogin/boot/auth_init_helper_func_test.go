//ff:func feature=gen-gogin type=test control=sequence
//ff:what authInitHelperFunc — cmd/configure_auth.go 에 emit 할 configureAuth 함수 본문 생성

package boot

import (
	"strings"
	"testing"
)

func TestAuthInitHelperFunc(t *testing.T) {
	cfg := authInitConfig{AccessName: "ac", RefreshName: "rf"}
	src := authInitHelperFunc(cfg)
	for _, must := range []string{
		"func configureAuth(accessTTLStr, refreshTTLStr, defaultMode, sameSiteStr, secretEnv string) {",
		"time.ParseDuration(accessTTLStr)",
		"time.ParseDuration(refreshTTLStr)",
		`os.Getenv("BACKEND_AUTH_MODE")`,
		`case "bearer", "cookie", "hybrid":`,
		"parseSameSite(sameSiteStr)",
		"auth.Configure(auth.Config{",
		`AccessName:  "ac",`,
		`RefreshName: "rf",`,
	} {
		if !strings.Contains(src, must) {
			t.Errorf("configureAuth helper missing %q, got:\n%s", must, src)
		}
	}
}
