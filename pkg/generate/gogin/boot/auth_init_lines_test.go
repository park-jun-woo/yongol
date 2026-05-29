//ff:func feature=gen-gogin type=test control=sequence
//ff:what authInitLines — resolve 결과를 받아 main.go 에 넣을 라인 슬라이스 조립

package boot

import (
	"strings"
	"testing"
)

func TestAuthInitLines(t *testing.T) {
	cfg := authInitConfig{
		AccessTTL:  "15m",
		RefreshTTL: "168h",
		Mode:       "cookie",
		SameSite:   "Lax",
		SecretEnv:  "JWT_SECRET",
	}
	body := strings.Join(authInitLines(cfg), "\n")
	if !strings.Contains(body, `configureAuth("15m", "168h", "cookie", "Lax", "JWT_SECRET")`) {
		t.Errorf("authInitLines missing configureAuth call with resolved values, got:\n%s", body)
	}
	// Header preamble + store injection should be present.
	if !strings.Contains(body, "auth.Init(infraauth.NewPostgres(queries))") {
		t.Errorf("authInitLines missing RefreshStore install line, got:\n%s", body)
	}
}
