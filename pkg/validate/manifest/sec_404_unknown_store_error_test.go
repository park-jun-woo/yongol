//ff:func feature=validate type=test control=sequence topic=manifest-auth
//ff:what SEC-404 테스트 — 오타 store 값은 ERROR

package manifest

import (
	"strings"
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestSec404_UnknownStore_Error(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{
			Frontend: pmanifest.Frontend{
				Auth: &pmanifest.FrontendAuth{TokenField: "access_token", Store: "localstorage"}, // wrong casing
			},
		},
	}
	got := sec404FrontendAuthStoreEnum(fs)
	if len(got) != 1 {
		t.Fatalf("expected 1 diagnostic for typo store, got %d: %+v", len(got), got)
	}
	if !strings.Contains(got[0].Message, "[SEC-404]") {
		t.Fatalf("message missing [SEC-404] prefix: %q", got[0].Message)
	}
}
