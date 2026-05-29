//ff:func feature=agent type=test control=iteration dimension=2
//ff:what TestAssemblePathBlock — security/params/requestBody/error-response 분기별 path 블록 조립 검증

package agent

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func TestAssemblePathBlock(t *testing.T) {
	schema200 := map[string]any{"type": "object"}

	t.Run("CreatePublicWithParamsAndBody", func(t *testing.T) {
		feat := features.Feature{Op: "CreateWorkflow", Desc: "make one", Public: true}
		params := []any{map[string]any{"name": "q", "in": "query"}}
		reqBody := map[string]any{"required": true}
		errResps := map[string]any{"404": map[string]any{"description": "nope"}}

		block := assemblePathBlock(feat, params, reqBody, schema200, errResps)

		op, ok := block["post"].(map[string]any)
		if !ok {
			t.Fatalf("expected 'post' method block, got: %v", block)
		}
		if op["operationId"] != "CreateWorkflow" {
			t.Errorf("operationId: got %v", op["operationId"])
		}
		if op["summary"] != "make one" {
			t.Errorf("summary: got %v", op["summary"])
		}
		// Public → no security.
		if _, has := op["security"]; has {
			t.Error("public op must not carry security")
		}
		// params present.
		if _, has := op["parameters"]; !has {
			t.Error("expected parameters")
		}
		// post needs request body.
		if _, has := op["requestBody"]; !has {
			t.Error("expected requestBody for post")
		}
		responses, _ := op["responses"].(map[string]any)
		if _, has := responses["200"]; !has {
			t.Error("expected 200 response")
		}
		if _, has := responses["404"]; !has {
			t.Error("expected merged 404 error response")
		}
	})

	t.Run("GetPrivateNoParamsNoBody", func(t *testing.T) {
		feat := features.Feature{Op: "GetWorkflow", Desc: "fetch", Public: false}

		block := assemblePathBlock(feat, nil, map[string]any{"x": 1}, schema200, nil)

		op, ok := block["get"].(map[string]any)
		if !ok {
			t.Fatalf("expected 'get' method block, got: %v", block)
		}
		// Private → security present.
		if _, has := op["security"]; !has {
			t.Error("private op must carry security")
		}
		// No params provided → omitted.
		if _, has := op["parameters"]; has {
			t.Error("did not expect parameters")
		}
		// get does not need a request body even though reqBody is non-nil.
		if _, has := op["requestBody"]; has {
			t.Error("get must not carry requestBody")
		}
		responses, _ := op["responses"].(map[string]any)
		if len(responses) != 1 {
			t.Errorf("expected only 200 response, got %d", len(responses))
		}
	})
}
