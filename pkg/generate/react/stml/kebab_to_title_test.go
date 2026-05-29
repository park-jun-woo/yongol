//ff:func feature=stml-gen type=test control=iteration dimension=1
//ff:what kebabToTitle 함수의 케밥케이스 → Title Case 변환을 검증
package stml

import "testing"

func TestKebabToTitle(t *testing.T) {
	cases := []struct {
		input string
		want  string
	}{
		{"workflows", "Workflows"},
		{"workflow-detail", "Workflow Detail"},
		{"audit-logs", "Audit Logs"},
		{"login-page", "Login"},
		{"my-reservations-page", "My Reservations"},
		{"room-edit-page", "Room Edit"},
		{"reservation-detail-page", "Reservation Detail"},
		{"activate-page", "Activate"},
		{"create-workflow-page", "Create Workflow"},
		{"list-page", "List"},
		{"form-page", "Form"},
		{"action-page", "Action"},
		{"edit-page", "Edit"},
	}
	for _, tc := range cases {
		t.Run(tc.input, func(t *testing.T) {
			got := kebabToTitle(tc.input)
			if got != tc.want {
				t.Errorf("kebabToTitle(%q) = %q, want %q", tc.input, got, tc.want)
			}
		})
	}
}
