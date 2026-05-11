//ff:func feature=stml-gen type=generator control=sequence
//ff:what ComponentRef의 JSX를 생성한다
package stml

import (
	"fmt"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// renderComponentJSX generates JSX for a ComponentRef.
func renderComponentJSX(c stmlparser.ComponentRef, dataVar string, indent int) string {
	ind := indentStr(indent)
	if c.Bind != "" {
		return fmt.Sprintf("%s<%s data={%s.%s} />", ind, c.Name, dataVar, c.Bind)
	}
	return fmt.Sprintf("%s<%s />", ind, c.Name)
}
