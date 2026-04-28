//ff:func feature=validate type=test control=iteration dimension=1 topic=ssac-structural
//ff:what S-25 정합성 — parser ValidSequenceTypes ⊆ knownSeqTypes (BUG-001 회귀 방지)

package ssac

import (
	"testing"

	parsessac "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// TestS25KnownSeqTypesCoversParserSet asserts that every sequence directive
// declared by the parser (parsessac.ValidSequenceTypes) is also present in
// validate's knownSeqTypes whitelist. BUG-001 was a violation of this
// invariant: parser.SeqEval was added but knownSeqTypes was not updated, so
// every @eval directive raised a spurious S-25 error.
func TestS25KnownSeqTypesCoversParserSet(t *testing.T) {
	for seqType := range parsessac.ValidSequenceTypes {
		if !knownSeqTypes[seqType] {
			t.Errorf("knownSeqTypes is missing parser-declared sequence type %q "+
				"(parser exports it via ValidSequenceTypes; S-25 would falsely "+
				"reject every use of @%s)", seqType, seqType)
		}
	}
}
