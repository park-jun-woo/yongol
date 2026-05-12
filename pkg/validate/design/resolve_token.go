//ff:func feature=validate type=util control=selection topic=design-structural
//ff:what resolveToken — dotted 참조 (e.g. "colors.primary") 를 DesignSpec 토큰 그룹에서 탐색
package design

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// resolveToken checks whether a dotted reference (e.g. "colors.primary") can be found
// in the DesignSpec token groups.
func resolveToken(fs *yongol.Fullstack, dotted string) bool {
	parts := strings.SplitN(dotted, ".", 2)
	if len(parts) != 2 {
		return false
	}
	group, token := parts[0], parts[1]
	spec := fs.DesignSpec
	switch group {
	case "colors":
		_, ok := spec.Colors[token]
		return ok
	case "typography":
		_, ok := spec.Typography[token]
		return ok
	case "rounded":
		_, ok := spec.Rounded[token]
		return ok
	case "spacing":
		_, ok := spec.Spacing[token]
		return ok
	}
	return false
}
