//ff:func feature=gen-gogin type=test control=iteration topic=sqlc-post
//ff:what TestSortedColumnKeys — 컬럼 맵 키를 오름차순 정렬 반환 + 빈 맵 처리 검증

package sqlcpost

import (
	"reflect"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

func TestSortedColumnKeys(t *testing.T) {
	m := map[string]ddl.Column{
		"name":  {Name: "name"},
		"id":    {Name: "id"},
		"email": {Name: "email"},
	}
	got := sortedColumnKeys(m)
	want := []string{"email", "id", "name"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("sortedColumnKeys = %v, want %v", got, want)
	}
}

func TestSortedColumnKeys_Empty(t *testing.T) {
	if got := sortedColumnKeys(map[string]ddl.Column{}); len(got) != 0 {
		t.Errorf("expected empty result, got %v", got)
	}
}
