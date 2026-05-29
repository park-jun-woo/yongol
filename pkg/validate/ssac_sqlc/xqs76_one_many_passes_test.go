//ff:func feature=validate type=test control=sequence topic=ssac-sqlc
//ff:what TestXQS76_OneManyPasses — @get/:one, @post/:many 정상 통과 검증

package ssac_sqlc

import (
	"testing"

	sqlcparser "github.com/park-jun-woo/yongol/pkg/parser/sqlc"
	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXQS76_OneManyPasses(t *testing.T) {
	t.Run("GetWithOnePasses", func(t *testing.T) {
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
		diags := xqs76GetPostExecCardinality(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0 diags, got %d: %v", len(diags), diags)
		}
	})

	t.Run("PostWithManyPasses", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssacparser.ServiceFunc{{
				Name: "CreateCourses", FileName: "createcourses.ssac",
				Sequences: []ssacparser.Sequence{{Type: "post", Model: "Course.CreateBatch", Result: &ssacparser.Result{Type: "[]Course", Var: "courses"}, Line: 5}},
			}},
			SQLcQueries: []sqlcparser.QuerySpec{{
				Name: "CourseCreateBatch", Model: "Course", Method: "CreateBatch", Cardinality: "many",
				Body: "INSERT INTO courses (title) VALUES (@title) RETURNING *", File: "courses.sql", Line: 10,
			}},
		}
		diags := xqs76GetPostExecCardinality(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0 diags, got %d: %v", len(diags), diags)
		}
	})
}
