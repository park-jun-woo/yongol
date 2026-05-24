//ff:func feature=validate type=util control=sequence topic=ssac-sqlc
//ff:what checkXqs73Response — @response 시퀀스에서 부분 SELECT 필드 참조 위반 검사

package ssac_sqlc

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// checkXqs73Response checks @response sequences for partial SELECT field violations.
func checkXqs73Response(fn ssacparser.ServiceFunc, seq ssacparser.Sequence, vars map[string]xqs73VarInfo) []diagnostic.Diagnostic {
	// @response { field: var.Field, ... } — check each field reference
	if len(seq.Fields) > 0 {
		return checkXqs73ResponseFields(fn, seq, vars)
	}

	// @response target — the convert<Model> function accesses ALL schema fields.
	if seq.Target == "" {
		return nil
	}
	vi, ok := vars[seq.Target]
	if !ok {
		return nil
	}
	return []diagnostic.Diagnostic{{
		File:    fn.FileName,
		Line:    seq.Line,
		Phase:   diagnostic.PhaseValidate,
		Level:   diagnostic.LevelError,
		Message: fmt.Sprintf("[XQS-73] @response target %q is bound to query %q which uses partial SELECT (%s); convert<%s>() will access fields not present in the Row struct", seq.Target, vi.query.Name, strings.Join(vi.query.SelectCols, ", "), vi.seq.Result.Type),
		Advice:  fmt.Sprintf("Use SELECT * in query %q, or switch to @response { field: %s.Col } to reference only available columns.", vi.query.Name, seq.Target),
	}}
}
