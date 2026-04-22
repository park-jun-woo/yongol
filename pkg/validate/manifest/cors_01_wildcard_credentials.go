//ff:func feature=validate type=rule control=sequence topic=manifest-cors
//ff:what CORS-01 — allow_origins=["*"] + allow_credentials=true 동시 허용 금지

package manifest

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// cors01WildcardCredentials enforces the browser fetch spec: when
// allow_origins contains only the wildcard "*", allow_credentials must be
// false. Real browsers reject responses that combine them, so this guard
// prevents silent production breakage.
func cors01WildcardCredentials(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs.Manifest == nil || fs.Manifest.Backend.CORS == nil {
		return nil
	}
	c := fs.Manifest.Backend.CORS
	if !c.Enabled {
		return nil
	}
	if !c.AllowCredentials {
		return nil
	}
	if !corsAllowOriginsHasWildcard(c.AllowOrigins) {
		return nil
	}
	return []diagnostic.Diagnostic{{
		File:    "manifest.yaml",
		Phase:   diagnostic.PhaseValidate,
		Level:   diagnostic.LevelError,
		Message: "[CORS-01] allow_origins=\"*\" 와 allow_credentials=true 는 동시에 사용할 수 없습니다",
		Advice:  "allow_origins 를 명시적 origin 목록으로 바꾸거나 allow_credentials 를 false 로 설정하세요",
	}}
}
