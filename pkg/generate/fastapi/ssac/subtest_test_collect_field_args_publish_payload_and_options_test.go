//ff:func feature=gen-fastapi type=test-helper control=sequence
//ff:what subtestTestCollectFieldArgsPublishPayloadAndOptions — PublishPayloadAndOptions 서브테스트
package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func subtestTestCollectFieldArgsPublishPayloadAndOptions(t *testing.T) {

	op := ir.Op{Kind: ir.OpPublish, Publish: &ir.PublishOp{
		Payload: []ir.FieldArg{faKey("pl")},
		Options: []ir.FieldArg{faKey("op")},
	}}
	got := collectFieldArgs(op)
	if len(got) != 2 || got[0].Key != "pl" || got[1].Key != "op" {
		t.Errorf("got %v, want [pl op]", got)
	}

}
