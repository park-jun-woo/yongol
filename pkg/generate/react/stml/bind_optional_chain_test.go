//ff:func feature=stml-gen type=test control=iteration dimension=1
//ff:what data-bind의 dotted path에 optional chaining이 적용되는지 검증
package stml

import (
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestBindOptionalChain(t *testing.T) {
	tests := []struct {
		name    string
		bind    string
		want    string
		notWant string
	}{
		{
			name: "single field unchanged",
			bind: "name",
			want: "{data.name}",
		},
		{
			name:    "dotted path gets optional chaining",
			bind:    "building.name",
			want:    "{data.building?.name}",
			notWant: "{data.building.name}",
		},
		{
			name:    "deeply nested path",
			bind:    "a.b.c",
			want:    "{data.a?.b?.c}",
			notWant: "{data.a.b.c}",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := renderBindJSX(
				stmlparser.FieldBind{Name: tt.bind, Tag: "span"},
				"data", 0, bindCtx{},
			)
			assertContains(t, got, tt.want)
			if tt.notWant != "" {
				assertNotContains(t, got, tt.notWant)
			}
		})
	}
}
