//ff:func feature=ssac-parse type=parser control=sequence
//ff:what import에서 net/http 제외 확인 — 표준 라이브러리 필터링

package ssac

import "testing"

func TestParseImportsExcludeNetHTTP(t *testing.T) {
	src := `package service

import (
	"net/http"
	"myapp/billing"
)

// @get User user = User.FindByID({UserID: request.UserID})
func GetUser(c *gin.Context) {}
`
	sfs := parseTestFile(t, src)
	if len(sfs[0].Imports) != 1 {
		t.Fatalf("expected 1 import, got %d", len(sfs[0].Imports))
	}
	assertEqual(t, "Import", sfs[0].Imports[0], "myapp/billing")
}
