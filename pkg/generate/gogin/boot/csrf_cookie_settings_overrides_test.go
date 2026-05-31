//ff:func feature=gen-gogin type=test control=sequence topic=csrf
//ff:what csrfCookieSettings — CsrfConfig 기본값 적용 후 생성에 필요한 값들 추출
package boot

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

func TestCsrfCookieSettings_Overrides(t *testing.T) {
	c := &manifest.CsrfConfig{
		CookieName:  "C",
		HeaderName:  "H",
		ExemptPaths: []string{"/a", "/b"},
		MaxAge:      120,
	}
	cookie, header, exempt, maxAge, secure := csrfCookieSettings(c)
	if cookie != "C" || header != "H" {
		t.Errorf("name overrides not applied: %q %q", cookie, header)
	}
	if strings.Join(exempt, ",") != "/a,/b" {
		t.Errorf("exempt paths wrong: %v", exempt)
	}
	if maxAge != 120 {
		t.Errorf("maxAge override = %d, want 120", maxAge)
	}
	if !secure {
		t.Errorf("secure must always be true")
	}
}
