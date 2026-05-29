//ff:func feature=validate type=test control=sequence topic=ssac-sqlc
//ff:what TestXQS76_ErrorCases — @get/@post + :exec/:execrows/:execresult cardinality 불일치 에러 검증

package ssac_sqlc

import (
	"strings"
	"testing"

	sqlcparser "github.com/park-jun-woo/yongol/pkg/parser/sqlc"
	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXQS76_ErrorCases(t *testing.T) {
	t.Run("GetWithExecRaises", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssacparser.ServiceFunc{{
				Name: "GetCourse", FileName: "getcourse.ssac",
				Sequences: []ssacparser.Sequence{{Type: "get", Model: "Course.FindByID", Result: &ssacparser.Result{Type: "Course", Var: "course"}, Line: 5}},
			}},
			SQLcQueries: []sqlcparser.QuerySpec{{
				Name: "CourseFindByID", Model: "Course", Method: "FindByID", Cardinality: "exec",
				Body: "SELECT * FROM courses WHERE id = @id", File: "courses.sql", Line: 10,
			}},
		}
		diags := xqs76GetPostExecCardinality(fs)
		if len(diags) != 1 {
			t.Fatalf("expected 1 diag, got %d: %v", len(diags), diags)
		}
		if !strings.Contains(diags[0].Message, "XQS-76") {
			t.Fatalf("expected XQS-76, got: %s", diags[0].Message)
		}
		if !strings.Contains(diags[0].Message, "@get") {
			t.Fatalf("expected @get, got: %s", diags[0].Message)
		}
	})

	t.Run("PostWithExecrowsRaises", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssacparser.ServiceFunc{{
				Name: "CreateCourse", FileName: "createcourse.ssac",
				Sequences: []ssacparser.Sequence{{Type: "post", Model: "Course.Create", Result: &ssacparser.Result{Type: "Course", Var: "course"}, Line: 5}},
			}},
			SQLcQueries: []sqlcparser.QuerySpec{{
				Name: "CourseCreate", Model: "Course", Method: "Create", Cardinality: "execrows",
				Body: "INSERT INTO courses (title) VALUES (@title)", File: "courses.sql", Line: 10,
			}},
		}
		diags := xqs76GetPostExecCardinality(fs)
		if len(diags) != 1 {
			t.Fatalf("expected 1 diag, got %d: %v", len(diags), diags)
		}
		if !strings.Contains(diags[0].Message, "@post") {
			t.Fatalf("expected @post, got: %s", diags[0].Message)
		}
	})

	t.Run("GetWithExecresultRaises", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssacparser.ServiceFunc{{
				Name: "GetCourse", FileName: "getcourse.ssac",
				Sequences: []ssacparser.Sequence{{Type: "get", Model: "Course.FindByID", Result: &ssacparser.Result{Type: "Course", Var: "course"}, Line: 5}},
			}},
			SQLcQueries: []sqlcparser.QuerySpec{{
				Name: "CourseFindByID", Model: "Course", Method: "FindByID", Cardinality: "execresult",
				Body: "SELECT * FROM courses WHERE id = @id", File: "courses.sql", Line: 10,
			}},
		}
		diags := xqs76GetPostExecCardinality(fs)
		if len(diags) != 1 {
			t.Fatalf("expected 1 diag, got %d: %v", len(diags), diags)
		}
		if !strings.Contains(diags[0].Message, ":execresult") {
			t.Fatalf("expected :execresult, got: %s", diags[0].Message)
		}
	})
}
