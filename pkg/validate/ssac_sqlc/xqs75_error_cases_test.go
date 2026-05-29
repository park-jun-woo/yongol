//ff:func feature=validate type=test control=sequence topic=ssac-sqlc
//ff:what TestXQS75_ErrorCases — @put/@delete + :one/:many cardinality 불일치 에러 검증

package ssac_sqlc

import (
	"strings"
	"testing"

	sqlcparser "github.com/park-jun-woo/yongol/pkg/parser/sqlc"
	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXQS75_ErrorCases(t *testing.T) {
	t.Run("PutWithOneRaises", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssacparser.ServiceFunc{{
				Name: "UpdateCourse", FileName: "updatecourse.ssac",
				Sequences: []ssacparser.Sequence{{Type: "put", Model: "Course.Update", Line: 5}},
			}},
			SQLcQueries: []sqlcparser.QuerySpec{{
				Name: "CourseUpdate", Model: "Course", Method: "Update", Cardinality: "one",
				Body: "UPDATE courses SET title = @title WHERE id = @id RETURNING *", File: "courses.sql", Line: 10,
			}},
		}
		diags := xqs75PutDeleteExecCardinality(fs)
		if len(diags) != 1 {
			t.Fatalf("expected 1 diag, got %d: %v", len(diags), diags)
		}
		if !strings.Contains(diags[0].Message, "XQS-75") {
			t.Fatalf("expected XQS-75, got: %s", diags[0].Message)
		}
		if !strings.Contains(diags[0].Message, "@put") {
			t.Fatalf("expected @put, got: %s", diags[0].Message)
		}
	})

	t.Run("DeleteWithManyRaises", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssacparser.ServiceFunc{{
				Name: "DeleteExpired", FileName: "deleteexpired.ssac",
				Sequences: []ssacparser.Sequence{{Type: "delete", Model: "Token.DeleteExpired", Line: 3}},
			}},
			SQLcQueries: []sqlcparser.QuerySpec{{
				Name: "TokenDeleteExpired", Model: "Token", Method: "DeleteExpired", Cardinality: "many",
				Body: "DELETE FROM tokens WHERE expires_at < NOW() RETURNING id", File: "tokens.sql", Line: 5,
			}},
		}
		diags := xqs75PutDeleteExecCardinality(fs)
		if len(diags) != 1 {
			t.Fatalf("expected 1 diag, got %d: %v", len(diags), diags)
		}
		if !strings.Contains(diags[0].Message, "@delete") {
			t.Fatalf("expected @delete, got: %s", diags[0].Message)
		}
	})
}
