//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestCollectFuncInnerTypeNames — @call 응답 내부 복합 타입 수집 검증

package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/funcspec"
)

func TestCollectFuncInnerTypeNames(t *testing.T) {
	baseSpecs := []funcspec.FuncSpec{{
		Name: "LoadMessages", Package: "chat",
		ResponseFields: []funcspec.Field{
			{Name: "Messages", Type: "[]ChatMessage", JSONName: "messages"},
			{Name: "Pagination", Type: "ChatPagination", JSONName: "pagination"},
			{Name: "Total", Type: "int", JSONName: "total"},
		},
	}}
	chatInfo := funcRespInfo{PkgAlias: "chat", ImportPath: "example.com/app/internal/chat"}
	baseResp := map[string]funcRespInfo{"LoadMessagesResponse": chatInfo}

	t.Run("Basic", func(t *testing.T) {
		needed := map[string]bool{"LoadMessagesResponse": true, "ChatMessage": true, "ChatPagination": true}
		result := collectFuncInnerTypeNames(baseResp, baseSpecs, needed, map[string]bool{})
		if len(result) != 2 {
			t.Fatalf("expected 2 inner types, got %d: %v", len(result), result)
		}
		if result["ChatMessage"].PkgAlias != "chat" {
			t.Errorf("ChatMessage PkgAlias = %q", result["ChatMessage"].PkgAlias)
		}
	})

	t.Run("ExcludesSqlcModels", func(t *testing.T) {
		needed := map[string]bool{"ChatMessage": true}
		sqlc := map[string]bool{"ChatMessage": true}
		result := collectFuncInnerTypeNames(baseResp, baseSpecs, needed, sqlc)
		if len(result) != 0 {
			t.Fatalf("expected 0 (sqlc excluded), got %d", len(result))
		}
	})

	t.Run("ExcludesDirectRespTypes", func(t *testing.T) {
		resp := map[string]funcRespInfo{"LoadMessagesResponse": chatInfo, "ChatMessage": chatInfo}
		needed := map[string]bool{"ChatMessage": true}
		result := collectFuncInnerTypeNames(resp, baseSpecs, needed, map[string]bool{})
		if len(result) != 0 {
			t.Fatalf("expected 0 (already in funcRespNames), got %d", len(result))
		}
	})

	t.Run("NotInNeeded", func(t *testing.T) {
		result := collectFuncInnerTypeNames(baseResp, baseSpecs, map[string]bool{}, map[string]bool{})
		if len(result) != 0 {
			t.Fatalf("expected 0 (not in needed), got %d", len(result))
		}
	})

	t.Run("Empty", func(t *testing.T) {
		result := collectFuncInnerTypeNames(nil, nil, nil, nil)
		if len(result) != 0 {
			t.Fatalf("expected 0, got %d", len(result))
		}
	})
}
