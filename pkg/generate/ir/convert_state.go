//ff:func feature=gen-ir type=util control=sequence
//ff:what convertState -- @state 시퀀스 → StateOp IR 변환 (AllowedFromStates 이식)

package ir

import (
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// convertState converts a @state sequence to an IR Op. When the Fullstack
// context contains StateDiagrams, AllowedFromStates is populated by
// searching for the matching diagram and collecting source states for the
// transition event.
func convertState(seq ssac.Sequence, fs *yongol.Fullstack) Op {
	statusCode := seq.ErrStatus
	if statusCode == 0 {
		statusCode = 409
	}
	op := StateOp{
		Diagram:    seq.DiagramID,
		Inputs:     convertInputsToFieldArgs(seq.Inputs),
		Transition: seq.Transition,
		Message:    seq.Message,
		StatusCode: statusCode,
	}

	// Enrich with allowed source states from Mermaid stateDiagram.
	if fs != nil {
		for _, d := range fs.StateDiagrams {
			if d.ID == seq.DiagramID || d.Symbol == seq.DiagramID {
				op.AllowedFromStates = d.ValidFromStates(seq.Transition)
				break
			}
		}
	}

	return Op{Kind: OpState, State: &op}
}
