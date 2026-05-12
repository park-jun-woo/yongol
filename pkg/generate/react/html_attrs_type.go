//ff:func feature=gen-react type=util control=selection
//ff:what HTML 태그에 대응하는 React HTML attributes 타입을 반환한다

package react

// htmlAttrsType returns the React HTML attributes type for a given tag.
func htmlAttrsType(tag string) string {
	switch tag {
	case "button":
		return "ButtonHTMLAttributes<HTMLButtonElement>"
	case "input":
		return "InputHTMLAttributes<HTMLInputElement>"
	case "select":
		return "SelectHTMLAttributes<HTMLSelectElement>"
	case "textarea":
		return "TextareaHTMLAttributes<HTMLTextAreaElement>"
	case "form":
		return "FormHTMLAttributes<HTMLFormElement>"
	case "table":
		return "TableHTMLAttributes<HTMLTableElement>"
	case "label":
		return "LabelHTMLAttributes<HTMLLabelElement>"
	case "a":
		return "AnchorHTMLAttributes<HTMLAnchorElement>"
	default:
		return "HTMLAttributes<HTMLDivElement>"
	}
}
