//ff:func feature=validate type=test control=iteration dimension=1 topic=stml-statemachine
//ff:what stateInSlice — 존재/부재·빈 슬라이스·대소문자 구분·중복 포함 여부 검증

package stml_statemachine

import "testing"

func TestStateInSlice(t *testing.T) {
	states := []string{"draft", "active", "archived"}
	tests := []struct {
		name   string
		states []string
		target string
		want   bool
	}{
		{name: "present in middle", states: states, target: "active", want: true},
		{name: "present at end", states: states, target: "archived", want: true},
		{name: "absent", states: states, target: "pending", want: false},
		{name: "case sensitive miss", states: states, target: "Draft", want: false},
		{name: "empty slice", states: nil, target: "draft", want: false},
		{name: "empty target absent", states: states, target: "", want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stateInSlice(tt.states, tt.target); got != tt.want {
				t.Errorf("stateInSlice(%v, %q) = %v, want %v", tt.states, tt.target, got, tt.want)
			}
		})
	}
}
