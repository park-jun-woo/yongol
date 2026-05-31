//ff:func feature=validate type=test control=sequence
//ff:what TestByName_ZeroCov — design 토큰 참조/미지 prop 검사 헬퍼 직접 호출
package design

import (
	pdesign "github.com/park-jun-woo/yongol/pkg/parser/design"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func designFS() *yongol.Fullstack {
	return &yongol.Fullstack{DesignSpec: &pdesign.DesignSpec{
		File:    "DESIGN.md",
		Colors:  map[string]string{"primary": "#fff"},
		Rounded: map[string]string{"md": "8px"},
		Spacing: map[string]string{"sm": "4px"},
	}}
}
