//ff:func feature=gen-fastapi type=test control=iteration dimension=1
//ff:what RenderEnvExample test — .env.example 기본 설정 라인 전체 검증
package scaffold

import (
	"strings"
	"testing"
)

func TestRenderEnvExample(t *testing.T) {
	got, err := RenderEnvExample()
	if err != nil {
		t.Fatalf("RenderEnvExample() error = %v", err)
	}
	for _, want := range []string{
		"DATABASE_URL=postgresql+asyncpg://postgres:postgres@localhost:5432/app",
		"JWT_SECRET=change-me-in-production",
		"JWT_ALGORITHM=HS256",
		"JWT_EXPIRE_MINUTES=30",
		"DEBUG=false",
	} {
		if !strings.Contains(got, want) {
			t.Errorf("RenderEnvExample() missing %q\ngot:\n%s", want, got)
		}
	}
}
