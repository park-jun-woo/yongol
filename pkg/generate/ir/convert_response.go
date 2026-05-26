//ff:func feature=gen-ir type=util control=sequence
//ff:what convertResponse -- @response 시퀀스 → ResponseOp IR 변환

package ir

import (
	"sort"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// convertResponse converts a @response sequence to an IR Op.
func convertResponse(seq ssac.Sequence) Op {
	op := ResponseOp{}

	if len(seq.Fields) > 0 {
		fields := make([]ResponseField, 0, len(seq.Fields))
		keys := make([]string, 0, len(seq.Fields))
		for k := range seq.Fields {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for _, k := range keys {
			fields = append(fields, ResponseField{
				Name:   k,
				Source: seq.Fields[k],
			})
		}
		op.Fields = fields
	} else if seq.Target != "" {
		op.SingleVar = seq.Target
	}

	return Op{Kind: OpResponse, Response: &op}
}
