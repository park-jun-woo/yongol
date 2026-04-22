//ff:func feature=validate type=rule control=sequence topic=manifest-security-headers
//ff:what SEC-302 — HSTS max_age 가 180일(15552000초) 미만이면 WARNING

package manifest

import (
	"strconv"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// hstsPreloadMinSeconds is the de-facto minimum max-age accepted by major
// browser HSTS preload lists (180 days). Phase007 plan §D references this
// threshold.
const hstsPreloadMinSeconds = 15552000

// sec302HSTSShort warns when backend.security_headers.hsts.max_age is set
// to a value below the HSTS preload minimum (180 days). A short max-age
// narrows the window during which browsers remember to force HTTPS, leaving
// users vulnerable to SSL-strip attacks on revisits. The rule only fires
// when HSTS is explicitly configured — omitted blocks inherit the default
// (31536000, 1 year) which satisfies the threshold.
func sec302HSTSShort(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs == nil || fs.Manifest == nil {
		return nil
	}
	sh := fs.Manifest.Backend.SecurityHeaders
	if sh == nil || sh.HSTS == nil {
		return nil
	}
	if sh.HSTS.MaxAge <= 0 {
		return nil
	}
	if sh.HSTS.MaxAge >= hstsPreloadMinSeconds {
		return nil
	}
	return []diagnostic.Diagnostic{{
		File:    "manifest.yaml",
		Phase:   diagnostic.PhaseValidate,
		Level:   diagnostic.LevelWarning,
		Message: "[SEC-302] backend.security_headers.hsts.max_age=" + strconv.Itoa(sh.HSTS.MaxAge) + " 는 HSTS preload 최소치(" + strconv.Itoa(hstsPreloadMinSeconds) + "초, 180일) 미만입니다",
		Advice:  "운영 환경에서는 max_age 를 최소 15552000 (180일) 이상, 권장 31536000 (1년) 으로 설정하세요",
	}}
}
