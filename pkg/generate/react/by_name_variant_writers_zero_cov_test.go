//ff:func feature=gen-react type=test control=sequence
//ff:what TestByName_ZeroCov — react 코드젠 헬퍼들을 이름으로 직접 호출해 커버리지 귀속
package react

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/design"
)

func TestByNameVariantWriters_ZeroCov(t *testing.T) {
	vk := []string{"primary", "ghost"}
	sk := []string{"sm", "lg"}
	tok := design.ComponentToken{
		Variants: map[string]string{"primary": "bg-blue", "ghost": "bg-none"},
		Sizes:    map[string]string{"sm": "text-sm", "lg": "text-lg"},
	}
	var b strings.Builder
	writeVariantTypes(&b, vk, sk)
	writeVariantProps(&b, "Button", "button", vk, sk)
	writeVariantRecords(&b, vk, sk, tok)
	if b.Len() == 0 {
		t.Errorf("variant writers produced nothing")
	}
}
