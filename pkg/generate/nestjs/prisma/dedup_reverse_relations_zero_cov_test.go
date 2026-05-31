//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestIsPrimaryKey_ZeroCov — isPrimaryKey 포함/미포함 분기 검증
package prisma

import (
	"testing"
)

func TestDedupReverseRelations_ZeroCov(t *testing.T) {
	rels := []reverseRelation{
		{FieldName: "posts", ModelName: "Post"},
		{FieldName: "posts", ModelName: "Post"},
		{FieldName: "comments", ModelName: "Comment"},
	}
	got := dedupReverseRelations(rels)
	if len(got) != 2 {
		t.Errorf("expected 2 deduped, got %d", len(got))
	}
}
