//ff:func feature=gen-gogin type=test control=sequence topic=csrf
//ff:what csrfCookieSettings — CsrfConfig 기본값 적용 후 생성에 필요한 값들 추출
package boot

import (
	"testing"
)

func TestCsrfCookieSettings_Defaults(t *testing.T) {
	cookie, header, exempt, maxAge, secure := csrfCookieSettings(nil)
	if cookie != "XSRF-TOKEN" || header != "X-XSRF-TOKEN" {
		t.Errorf("default names wrong: %q %q", cookie, header)
	}
	if len(exempt) != 0 {
		t.Errorf("default exempt should be empty, got %v", exempt)
	}
	if maxAge != 86400 || !secure {
		t.Errorf("default maxAge/secure wrong: %d %v", maxAge, secure)
	}
}
