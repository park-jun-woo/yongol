//ff:func feature=gen-ir type=util control=sequence
//ff:what convertPublish -- @publish 시퀀스 → PublishOp IR 변환

package ir

import "github.com/park-jun-woo/yongol/pkg/parser/ssac"

// convertPublish converts a @publish sequence to an IR Op.
func convertPublish(seq ssac.Sequence) Op {
	return Op{Kind: OpPublish, Publish: &PublishOp{
		Topic:   seq.Topic,
		Payload: convertInputsToFieldArgs(seq.Inputs),
		Options: convertInputsToFieldArgs(seq.Options),
	}}
}
