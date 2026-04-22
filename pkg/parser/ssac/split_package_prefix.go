//ff:func feature=ssac-parse type=util control=sequence
//ff:what splitPackagePrefix — splits "session.Session.Get" into ("session", "Session.Get") package prefix
package ssac

import "strings"

// splitPackagePrefix splits "session.Session.Get" into ("session", "Session.Get").
// "Course.FindByID" → ("", "Course.FindByID") — 2-part form has no package prefix.
// @call already uses the pkg.Func form so this function is not used for it.
func splitPackagePrefix(model string) (pkg, rest string) {
	// dot count: 1 → legacy Model.Method, 2+ → pkg.Model.Method
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
