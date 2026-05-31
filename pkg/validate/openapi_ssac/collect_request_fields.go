//ff:func feature=validate type=util control=iteration dimension=1 topic=ssac-openapi
//ff:what collectRequestFields — ServiceFunc의 시퀀스에서 request.X 참조 필드 수집

package openapi_ssac

import (
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// collectRequestFields collects every `request.<Field>` reference used by fn.
func collectRequestFields(fn ssac.ServiceFunc) map[string]bool {
	fields := make(map[string]bool)
	for _, seq := range fn.Sequences {
		collectFromArgs(fields, seq.Args)
		collectFromValueMap(fields, seq.Inputs)
		collectFromValueMap(fields, seq.Fields)
	}
	return fields
}
