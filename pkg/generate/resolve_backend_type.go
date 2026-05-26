//ff:func feature=generate type=util control=selection
//ff:what ResolveBackendType — lang+framework 조합을 BackendType 으로 변환
package generate

import "fmt"

// ResolveBackendType maps a (lang, framework) pair from manifest.yaml to the
// corresponding BackendType. Unknown combinations return an error.
func ResolveBackendType(lang, framework string) (BackendType, error) {
	switch lang + "+" + framework {
	case "go+gin":
		return GoGin, nil
	case "typescript+nestjs":
		return NestJS, nil
	case "python+fastapi":
		return FastAPI, nil
	default:
		return "", fmt.Errorf("unsupported backend: lang=%q framework=%q", lang, framework)
	}
}
