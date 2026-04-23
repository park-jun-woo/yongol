//ff:func feature=gen-react type=util control=sequence
//ff:what findOpenAPISpec — openapi-typescript 용 OpenAPI yaml 경로 반환 (<specsDir>/api/openapi.yaml)

package react

import "github.com/park-jun-woo/yongol/pkg/yongol"

// findOpenAPISpec returns the path to the OpenAPI yaml used by openapi-typescript.
// Convention: <specsDir>/api/openapi.yaml.
func findOpenAPISpec(fs *yongol.Fullstack) string {
	if fs == nil || fs.SpecsDir == "" {
		return ""
	}
	return fs.SpecsDir + "/api/openapi.yaml"
}
