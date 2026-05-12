//ff:func feature=gen-react type=util control=selection
//ff:what 컴포넌트 이름에서 HTML 태그를 추론한다

package react

import "strings"

// inferHTMLTag maps component names to HTML tags.
func inferHTMLTag(name string) string {
	lower := strings.ToLower(name)
	switch {
	case lower == "button":
		return "button"
	case lower == "input":
		return "input"
	case lower == "select":
		return "select"
	case lower == "textarea":
		return "textarea"
	case lower == "form":
		return "form"
	case lower == "table":
		return "table"
	case lower == "label":
		return "label"
	case lower == "a" || lower == "link":
		return "a"
	default:
		return "div"
	}
}
