//ff:func feature=stml-gen type=generator control=sequence
//ff:what 폼 필드(input/component)의 JSX를 생성한다
package stml

import (
	"fmt"
	"strings"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// renderFieldJSX generates JSX for a form field.
func renderFieldJSX(f stmlparser.FieldBind, formName string, indent int) string {
	ind := indentStr(indent)

	// data-component field
	if strings.HasPrefix(f.Tag, "data-component:") {
		comp := strings.TrimPrefix(f.Tag, "data-component:")
		return fmt.Sprintf("%s<%s {...%s.register('%s')} />", ind, comp, formName, f.Name)
	}

	var attrs []string
	if f.Type != "" {
		attrs = append(attrs, fmt.Sprintf(`type="%s"`, f.Type))
	}
	if f.Placeholder != "" {
		attrs = append(attrs, fmt.Sprintf(`placeholder="%s"`, f.Placeholder))
	}
	if f.ClassName != "" {
		attrs = append(attrs, fmt.Sprintf(`className="%s"`, f.ClassName))
	}

	reg := fmt.Sprintf("{...%s.register('%s'", formName, f.Name)
	if f.Type == "number" {
		reg += ", { valueAsNumber: true }"
	}
	reg += ")}"

	attrStr := ""
	if len(attrs) > 0 {
		attrStr = " " + strings.Join(attrs, " ")
	}

	return fmt.Sprintf("%s<input%s %s />", ind, attrStr, reg)
}
