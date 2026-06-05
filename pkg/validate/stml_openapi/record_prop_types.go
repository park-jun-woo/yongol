//ff:func feature=validate type=util control=iteration dimension=1 topic=stml-openapi
//ff:what recordPropTypes — property 맵의 각 항목 이름 -> 첫 번째 선언 type 을 out 에 복사

package stml_openapi

import "github.com/getkin/kin-openapi/openapi3"

// recordPropTypes copies each property's first declared type into out.
// Properties without a declared type (incl. nil refs) map to "".
func recordPropTypes(out map[string]string, props openapi3.Schemas) {
	for name, ref := range props {
		if ref == nil || ref.Value == nil {
			out[name] = ""
			continue
		}
		if ts := ref.Value.Type.Slice(); len(ts) > 0 {
			out[name] = ts[0]
		} else {
			out[name] = ""
		}
	}
}
