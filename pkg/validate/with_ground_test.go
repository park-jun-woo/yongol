//ff:func feature=validate type=test control=sequence
//ff:what TestWithGround — 주입된 Ground가 config에 설정되는지 검증

package validate

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/ground"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestWithGround(t *testing.T) {
	g := ground.Build(&yongol.Fullstack{})
	c := &config{}
	WithGround(g)(c)
	if c.ground != g {
		t.Fatal("expected injected ground to be set on config")
	}
}
