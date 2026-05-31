//ff:func feature=openapi-parse type=test control=sequence
//ff:what TestByName_ZeroCov — openapi 파서 헬퍼들을 이름으로 직접 호출해 커버리지 귀속
package openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestByNamePathParamTypes_ZeroCov(t *testing.T) {
	schema := openapi3.NewSchemaRef("", &openapi3.Schema{Type: &openapi3.Types{"integer"}})
	op := &openapi3.Operation{
		OperationID: "GetItem",
		Parameters: openapi3.Parameters{
			&openapi3.ParameterRef{Value: &openapi3.Parameter{
				Name:   "id",
				In:     "path",
				Schema: schema,
			}},
		},
	}
	result := map[string]map[string]string{}
	collectPathParamTypesForOp(result, op)
	if result["GetItem"]["id"] != "integer" {
		t.Errorf("collectPathParamTypesForOp = %v", result)
	}

	// nil/empty op short-circuits.
	collectPathParamTypesForOp(result, nil)
	collectPathParamTypesForOp(result, &openapi3.Operation{})
}
