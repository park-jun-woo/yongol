//ff:func feature=validate type=test control=sequence topic=ssac-sqlc
//ff:what TestXQS75_SkipCases — @get/package/missing/nil 스킵 검증

package ssac_sqlc

import (
	"testing"

	sqlcparser "github.com/park-jun-woo/yongol/pkg/parser/sqlc"
	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXQS75_SkipCases(t *testing.T) {
	t.Run("GetWithOneSkipped", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssacparser.ServiceFunc{{
				Name: "GetCourse", FileName: "getcourse.ssac",
				Sequences: []ssacparser.Sequence{{Type: "get", Model: "Course.FindByID", Result: &ssacparser.Result{Type: "Course", Var: "course"}, Line: 5}},
			}},
			SQLcQueries: []sqlcparser.QuerySpec{{
				Name: "CourseFindByID", Model: "Course", Method: "FindByID", Cardinality: "one",
				Body: "SELECT * FROM courses WHERE id = @id", File: "courses.sql", Line: 10,
			}},
		}
		diags := xqs75PutDeleteExecCardinality(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0 diags, got %d: %v", len(diags), diags)
		}
	})

	t.Run("PutWithPackageSkipped", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssacparser.ServiceFunc{{
				Name: "UpdateSession", FileName: "updatesession.ssac",
				Sequences: []ssacparser.Sequence{{Type: "put", Package: "session", Model: "Session.Extend", Line: 5}},
			}},
			SQLcQueries: []sqlcparser.QuerySpec{{
				Name: "SessionExtend", Model: "Session", Method: "Extend", Cardinality: "one",
				Body: "UPDATE sessions SET expires_at = NOW() RETURNING *", File: "sessions.sql", Line: 10,
			}},
		}
		diags := xqs75PutDeleteExecCardinality(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0 diags, got %d: %v", len(diags), diags)
		}
	})

	t.Run("MissingQuerySkipped", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssacparser.ServiceFunc{{
				Name: "UpdateCourse", FileName: "updatecourse.ssac",
				Sequences: []ssacparser.Sequence{{Type: "put", Model: "Course.Update", Line: 5}},
			}},
			SQLcQueries: []sqlcparser.QuerySpec{},
		}
		diags := xqs75PutDeleteExecCardinality(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0 diags, got %d: %v", len(diags), diags)
		}
	})

	t.Run("NilFullstackPasses", func(t *testing.T) {
		diags := xqs75PutDeleteExecCardinality(nil)
		if len(diags) != 0 {
			t.Fatalf("expected 0 diags, got %d: %v", len(diags), diags)
		}
	})
}
