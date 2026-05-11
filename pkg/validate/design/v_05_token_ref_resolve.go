//ff:func feature=validate type=rule control=sequence topic=design-structural
//ff:what V-05 — components 내 토큰 참조 {group.token}이 같은 파일 내 실제 토큰으로 resolve 검증
package design

import (
	"regexp"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

var tokenRefRe = regexp.MustCompile(`\{([^}]+)\}`)

// v05TokenRefResolve validates that {group.token} references in component props
// resolve to actual tokens defined in the same DESIGN.md file.
func v05TokenRefResolve(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	spec := fs.DesignSpec
	var diags []diagnostic.Diagnostic
	for compName, comp := range spec.Components {
		for propName, propVal := range comp.Props {
			refs := tokenRefRe.FindAllStringSubmatch(propVal, -1)
			for _, ref := range refs {
				dotted := ref[1] // e.g. "colors.primary"
				if !resolveToken(fs, dotted) {
					diags = append(diags, diagnostic.Diagnostic{
						File:    spec.File,
						Phase:   diagnostic.PhaseValidate,
						Level:   diagnostic.LevelError,
						Message: "[V-05] component \"" + compName + "\" prop \"" + propName + "\" references unresolved token: {" + dotted + "}",
						Advice:  "Ensure the referenced token exists in the DESIGN.md (colors, typography, rounded, or spacing group)",
					})
				}
			}
		}
	}
	return diags
}

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
