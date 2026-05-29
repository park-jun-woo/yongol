//ff:func feature=agent type=test control=iteration dimension=2
//ff:what TestMatchBySchema — $ref 스키마 이름을 feature op 와 매핑 검증

package agent

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func TestMatchBySchema(t *testing.T) {
	feats := []features.Feature{
		{Op: "CreateUser"},
		{Op: "ListOrders"},
	}
	msg := "schema $ref #/components/schemas/createuser invalid"
	got := matchBySchema(msg, nil, feats)
	if len(got) != 1 || got[0] != "CreateUser" {
		t.Errorf("matchBySchema = %v, want [CreateUser]", got)
	}
	// No schema ref -> nil.
	if got := matchBySchema("plain error", nil, feats); got != nil {
		t.Errorf("expected nil, got %v", got)
	}
}
