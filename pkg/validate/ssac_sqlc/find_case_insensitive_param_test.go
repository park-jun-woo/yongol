//ff:func feature=validate type=test control=sequence topic=ssac-sqlc
//ff:what TestSSaCSqlcHelpers — unit tests for the pure ssac_sqlc helper functions
package ssac_sqlc

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/rule"
)

func TestFindCaseInsensitiveParam(t *testing.T) {
	params := rule.StringSet{"UserID": true, "email": true}
	if got := findCaseInsensitiveParam("userid", params); got != "UserID" {
		t.Errorf("findCaseInsensitiveParam(userid) = %q, want UserID", got)
	}
	if got := findCaseInsensitiveParam("EMAIL", params); got != "email" {
		t.Errorf("findCaseInsensitiveParam(EMAIL) = %q, want email", got)
	}
	if got := findCaseInsensitiveParam("missing", params); got != "" {
		t.Errorf("findCaseInsensitiveParam(missing) = %q, want empty", got)
	}
}
