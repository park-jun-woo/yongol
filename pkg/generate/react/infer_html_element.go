//ff:func feature=gen-react type=util control=selection
//ff:what HTML 태그에서 TypeScript element 타입을 추론한다

package react

// inferHTMLElement maps HTML tags to their TypeScript element types.
func inferHTMLElement(tag string) string {
	switch tag {
	case "button":
		return "HTMLButtonElement"
	case "input":
		return "HTMLInputElement"
	case "select":
		return "HTMLSelectElement"
	case "textarea":
		return "HTMLTextAreaElement"
	case "form":
		return "HTMLFormElement"
	case "table":
		return "HTMLTableElement"
	case "label":
		return "HTMLLabelElement"
	case "a":
		return "HTMLAnchorElement"
	default:
		return "HTMLDivElement"
	}
}
