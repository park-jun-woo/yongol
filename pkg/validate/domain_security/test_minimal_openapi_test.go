//ff:func feature=validate type=test-helper control=iteration dimension=2 topic=domain-security
//ff:what minimalOpenAPI — 테스트용 최소 OpenAPI YAML 생성
package domain_security

// minimalOpenAPI returns minimal valid OpenAPI YAML with given operations.
// ops is a map of path -> method -> operationId.
func minimalOpenAPI(ops map[string]map[string]opDef) string {
	yaml := "openapi: '3.0.0'\ninfo:\n  title: test\n  version: '1.0'\npaths:\n"
	for path, methods := range ops {
		yaml += "  " + path + ":\n"
		for method, def := range methods {
			yaml += "    " + method + ":\n"
			yaml += "      operationId: " + def.ID + "\n"
			yaml += "      responses:\n        '200':\n          description: ok\n"
			if def.Security != "" {
				yaml += "      security: " + def.Security + "\n"
			}
		}
	}
	return yaml
}
