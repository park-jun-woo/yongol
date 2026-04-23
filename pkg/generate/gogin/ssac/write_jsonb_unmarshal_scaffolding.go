//ff:func feature=gen-gogin type=util control=iteration dimension=1
//ff:what writeJSONBUnmarshalScaffolding — 각 JSONB 필드에 대해 Unmarshal 블록을 strings.Builder 에 기록

package ssac

import "strings"

func writeJSONBUnmarshalScaffolding(sb *strings.Builder, jsonbs []jsonbFieldAlias) {
	for _, j := range jsonbs {
		sb.WriteString("\tvar " + j.localVar + " map[string]interface{}\n")
		sb.WriteString("\tif len(row." + j.dbField + ") > 0 {\n")
		sb.WriteString("\t\tif err := json.Unmarshal(row." + j.dbField + ", &" + j.localVar + "); err != nil {\n")
		sb.WriteString("\t\t\treturn nil, err\n")
		sb.WriteString("\t\t}\n")
		sb.WriteString("\t}\n")
	}
}
