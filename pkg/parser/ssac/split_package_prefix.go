//ff:func feature=ssac-parse type=util control=sequence
//ff:what "session.Session.Get" → ("session", "Session.Get") 패키지 접두사 분리
package ssac

import "strings"

// splitPackagePrefix는 "session.Session.Get" → ("session", "Session.Get")로 분리한다.
// "Course.FindByID" → ("", "Course.FindByID") — 2-part는 패키지 없음.
// @call은 이미 pkg.Func 형식이므로 이 함수를 사용하지 않는다.
func splitPackagePrefix(model string) (pkg, rest string) {
	// dot 개수: 1개 → 기존 Model.Method, 2개 이상 → pkg.Model.Method
	firstDot := strings.IndexByte(model, '.')
	if firstDot < 0 {
		return "", model
	}
	secondDot := strings.IndexByte(model[firstDot+1:], '.')
	if secondDot < 0 {
		// "Model.Method" — no package prefix
		return "", model
	}
	// "pkg.Model.Method" — first part is package (lowercase check)
	pkg = model[:firstDot]
	if len(pkg) > 0 && pkg[0] >= 'a' && pkg[0] <= 'z' {
		return pkg, model[firstDot+1:]
	}
	// If first part starts with uppercase, it's not a package prefix
	return "", model
}
