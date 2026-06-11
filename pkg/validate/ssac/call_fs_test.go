//ff:func feature=validate type=test-helper control=sequence topic=ssac-structural
//ff:what callFS — 단일 @call Arg를 가진 최소 Fullstack 생성 (S-59 테스트용)

package ssac

import (
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/rule"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// callFS builds a Fullstack with a single @call whose Arg references
// source.field, plus the given S-59 schema for that source.
func callFS(funcName, source, field string, schema map[string][]string) *yongol.Fullstack {
	fs := &yongol.Fullstack{ServiceFuncs: []ssac.ServiceFunc{{
		Name: funcName, FileName: "r.ssac", Line: 1,
		Sequences: []ssac.Sequence{{
			Type: "call", Line: 3, Model: "pkg.Func",
			Args: []ssac.Arg{{Source: source, Field: field}},
		}},
	}}}
	fs.SetGround(&rule.Ground{Schemas: schema, Types: map[string]string{}})
	return fs
}
