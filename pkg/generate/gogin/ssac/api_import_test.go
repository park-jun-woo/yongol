//ff:func feature=gen-gogin type=test control=sequence
//ff:what apiImportLine — 단일사이트 평문 import vs 도메인 alias import 검증
package ssac

import "testing"

func TestApiImportLine(t *testing.T) {
	if got := apiImportLine("example.com/app", ""); got != `"example.com/app/internal/api"` {
		t.Fatalf("single-site apiImportLine = %q", got)
	}
	if got := apiImportLine("example.com/app", "public"); got != `api "example.com/app/internal/api_public"` {
		t.Fatalf("domain apiImportLine = %q", got)
	}
}
