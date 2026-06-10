//ff:func feature=stml-gen type=util control=sequence
//ff:what data-link-params 소스를 JSX 표현식으로 변환한다 (route.X → useParams 변수, item.* 그대로)
package stml

import "strings"

// linkParamExpr converts a data-link-params source into the JSX expression
// interpolated into the Link `to` template: route.<Name> becomes the
// useParams() destructured variable, item.<Field> is valid as-is inside
// the data-each map callback scope (paramSourceExpr's contract).
func linkParamExpr(source string) string {
	if strings.HasPrefix(source, "route.") {
		return strings.TrimPrefix(source, "route.")
	}
	return source
}
