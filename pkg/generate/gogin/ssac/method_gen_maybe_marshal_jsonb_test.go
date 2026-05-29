//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what methodGen.maybeMarshalJSONB 단위 테스트 (body JSONB 필드일 때만 json.Marshal 프리앰블)

package ssac

import "testing"

func TestMethodGenMaybeMarshalJSONB(t *testing.T) {
	t.Run("no jsonb fields → not applicable", func(t *testing.T) {
		g := &methodGen{}
		_, _, ok := g.maybeMarshalJSONB("Payload", "request.payload")
		if ok {
			t.Errorf("expected ok=false when no BodyJSONBFields")
		}
	})
	t.Run("non-jsonb source → not applicable", func(t *testing.T) {
		g := &methodGen{BodyJSONBFields: map[string]bool{"payload": true}}
		_, _, ok := g.maybeMarshalJSONB("Title", "request.title")
		if ok {
			t.Errorf("expected ok=false for non-jsonb source")
		}
	})
	t.Run("jsonb body field → marshal preamble", func(t *testing.T) {
		g := &methodGen{
			BodyJSONBFields: map[string]bool{"payload": true},
		}
		rawVar, pre, ok := g.maybeMarshalJSONB("payload", "request.payload")
		if !ok {
			t.Fatal("expected ok=true")
		}
		if rawVar != "payloadRaw" {
			t.Errorf("rawVar = %q, want payloadRaw", rawVar)
		}
		if len(pre) != 2 {
			t.Fatalf("expected 2 preamble lines, got %d: %v", len(pre), pre)
		}
		if pre[0] != "payloadRaw, err := json.Marshal(request.Body.Payload)" {
			t.Errorf("preamble[0] = %q", pre[0])
		}
		if pre[1] != "if err != nil { return nil, err }" {
			t.Errorf("preamble[1] = %q", pre[1])
		}
	})
}
