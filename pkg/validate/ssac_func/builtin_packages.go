//ff:func feature=validate type=util control=selection topic=func-check
//ff:what builtinPackages — SSaC @call 대상 ssac 내장 패키지 이름 집합

package ssac_func

// builtinPackages is the set of ssac runtime package names that are always
// available for SSaC @call without a project-local func/ directory. Any @call
// targeting one of these packages is validated against YongolPkgSpecs.
var builtinPackages = map[string]bool{
	"auth":    true,
	"session": true,
	"cache":   true,
	"file":    true,
	"mail":    true,
	"queue":   true,
	"storage": true,
	"crypto":  true,
	"text":    true,
	"image":   true,
}
