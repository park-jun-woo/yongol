//ff:func feature=stml-gen type=test control=iteration dimension=1
//ff:what wrapNumberArg — required는 Number() 래핑/optional은 null 가드 포함 분기 검증

package stml

import "testing"

func TestWrapNumberArg(t *testing.T) {
	tests := []struct {
		name     string
		expr     string
		optional bool
		want     string
	}{
		{
			name:     "required param wraps in Number()",
			expr:     "RoomID",
			optional: false,
			want:     "Number(RoomID)",
		},
		{
			name:     "optional param adds null guard",
			expr:     "RoomID",
			optional: true,
			want:     "RoomID != null ? Number(RoomID) : undefined",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wrapNumberArg(tt.expr, tt.optional); got != tt.want {
				t.Errorf("wrapNumberArg(%q, %v) = %q, want %q", tt.expr, tt.optional, got, tt.want)
			}
		})
	}
}
