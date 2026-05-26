//ff:func feature=validate type=rule control=sequence topic=manifest-structural
//ff:what C-9 — backend.lang+framework 조합이 지원되는 BackendType 인지 검증

package manifest

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/generate"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// c09BackendLangFramework validates that the manifest backend.lang and
// backend.framework combination resolves to a known BackendType. An unknown
// pair means yongol has no generator for the requested stack.
func c09BackendLangFramework(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs == nil || fs.Manifest == nil {
		return nil
	}
	lang := fs.Manifest.Backend.Lang
	fw := fs.Manifest.Backend.Framework
	if lang == "" && fw == "" {
		// Both empty — defaults to go+gin at generate time.
		return nil
	}
	if _, err := generate.ResolveBackendType(lang, fw); err != nil {
		return []diagnostic.Diagnostic{{
			File:  "manifest.yaml",
			Phase: diagnostic.PhaseValidate,
			Level: diagnostic.LevelError,
			Message: "[C-9] unsupported backend.lang + backend.framework combination: " +
				lang + " + " + fw,
			Advice: "Supported combinations: go+gin, typescript+nestjs, python+fastapi. " +
				"Set backend.lang and backend.framework to one of these pairs.",
		}}
	}
	return nil
}
