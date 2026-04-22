//ff:func feature=validate type=util control=sequence topic=ssac-openapi
//ff:what resolveSeqErrStatus — guard sequence 의 effective HTTP status code 산출

package openapi_ssac

import (
	"github.com/park-jun-woo/yongol/pkg/parser/funcspec"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// resolveSeqErrStatus returns the effective HTTP status code for a guard-style
// sequence and true when the sequence type is error-eligible. For @call, the
// status resolution order is: seq.ErrStatus > FuncSpec @error > errStatusTypes
// default. Non-guard sequence types return (0, false).
func resolveSeqErrStatus(seq ssac.Sequence, funcSpecs []funcspec.FuncSpec) (int, bool) {
	defaultStatus, ok := errStatusTypes[seq.Type]
	if !ok {
		return 0, false
	}
	status := defaultStatus
	if seq.ErrStatus != 0 {
		return seq.ErrStatus, true
	}
	if seq.Type == "call" {
		if fsStatus := lookupFuncSpecErrStatus(seq.Model, funcSpecs); fsStatus != 0 {
			return fsStatus, true
		}
	}
	return status, true
}
