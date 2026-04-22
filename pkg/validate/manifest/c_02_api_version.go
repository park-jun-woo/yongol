//ff:func feature=validate type=rule control=sequence topic=manifest-structural
//ff:what C-2 — manifest apiVersion 값 검증

package manifest

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// c02APIVersion validates that manifest.apiVersion is "yongol/v1".
func c02APIVersion(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs.Manifest == nil {
		return nil
	}
	if fs.Manifest.APIVersion == "yongol/v1" {
		return nil
	}
	return []diagnostic.Diagnostic{{
		File:    "manifest.yaml",
		Phase:   diagnostic.PhaseValidate,
		Level:   diagnostic.LevelError,
		Message: "[C-2] apiVersion " + quoted(fs.Manifest.APIVersion) + " is not \"yongol/v1\"",
		Advice:  "apiVersion 을 yongol/v1 로 설정하세요",
	}}
}

