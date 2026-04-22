//ff:func feature=external type=util control=iteration dimension=1
//ff:what OpenAPI 제목 또는 파일명에서 서비스 이름을 추론한다
package external

import (
	"path/filepath"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

func inferServiceName(source string, doc *openapi3.T) string {
	if doc.Info != nil && doc.Info.Title != "" {
		return toPascalCase(doc.Info.Title)
	}
	base := filepath.Base(source)
	// Strip known suffixes: .openapi.yaml, .openapi.json, .yaml, .json
	for _, suffix := range []string{".openapi.yaml", ".openapi.json", ".openapi.yml", ".yaml", ".yml", ".json"} {
		if strings.HasSuffix(base, suffix) {
			base = base[:len(base)-len(suffix)]
			break
		}
	}
	return toPascalCase(base)
}
