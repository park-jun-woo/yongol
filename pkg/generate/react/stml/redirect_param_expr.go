//ff:func feature=stml-gen type=util control=sequence
//ff:what data-redirect-params 소스를 JSX 표현식으로 변환한다 (route.X → useParams 변수, respField → data.<field>)
package stml

import "strings"

// redirectParamExpr converts a data-redirect-params source into the JSX
// expression interpolated into the navigate() template: route.<Name>
// becomes the useParams() destructured variable, an unprefixed respField
// reads from the mutation's 2xx response object (`data`, the onSuccess
// parameter — page-flow Phase008).
func redirectParamExpr(source string) string {
	if strings.HasPrefix(source, "route.") {
		return strings.TrimPrefix(source, "route.")
	}
	return "data." + source
}
