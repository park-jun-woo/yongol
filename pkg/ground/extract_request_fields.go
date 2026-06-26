//ff:func feature=rule type=util control=iteration dimension=1
//ff:what extractRequestFields — requestBody에서 필드명 집합 추출
package ground

import (
	"github.com/getkin/kin-openapi/openapi3"

	"github.com/park-jun-woo/yongol/pkg/rule"
)

func extractRequestFields(body *openapi3.RequestBody) rule.StringSet {
	if body.Content == nil {
		return nil
	}
	ct := body.Content.Get("application/json")
	if ct == nil {
		ct = body.Content.Get("multipart/form-data")
	}
	if ct == nil || ct.Schema == nil || ct.Schema.Value == nil {
		return nil
	}
	fields := make(rule.StringSet, len(ct.Schema.Value.Properties))
	for name := range ct.Schema.Value.Properties {
		fields[name] = true
	}
	return fields
}
