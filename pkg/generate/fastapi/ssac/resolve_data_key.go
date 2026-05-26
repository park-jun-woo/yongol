//ff:func feature=gen-fastapi type=util control=sequence
//ff:what resolveDataKey — FieldArg → SQLAlchemy data 절 키 추출

package ssac

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// resolveDataKey extracts a key for the data clause from a FieldArg.
func resolveDataKey(a ir.FieldArg) string {
	if a.Key != "" {
		return a.Key
	}
	return strings.TrimPrefix(a.Field, ".")
}
