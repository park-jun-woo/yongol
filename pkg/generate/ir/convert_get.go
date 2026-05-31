//ff:func feature=gen-ir type=util control=sequence
//ff:what convertGet -- @get 시퀀스 → GetOp IR 변환 (PaginationArgs 분리 포함)

package ir

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// paginationKeys lists the well-known pagination input keys that should be
// separated from where-clause Args into GetOp.PaginationArgs.
var paginationKeys = map[string]bool{
	"cursor":      true,
	"per_page":    true,
	"page_offset": true,
	"page":        true,
	"limit":       true,
	"offset":      true,
}

// convertGet converts a @get sequence to an IR Op. The Fullstack context
// is accepted for signature consistency but is not used directly here;
// enrichment happens in the BuildServicePlan post-processing passes.
func convertGet(seq ssac.Sequence, _ *yongol.Fullstack) Op {
	model, method := splitModelMethod(seq.Model)
	allArgs := convertInputsToFieldArgs(seq.Inputs)
	whereArgs, pagArgs := splitPaginationArgs(allArgs)

	op := GetOp{
		Model:          model,
		Method:         method,
		Args:           whereArgs,
		PaginationArgs: pagArgs,
	}
	if seq.Result != nil {
		op.VarName = seq.Result.Var
		op.VarType = seq.Result.Type
		op.IsList = seq.Result.Wrapper != "" || strings.HasPrefix(seq.Result.Type, "[]")
		op.IsCount = isCountResultType(seq.Result.Type)
	}
	return Op{Kind: OpGet, Get: &op}
}
