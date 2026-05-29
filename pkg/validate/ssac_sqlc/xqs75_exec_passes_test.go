//ff:func feature=validate type=test control=sequence topic=ssac-sqlc
//ff:what TestXQS75_ExecPasses — @put/@delete + :exec/:execresult 정상 통과 검증

package ssac_sqlc

import (
	"testing"

	sqlcparser "github.com/park-jun-woo/yongol/pkg/parser/sqlc"
	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXQS75_ExecPasses(t *testing.T) {
	t.Run("PutWithExecPasses", func(t *testing.T) {
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
		diags := xqs75PutDeleteExecCardinality(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0 diags, got %d: %v", len(diags), diags)
		}
	})

	t.Run("DeleteWithExecPasses", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssacparser.ServiceFunc{{
				Name: "DeleteCourse", FileName: "deletecourse.ssac",
				Sequences: []ssacparser.Sequence{{Type: "delete", Model: "Course.Delete", Line: 5}},
			}},
			SQLcQueries: []sqlcparser.QuerySpec{{
				Name: "CourseDelete", Model: "Course", Method: "Delete", Cardinality: "exec",
				Body: "DELETE FROM courses WHERE id = @id", File: "courses.sql", Line: 10,
			}},
		}
		diags := xqs75PutDeleteExecCardinality(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0 diags, got %d: %v", len(diags), diags)
		}
	})

	t.Run("PutWithExecresultPasses", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssacparser.ServiceFunc{{
				Name: "UpdateCourse", FileName: "updatecourse.ssac",
				Sequences: []ssacparser.Sequence{{Type: "put", Model: "Course.Update", Line: 5}},
			}},
			SQLcQueries: []sqlcparser.QuerySpec{{
				Name: "CourseUpdate", Model: "Course", Method: "Update", Cardinality: "execresult",
				Body: "UPDATE courses SET title = @title WHERE id = @id", File: "courses.sql", Line: 10,
			}},
		}
		diags := xqs75PutDeleteExecCardinality(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0 diags, got %d: %v", len(diags), diags)
		}
	})
}
