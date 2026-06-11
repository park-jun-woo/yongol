//ff:func feature=stml-gen type=test control=iteration dimension=1
//ff:what TestBindValueExpr — 타입/포맷별 값 표현식 분기와 미상 타입 fallback(바이트 동일) 검증

package stml

import (
	"testing"

	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
)

func TestBindValueExpr(t *testing.T) {
	tests := []struct {
		name string
		info oapiparser.FieldTypeInfo
		want string
	}{
		{"boolean", oapiparser.FieldTypeInfo{Type: "boolean"}, "{v ? 'Yes' : 'No'}"},
		{"date", oapiparser.FieldTypeInfo{Type: "string", Format: "date"}, "{v ? new Date(v).toLocaleDateString() : ''}"},
		{"date-time", oapiparser.FieldTypeInfo{Type: "string", Format: "date-time"}, "{v ? new Date(v).toLocaleString() : ''}"},
		{"integer", oapiparser.FieldTypeInfo{Type: "integer"}, "{typeof v === 'number' ? v.toLocaleString() : v}"},
		{"number", oapiparser.FieldTypeInfo{Type: "number"}, "{typeof v === 'number' ? v.toLocaleString() : v}"},
		{"plain string", oapiparser.FieldTypeInfo{Type: "string"}, "{v}"},
		{"unknown (zero) falls back byte-identical", oapiparser.FieldTypeInfo{}, "{v}"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bindValueExpr("v", tt.info); got != tt.want {
				t.Errorf("bindValueExpr(%+v) = %q, want %q", tt.info, got, tt.want)
			}
		})
	}
}
