//ff:func feature=gen-ir type=util control=sequence
//ff:what convertAuth -- @auth 시퀀스 → AuthOp IR 변환 (Ownership 정보 이식)

package ir

import (
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// convertAuth converts a @auth sequence to an IR Op. When the Fullstack
// context contains ParsedPolicies with @ownership annotations matching the
// sequence's Resource, AuthOp.Ownership is populated with the lookup metadata.
func convertAuth(seq ssac.Sequence, fs *yongol.Fullstack) Op {
	statusCode := seq.ErrStatus
	if statusCode == 0 {
		statusCode = 403
	}
	op := AuthOp{
		Action:     seq.Action,
		Resource:   seq.Resource,
		Inputs:     convertInputsToFieldArgs(seq.Inputs),
		Message:    seq.Message,
		StatusCode: statusCode,
	}

	// Enrich with ownership info from Rego @ownership annotations.
	// Only populate Ownership when ResourceID is present and non-zero in
	// the sequence Inputs, mirroring gogin's ownership_lookup.go:23-25.
	rawRID, hasRID := seq.Inputs["ResourceID"]
	if fs != nil && hasRID && !isResourceIDZeroIR(rawRID) {
		op.Ownership = lookupOwnershipIR(fs, seq.Resource)
	}

	return Op{Kind: OpAuth, Auth: &op}
}
