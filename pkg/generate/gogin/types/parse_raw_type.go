//ff:func feature=gen-gogin type=util control=sequence
//ff:what parseRawType — Column.RawType 토큰을 rawTypeInfo 로 분해 (typemap.ParseRawType 위임)

package types

import "github.com/park-jun-woo/yongol/pkg/generate/typemap"

// parseRawType splits Column.RawType into family head, parameter list,
// and array marker. Delegates to the framework-agnostic
// typemap.ParseRawType and converts the result to the package-private
// rawTypeInfo.
func parseRawType(raw string) rawTypeInfo {
	r := typemap.ParseRawType(raw)
	return rawTypeInfo{
		Head:       r.Head,
		Param:      r.Param,
		IsArray:    r.IsArray,
		MultiToken: r.MultiToken,
	}
}
