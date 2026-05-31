//ff:func feature=validate type=test control=sequence topic=ssac-sqlc
//ff:what zz_zerocov_test — ssac_sqlc 0% 헬퍼 (Run / collectInputKeys / buildQueryParamMap / checkSingleInputKeyCase / checkSeqInputKeyCase) 단위 테스트
package ssac_sqlc

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestBuildXqs18OAPIParamTypeMap_ZeroCov(t *testing.T) {
	intType := &openapi3.Types{"integer"}
	op := &openapi3.Operation{
		Parameters: openapi3.Parameters{
			&openapi3.ParameterRef{Value: &openapi3.Parameter{
				Name: "page",
				Schema: &openapi3.SchemaRef{Value: &openapi3.Schema{
					Type:   intType,
					Format: "int64",
				}},
			}},
			nil, // skipped: nil ref
			&openapi3.ParameterRef{Value: &openapi3.Parameter{
				Name:   "noschema",
				Schema: nil, // skipped: no schema
			}},
		},
	}
	m := buildXqs18OAPIParamTypeMap(op)
	if m["page"] == "" {
		t.Errorf("expected page type, got %v", m)
	}
	if _, ok := m["noschema"]; ok {
		t.Error("param without schema should be skipped")
	}
}
