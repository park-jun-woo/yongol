//ff:func feature=gen-gogin type=generator control=sequence
//ff:what writeConvertListFunc — convert<Name>List([]db.X) []api.X 헬퍼 함수 렌더링

package ssac

import "strings"

// writeConvertListFunc generates: func convertWorkflowList(rows []db.Workflow) []api.Workflow
func writeConvertListFunc(sb *strings.Builder, name string) {
	sb.WriteString("\nfunc convert" + name + "List(rows []db." + name + ") []api." + name + " {\n")
	sb.WriteString("\tresult := make([]api." + name + ", len(rows))\n")
	sb.WriteString("\tfor i, row := range rows {\n")
	sb.WriteString("\t\tresult[i] = *convert" + name + "(row)\n")
	sb.WriteString("\t}\n")
	sb.WriteString("\treturn result\n")
	sb.WriteString("}\n")
}
