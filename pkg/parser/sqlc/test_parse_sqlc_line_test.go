//ff:func feature=orchestrator type=test control=iteration dimension=1
//ff:what parseSQLCLine 테이블 테스트 — :one/:many/:exec/:execresult 및 비매크로 라인

package sqlc

import "testing"

func TestParseSQLCLine_Table(t *testing.T) {
	cases := []parseSQLCLineCase{
		{"one", "-- name: GetUser :one", true, "GetUser", "one", "GetUserRow"},
		{"many", "-- name: ListUsers :many", true, "ListUsers", "many", "ListUsersRow"},
		{"exec", "-- name: DeleteUser :exec", true, "DeleteUser", "exec", ""},
		{"execresult", "-- name: UpdateUser :execresult", true, "UpdateUser", "execresult", ""},
		{"not-macro", "-- just a comment", false, "", "", ""},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			runParseSQLCLineCase(t, tc)
		})
	}
}
