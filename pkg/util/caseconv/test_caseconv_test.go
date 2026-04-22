//ff:func feature=util type=test control=sequence topic=string-convert
//ff:what caseconv — SnakeToPascal / SnakeToPascalSqlc / PascalToSnake / KebabToCamel 회귀 테이블 테스트

package caseconv

import "testing"

func TestSnakeToPascal(t *testing.T) {
	cases := []struct{ in, want string }{
		{"user", "User"},
		{"user_id", "UserId"},
		{"created_at", "CreatedAt"},
		{"one_two_three", "OneTwoThree"},
		{"", ""},
		{"already", "Already"},
		{"__a__", "A"},
		{"per_page", "PerPage"},
	}
	for _, c := range cases {
		t.Run(c.in, func(t *testing.T) {
			if got := SnakeToPascal(c.in); got != c.want {
				t.Errorf("SnakeToPascal(%q) = %q, want %q", c.in, got, c.want)
			}
		})
	}
}

func TestSnakeToPascalSqlc(t *testing.T) {
	cases := []struct{ in, want string }{
		{"id", "ID"},
		{"ids", "IDS"},
		{"org_id", "OrgID"},
		{"user_ids", "UserIDS"},
		{"per_page", "PerPage"},
		{"org_name", "OrgName"},
		{"email", "Email"},
		{"", ""},
	}
	for _, c := range cases {
		t.Run(c.in, func(t *testing.T) {
			if got := SnakeToPascalSqlc(c.in); got != c.want {
				t.Errorf("SnakeToPascalSqlc(%q) = %q, want %q", c.in, got, c.want)
			}
		})
	}
}

func TestPascalToSnake(t *testing.T) {
	cases := []struct{ in, want string }{
		{"UserID", "user_id"},
		{"OrgName", "org_name"},
		{"userName", "user_name"},
		{"HTTPServer", "http_server"},
		// ettle/strcase splits trailing lowercase letter after an uppercase run,
		// so "UserIDs" → "user_i_ds". Documented to pin behaviour for callers.
		{"UserIDs", "user_i_ds"},
	}
	for _, c := range cases {
		t.Run(c.in, func(t *testing.T) {
			if got := PascalToSnake(c.in); got != c.want {
				t.Errorf("PascalToSnake(%q) = %q, want %q", c.in, got, c.want)
			}
		})
	}
}

func TestKebabToCamel(t *testing.T) {
	cases := []struct{ in, want string }{
		{"project-id", "projectId"},
		{"ReservationID", "ReservationID"},
		{"room-id", "roomId"},
		{"a-b-c", "aBC"},
		{"data-fetch", "dataFetch"},
		{"nodash", "nodash"},
		{"", ""},
	}
	for _, c := range cases {
		t.Run(c.in, func(t *testing.T) {
			if got := KebabToCamel(c.in); got != c.want {
				t.Errorf("KebabToCamel(%q) = %q, want %q", c.in, got, c.want)
			}
		})
	}
}
