//ff:func feature=validate type=util control=sequence topic=config-check
//ff:what normalizeCallKey — SSaC @call model을 'pkg.camelCaseFunc' 형태로 정규화

package ssac_manifest

import (
	"strings"

	"github.com/ettle/strcase"
)

// normalizeCallKey converts an SSaC @call model string to the canonical
// "pkg.camelCaseFunc" key used by jwtBuiltinFuncs.
func normalizeCallKey(model string) string {
	parts := strings.SplitN(model, ".", 2)
	if len(parts) == 2 {
		return parts[0] + "." + strcase.ToGoCamel(parts[1])
	}
	return strcase.ToGoCamel(model)
}
