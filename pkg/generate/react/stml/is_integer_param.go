//ff:func feature=stml-gen type=util control=sequence
//ff:what operationId의 path 파라미터가 integer 타입인지 확인한다
package stml

// isIntegerParam returns true when the named path parameter is declared
// as "integer" in the OpenAPI spec for the given operation.
func isIntegerParam(opID, paramName string, pathParamTypes map[string]map[string]string) bool {
	if pathParamTypes == nil {
		return false
	}
	params, ok := pathParamTypes[opID]
	if !ok {
		return false
	}
	return params[paramName] == "integer"
}
