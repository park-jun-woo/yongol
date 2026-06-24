//ff:func feature=validate type=rule control=sequence dimension=1 topic=stml-openapi
//ff:what TM-59 — manifest refresh_field 선언 시 어떤 STML 액션도 auth.refresh를 캡처하지 않으면 WARNING

package stml_openapi

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func tm59RefreshFieldCapture(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs.Manifest == nil || fs.Manifest.Frontend.Auth == nil {
		return nil
	}
	if fs.Manifest.Frontend.Auth.RefreshField == "" {
		return nil
	}
	if pagesHaveRefreshCapture(fs.STMLPages) {
		return nil
	}
	return []diagnostic.Diagnostic{{
		Phase:   diagnostic.PhaseValidate,
		Level:   diagnostic.LevelWarning,
		Message: "[TM-59] manifest.frontend.auth.refresh_field가 선언되었지만 어떤 STML 액션도 auth.refresh를 캡처하지 않습니다. login에서 refresh token이 저장되지 않아 무음 갱신이 동작하지 않습니다.",
		Advice:  "login data-action의 data-capture에 auth.refresh 바인딩을 추가하세요 (예: refresh_token -> auth.refresh)",
	}}
}
