//ff:func feature=validate type=test-helper control=sequence
//ff:what subtestTestXos67ResponseFieldTypeMatchUuidFieldBoundToStringLiteralErrorsFalseNegativeClosed — uuid field bound to string literal errors (false negative closed) 서브테스트
package openapi_ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/rule"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func subtestTestXos67ResponseFieldTypeMatchUuidFieldBoundToStringLiteralErrorsFalseNegativeClosed(t *testing.T) {

	fs := &yongol.Fullstack{
		ServiceFuncs: []ssac.ServiceFunc{
			{
				Name:     "cancelMatch",
				FileName: "cancel_match.ssac",
				Sequences: []ssac.Sequence{
					{Type: "response", Fields: map[string]string{"match_id": `"abc"`}},
				},
			},
		},
	}
	g := &rule.Ground{
		Types: map[string]string{
			"OpenAPI.response.cancelMatch.match_id": "openapi_types.UUID",
		},
	}
	fs.SetGround(g)
	diags := xos67ResponseFieldType(fs)
	if len(diags) != 1 {
		t.Fatalf("expected 1, got %d: %+v", len(diags), diags)
	}
	if !strings.Contains(diags[0].Message, "XOS-67") {
		t.Errorf("Message missing XOS-67: %s", diags[0].Message)
	}

}
