//ff:func feature=stml-gen type=test control=iteration dimension=1
//ff:what toLabel 함수의 snake_case/PascalCase → Title Case 변환을 검증
package stml

import "testing"

func TestToLabel(t *testing.T) {
	cases := []struct {
		input string
		want  string
	}{
		{"trigger_event", "Trigger Event"},
		{"room_id", "Room Id"},
		{"email", "Email"},
		{"Email", "Email"},
		{"RoomID", "Room ID"},
		{"Password", "Password"},
		{"StartAt", "Start At"},
		{"created_at", "Created At"},
		{"Name", "Name"},
		{"", ""},
	}
	for _, tc := range cases {
		t.Run(tc.input, func(t *testing.T) {
			got := toLabel(tc.input)
			if got != tc.want {
				t.Errorf("toLabel(%q) = %q, want %q", tc.input, got, tc.want)
			}
		})
	}
}
