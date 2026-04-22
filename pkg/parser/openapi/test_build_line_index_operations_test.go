//ff:func feature=manifest type=parser control=sequence
//ff:what BuildLineIndex 가 operationId 줄 번호를 올바르게 색인하는지 검증

package openapi

import "testing"

func TestBuildLineIndex_Operations(t *testing.T) {
	path := writeFixture(t)
	idx, err := BuildLineIndex(path)
	if err != nil {
		t.Fatalf("BuildLineIndex: %v", err)
	}

	// "operationId: Login" 줄을 찾는다 — fixture 에서 16번 줄.
	if got, want := idx.OperationLine("Login"), 16; got != want {
		t.Errorf("OperationLine(Login) = %d, want %d", got, want)
	}
	// "operationId: GetCurrentUser" 줄
	if got, want := idx.OperationLine("GetCurrentUser"), 38; got != want {
		t.Errorf("OperationLine(GetCurrentUser) = %d, want %d", got, want)
	}
}
