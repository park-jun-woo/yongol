//ff:func feature=validate type=test control=sequence topic=ssac-sqlc
//ff:what TestXQS76_SkipCases — package/nil-result/put/missing/nil 스킵 검증

package ssac_sqlc

import (
	"testing"

	sqlcparser "github.com/park-jun-woo/yongol/pkg/parser/sqlc"
	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXQS76_SkipCases(t *testing.T) {
	t.Run("GetWithPackageSkipped", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssacparser.ServiceFunc{{
				Name: "GetSession", FileName: "getsession.ssac",
				Sequences: []ssacparser.Sequence{{Type: "get", Package: "session", Model: "Session.Get", Result: &ssacparser.Result{Type: "Session", Var: "session"}, Line: 5}},
			}},
			SQLcQueries: []sqlcparser.QuerySpec{{
				Name: "SessionGet", Model: "Session", Method: "Get", Cardinality: "exec",
				Body: "SELECT * FROM sessions WHERE id = @id", File: "sessions.sql", Line: 10,
			}},
		}
		diags := xqs76GetPostExecCardinality(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0 diags, got %d: %v", len(diags), diags)
		}
	})

	t.Run("GetWithNilResultSkipped", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssacparser.ServiceFunc{{
				Name: "GetCourse", FileName: "getcourse.ssac",
				Sequences: []ssacparser.Sequence{{Type: "get", Model: "Course.FindByID", Line: 5}},
			}},
			SQLcQueries: []sqlcparser.QuerySpec{{
				Name: "CourseFindByID", Model: "Course", Method: "FindByID", Cardinality: "exec",
				Body: "SELECT * FROM courses WHERE id = @id", File: "courses.sql", Line: 10,
			}},
		}
		diags := xqs76GetPostExecCardinality(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0 diags, got %d: %v", len(diags), diags)
		}
	})

	t.Run("PutWithExecSkipped", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssacparser.ServiceFunc{{
				Name: "UpdateCourse", FileName: "updatecourse.ssac",
				Sequences: []ssacparser.Sequence{{Type: "put", Model: "Course.Update", Line: 5}},
			}},
			SQLcQueries: []sqlcparser.QuerySpec{{
				Name: "CourseUpdate", Model: "Course", Method: "Update", Cardinality: "exec",
				Body: "UPDATE courses SET title = @title WHERE id = @id", File: "courses.sql", Line: 10,
			}},
		}
		diags := xqs76GetPostExecCardinality(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0 diags, got %d: %v", len(diags), diags)
		}
	})

	t.Run("MissingQuerySkipped", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssacparser.ServiceFunc{{
				Name: "GetCourse", FileName: "getcourse.ssac",
				Sequences: []ssacparser.Sequence{{Type: "get", Model: "Course.FindByID", Result: &ssacparser.Result{Type: "Course", Var: "course"}, Line: 5}},
			}},
			SQLcQueries: []sqlcparser.QuerySpec{},
		}
		diags := xqs76GetPostExecCardinality(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0 diags, got %d: %v", len(diags), diags)
		}
	})

	t.Run("NilFullstackPasses", func(t *testing.T) {
		diags := xqs76GetPostExecCardinality(nil)
		if len(diags) != 0 {
			t.Fatalf("expected 0 diags, got %d: %v", len(diags), diags)
		}
	})
}
