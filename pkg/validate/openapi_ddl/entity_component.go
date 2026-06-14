//ff:func feature=validate type=util control=sequence topic=openapi-ddl
//ff:what entityComponent — 모델명이 실제 DDL 테이블+component 에 매핑될 때만 canonical component 명 반환 (B-2 가드)

package openapi_ddl

// entityComponent is the strategy B-2 guard. A model name is accepted as a
// canonical entity only when it maps to a real DDL table AND that table has a
// corresponding components.schemas entry. This rejects single-converging but
// non-entity types (result / token / sessionResult / func-result structs).
// Returns the component schema name, or "" when not a DDL-backed entity.
func entityComponent(idx *entityIndex, model string) string {
	if model == "" {
		return ""
	}
	table := modelToTable(model)
	if !idx.tableExists[table] {
		return ""
	}
	return idx.schemaForTable[table]
}
