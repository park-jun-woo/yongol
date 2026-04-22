//ff:func feature=manifest type=parser control=sequence
//ff:what BuildLineIndex 가 components.schemas 의 스키마 및 property 줄 번호를 올바르게 색인하는지 검증

package openapi

import "testing"

func TestBuildLineIndex_Schemas(t *testing.T) {
	path := writeFixture(t)
	idx, _ := BuildLineIndex(path)
	// "User:" 줄
	if got, want := idx.SchemaLine("User"), 7; got != want {
		t.Errorf("SchemaLine(User) = %d, want %d", got, want)
	}
	if got, want := idx.SchemaPropertyLine("User", "email"), 11; got != want {
		t.Errorf("SchemaPropertyLine(User,email) = %d, want %d", got, want)
	}
	if got, want := idx.SchemaPropertyLine("User", "name"), 12; got != want {
		t.Errorf("SchemaPropertyLine(User,name) = %d, want %d", got, want)
	}
}
