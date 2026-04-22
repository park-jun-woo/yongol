//ff:func feature=validate type=util control=iteration dimension=2 topic=config-check
//ff:what missingSecurityNames — operation security 참조 중 middleware set에 없는 이름 반환

package openapi_manifest

import "github.com/getkin/kin-openapi/openapi3"

func missingSecurityNames(op *openapi3.Operation, mwSet map[string]bool) []string {
	if op == nil || op.Security == nil {
		return nil
	}
	var missing []string
	for _, req := range *op.Security {
		for name := range req {
			if !mwSet[name] {
				missing = append(missing, name)
			}
		}
	}
	return missing
}
