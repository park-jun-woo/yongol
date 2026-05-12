//ff:func feature=gen-react type=util control=sequence
//ff:what 레이아웃 이름을 PascalCase + "Layout" 접미사로 변환한다

package react

// layoutComponentName converts a layout name to PascalCase + "Layout" suffix.
// e.g. "app" → "AppLayout", "auth" → "AuthLayout", "main-nav" → "MainNavLayout"
func layoutComponentName(name string) string {
	return kebabToPascal(name) + "Layout"
}
