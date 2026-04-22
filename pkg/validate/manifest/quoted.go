//ff:func feature=validate type=util control=sequence topic=manifest-structural
//ff:what quoted — 문자열을 쌍따옴표로 감싸 메시지 리터럴로 사용

package manifest

// quoted wraps s with surrounding double quotes for use in diagnostic messages.
func quoted(s string) string {
	return "\"" + s + "\""
}
