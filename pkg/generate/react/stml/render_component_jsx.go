//ff:func feature=stml-gen type=generator control=sequence
//ff:what ComponentRef의 JSX를 생성한다 (className/variant/size prop passthrough 포함)
package stml

import (
	"fmt"
	"strings"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// renderComponentJSX generates JSX for a ComponentRef. Non-empty ClassName,
// Variant, and Size fields are emitted as className, variant, and size props
// respectively. When data-bind is present, data={dataVar.bind} is emitted.
func renderComponentJSX(c stmlparser.ComponentRef, dataVar string, indent int) string {
	ind := indentStr(indent)

	var props []string
	if c.Bind != "" {
		props = append(props, fmt.Sprintf("data={%s.%s}", dataVar, optionalChainPath(c.Bind)))
	}
	if c.ClassName != "" {
		props = append(props, fmt.Sprintf("className=%q", c.ClassName))
	}
	if c.Variant != "" {
		props = append(props, fmt.Sprintf("variant=%q", c.Variant))
	}
	if c.Size != "" {
		props = append(props, fmt.Sprintf("size=%q", c.Size))
	}

	if len(props) == 0 {
		return fmt.Sprintf("%s<%s />", ind, c.Name)
	}
	return fmt.Sprintf("%s<%s %s />", ind, c.Name, strings.Join(props, " "))
}
