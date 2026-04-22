//ff:func feature=gen-gogin type=util control=iteration dimension=1
//ff:what parseClaims — manifest claims → sorted []ClaimField

package auth

import (
	"sort"

	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

// parseClaims converts manifest ClaimDef map to a sorted slice of ClaimField.
// Sort by field name for deterministic output across runs.
func parseClaims(claims map[string]manifest.ClaimDef) []ClaimField {
	var fields []ClaimField
	for name, def := range claims {
		goType := def.GoType
		if goType == "" {
			goType = "string"
		}
		fields = append(fields, ClaimField{Name: name, Key: def.Key, GoType: goType})
	}
	sort.Slice(fields, func(i, j int) bool { return fields[i].Name < fields[j].Name })
	return fields
}
