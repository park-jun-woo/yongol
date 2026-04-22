//ff:func feature=validate type=util control=sequence topic=ssac-ddl
//ff:what collectReferencedTable — 단일 시퀀스에서 참조된 DDL 테이블을 set 에 반영

package ssac_ddl

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func collectReferencedTable(seq ssac.Sequence, tables map[string]bool) {
	if seq.Package != "" {
		return
	}
	if seq.Model != "" {
		parts := strings.SplitN(seq.Model, ".", 2)
		if len(parts) >= 1 && parts[0] != "" {
			tables[modelToTable(parts[0])] = true
		}
	}
	if seq.Result != nil && seq.Result.Type != "" {
		typeName := normalizeTypeName(seq.Result.Type)
		if typeName != "" && !primitiveTypes[typeName] {
			tables[modelToTable(typeName)] = true
		}
	}
}
