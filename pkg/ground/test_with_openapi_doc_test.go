//ff:func feature=rule type=test-helper control=sequence
//ff:what withOpenAPIDoc — prebuilt *openapi3.T 를 Fullstack.OpenAPIDoc 에 부착하는 option

package ground

import (
	"github.com/getkin/kin-openapi/openapi3"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// withOpenAPIDoc attaches a prebuilt *openapi3.T.
func withOpenAPIDoc(doc *openapi3.T) func(*yongol.Fullstack) {
	return func(fs *yongol.Fullstack) { fs.OpenAPIDoc = doc }
}
