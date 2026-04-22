//ff:func feature=validate type=rule control=iteration dimension=1 topic=funcspec-structural
//ff:what F-1 — 프로젝트 func 가 yongol 내장 같은 함수명을 재정의하는지 감지 (패키지 이름 중복 아님)

package funcspec

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// f01BuiltinOverride flags project FuncSpecs that override a built-in pkg
// function of the same `<pkg>.<name>` combination. Earlier implementation
// flagged any project func whose package name happened to match built-in
// package names (auth/session/cache/file/…), but that is the **expected**
// fallback-chain pattern per manual-for-ai.md: project `func/<pkg>/` takes
// precedence over `pkg/<pkg>/` for custom implementations. Only actual
// function-name collisions produce import ambiguity in codegen (BUG008).
func f01BuiltinOverride(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	builtinFuncs := make(map[string]bool, len(fs.YongolPkgSpecs))
	for _, sp := range fs.YongolPkgSpecs {
		builtinFuncs[sp.Package+"."+sp.Name] = true
	}
	var diags []diagnostic.Diagnostic
	for _, sp := range fs.ProjectFuncSpecs {
		key := sp.Package + "." + sp.Name
		if !builtinFuncs[key] {
			continue
		}
		diags = append(diags, diagnostic.Diagnostic{
			File:    sp.Package + "/" + sp.Name + ".go",
			Line:    sp.Line,
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelWarning,
			Message: "[F-1] project func " + key + " overrides built-in pkg/" + key,
			Advice:  "의도된 override 가 아니면 프로젝트 func 의 이름을 변경하세요 (의도된 override 라면 WARNING 은 무시해도 됩니다)",
		})
	}
	return diags
}
