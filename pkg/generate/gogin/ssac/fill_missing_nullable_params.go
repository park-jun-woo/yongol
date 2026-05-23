//ff:func feature=gen-gogin type=util control=iteration dimension=1
//ff:what fillMissingNullableParams — SSaC Inputs 에 없는 nullable sqlc Params 를 pgtype zero 값으로 채움
package ssac

import "github.com/park-jun-woo/yongol/pkg/generate/gogin/types"

// fillMissingNullableParams finds sqlc Params fields that the SSaC Inputs
// do not provide and, when the underlying DDL column is nullable (pgtype
// family), emits a zero-value literal (e.g. "pgtype.Int8{}").
func (g *methodGen) fillMissingNullableParams(method string, inputs map[string]string) ([]string, []string) {
	var params []string
	for _, q := range g.SQLcQueries {
		if q.Name == method {
			params = q.Params
			break
		}
	}
	if len(params) == 0 {
		return nil, nil
	}

	model := g.modelForSQLCMethod(method)
	if model == "" {
		return nil, nil
	}

	var fields []string
	var imports []string
	needPgtype := false

	for _, param := range params {
		if _, provided := inputs[param]; provided {
			continue
		}
		col := lookupDDLColumn(g.DDLTables, model, param)
		if col == nil {
			continue
		}
		binding := types.MapPGType(*col)
		if binding.Kind != types.KindPgtype {
			continue
		}
		fields = append(fields, param+": "+binding.SqlcGoType+"{}")
		needPgtype = true
	}

	if needPgtype {
		imports = append(imports, `"github.com/jackc/pgx/v5/pgtype"`)
	}
	return fields, imports
}
