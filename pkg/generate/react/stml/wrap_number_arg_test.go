//ff:func feature=stml-gen type=test control=iteration dimension=1
//ff:what wrapNumberArg — 항상 평문 Number() 래핑(타입 number, BUG-137) 검증

package stml

import "testing"

func TestWrapNumberArg(t *testing.T) {
	tests := []struct {
		name string
		expr string
		want string
	}{
		{
			name: "required param wraps in Number()",
			expr: "RoomID",
			want: "Number(RoomID)",
		},
		{
			// BUG-137: optional segments no longer widen the arg — the call-guard
			// (enabled / mutation disabled) prevents the empty-value call instead.
			name: "optional param also stays plain Number()",
			expr: "RoomID",
			want: "Number(RoomID)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wrapNumberArg(tt.expr); got != tt.want {
				t.Errorf("wrapNumberArg(%q) = %q, want %q", tt.expr, got, tt.want)
			}
		})
	}
}
