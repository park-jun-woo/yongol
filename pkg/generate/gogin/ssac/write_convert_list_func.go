//ff:func feature=gen-gogin type=generator control=sequence
//ff:what writeConvertListFunc — convert<Name>List([]db.X) ([]api.X, error) 헬퍼 함수 렌더링

package ssac

import "strings"

// writeConvertListFunc generates: func convertWorkflowList(rows []db.Workflow) ([]api.Workflow, error)
//
// Mirrors the out-of-file emitConvertListFile variant. Kept for callers
// that assemble multiple converters into a single builder before
// writing. Signature returns an error because convert<Name> can fail
// on JSONB unmarshal (BUG-005).
func writeConvertListFunc(sb *strings.Builder, name string) {
	sb.WriteString("\nfunc convert" + name + "List(rows []db." + name + ") ([]api." + name + ", error) {\n")
	sb.WriteString("\tresult := make([]api." + name + ", len(rows))\n")
	sb.WriteString("\tfor i, row := range rows {\n")
	sb.WriteString("\t\titem, err := convert" + name + "(row)\n")
	sb.WriteString("\t\tif err != nil {\n")
	sb.WriteString("\t\t\treturn nil, err\n")
	sb.WriteString("\t\t}\n")
	sb.WriteString("\t\tresult[i] = *item\n")
	sb.WriteString("\t}\n")
	sb.WriteString("\treturn result, nil\n")
	sb.WriteString("}\n")
}
