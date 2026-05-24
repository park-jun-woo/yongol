//ff:func feature=validate type=test control=sequence topic=ssac-sqlc
//ff:what TestXQS73_EmptyTargetMissingPK — @empty 시 부분 SELECT에 PK 컬럼 누락 에러 검증

package ssac_sqlc

import (
	"strings"
	"testing"

	sqlcparser "github.com/park-jun-woo/yongol/pkg/parser/sqlc"
	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// TestXQS73_EmptyTargetMissingPK verifies that @empty raises an error
// when the partial SELECT query does not include the PK column (id).
func TestXQS73_EmptyTargetMissingPK(t *testing.T) {
	fs := &yongol.Fullstack{
		ServiceFuncs: []ssacparser.ServiceFunc{{
			Name: "RefreshToken", FileName: "refreshtoken.ssac",
			Sequences: []ssacparser.Sequence{
				{
					Type:   "get",
					Model:  "RefreshToken.FindByHash",
					Result: &ssacparser.Result{Type: "RefreshToken", Var: "rt"},
					Line:   5,
				},
				{
					Type:    "empty",
					Target:  "rt",
					Message: "Invalid token",
					Line:    8,
				},
			},
		}},
		SQLcQueries: []sqlcparser.QuerySpec{{
			Name:        "RefreshTokenFindByHash",
			Model:       "RefreshToken",
			Method:      "FindByHash",
			Cardinality: "one",
			RowType:     "RefreshTokenFindByHashRow",
			Body:        "SELECT claims, expires_at, revoked_at FROM refresh_tokens WHERE token_hash = @token_hash",
			SelectStar:  false,
			SelectCols:  []string{"claims", "expires_at", "revoked_at"},
			File:        "refresh_tokens.sql",
			Line:        1,
		}},
	}
	diags := xqs73PartialSelectField(fs)
	if len(diags) != 1 {
		t.Fatalf("expected 1 diag, got %d: %v", len(diags), diags)
	}
	if !strings.Contains(diags[0].Message, "id") {
		t.Fatalf("expected error about 'id' PK column, got: %s", diags[0].Message)
	}
}
