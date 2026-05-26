//ff:func feature=validate type=test control=sequence dimension=2 topic=ssac-structural
//ff:what S-63 — @get []T 리스트 엔드포인트에 페이지네이션 파라미터 누락 시 경고

package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestS63ListNoPagination(t *testing.T) {
	t.Run("Fires_list_without_pagination", func(t *testing.T) {
		fs := &yongol.Fullstack{ServiceFuncs: []ssac.ServiceFunc{{
			Name: "ListCourses", FileName: "c.ssac", Line: 1,
			Sequences: []ssac.Sequence{{
				Type: "get", Line: 3, Model: "Course.ListAll",
				Result: &ssac.Result{Var: "courses", Type: "[]Course"},
			}},
		}}}
		assertDiag(t, s63ListNoPagination(fs), "[S-63]")
	})
	t.Run("Passes_with_page_key", func(t *testing.T) {
		fs := &yongol.Fullstack{ServiceFuncs: []ssac.ServiceFunc{{
			Name: "ListCourses", FileName: "c.ssac", Line: 1,
			Sequences: []ssac.Sequence{{
				Type: "get", Line: 3, Model: "Course.ListAll",
				Result: &ssac.Result{Var: "courses", Type: "[]Course"},
				Inputs: map[string]string{"Page": "query.Page", "PerPage": "query.PerPage"},
			}},
		}}}
		assertNoDiag(t, s63ListNoPagination(fs), "[S-63]")
	})
	t.Run("Passes_with_cursor_key", func(t *testing.T) {
		fs := &yongol.Fullstack{ServiceFuncs: []ssac.ServiceFunc{{
			Name: "ListCourses", FileName: "c.ssac", Line: 1,
			Sequences: []ssac.Sequence{{
				Type: "get", Line: 3, Model: "Course.ListAll",
				Result: &ssac.Result{Var: "courses", Type: "[]Course"},
				Inputs: map[string]string{"Cursor": "query.Cursor"},
			}},
		}}}
		assertNoDiag(t, s63ListNoPagination(fs), "[S-63]")
	})
	t.Run("Passes_no_pagination_annotation", func(t *testing.T) {
		fs := &yongol.Fullstack{ServiceFuncs: []ssac.ServiceFunc{{
			Name: "ListAll", FileName: "c.ssac", Line: 1, NoPagination: true,
			Sequences: []ssac.Sequence{{
				Type: "get", Line: 3, Model: "Course.ListAll",
				Result: &ssac.Result{Var: "courses", Type: "[]Course"},
			}},
		}}}
		assertNoDiag(t, s63ListNoPagination(fs), "[S-63]")
	})
	t.Run("Skips_non_list_result", func(t *testing.T) {
		fs := &yongol.Fullstack{ServiceFuncs: []ssac.ServiceFunc{{
			Name: "GetCourse", FileName: "c.ssac", Line: 1,
			Sequences: []ssac.Sequence{{
				Type: "get", Line: 3, Model: "Course.Find",
				Result: &ssac.Result{Var: "course", Type: "Course"},
			}},
		}}}
		assertNoDiag(t, s63ListNoPagination(fs), "[S-63]")
	})
	t.Run("Skips_non_get", func(t *testing.T) {
		fs := &yongol.Fullstack{ServiceFuncs: []ssac.ServiceFunc{{
			Name: "X", FileName: "x.ssac", Line: 1,
			Sequences: []ssac.Sequence{{
				Type: "post", Line: 3, Model: "M.Create",
				Result: &ssac.Result{Var: "items", Type: "[]M"},
			}},
		}}}
		assertNoDiag(t, s63ListNoPagination(fs), "[S-63]")
	})
	t.Run("Skips_nil_result", func(t *testing.T) {
		fs := &yongol.Fullstack{ServiceFuncs: []ssac.ServiceFunc{{
			Name: "X", FileName: "x.ssac", Line: 1,
			Sequences: []ssac.Sequence{{
				Type: "get", Line: 3, Model: "M.Find",
			}},
		}}}
		assertNoDiag(t, s63ListNoPagination(fs), "[S-63]")
	})
	t.Run("Empty_funcs", func(t *testing.T) {
		assertNoDiag(t, s63ListNoPagination(&yongol.Fullstack{}), "[S-63]")
	})
}
