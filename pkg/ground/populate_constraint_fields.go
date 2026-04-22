//ff:func feature=rule type=loader control=iteration dimension=1
//ff:what populateConstraintFields — 필드 제약조건을 Ground.Types/Schemas에 등록
package ground

import (
	"strconv"
	"strings"

	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
	"github.com/park-jun-woo/yongol/pkg/rule"
)

func populateConstraintFields(g *rule.Ground, prefix string, fields map[string]oapiparser.FieldConstraint) {
	var required []string
	var enumFields []string
	for name, fc := range fields {
		if fc.MaxLength != nil {
			g.Types[prefix+".maxLength."+name] = strconv.Itoa(*fc.MaxLength)
		}
		if fc.Format != "" {
			g.Types[prefix+".format."+name] = fc.Format
		}
		if fc.Required {
			required = append(required, name)
		}
		if len(fc.Enum) > 0 {
			g.Schemas[prefix+".enum."+name] = fc.Enum
			g.Types[prefix+".enum."+name] = strings.Join(fc.Enum, ",")
			enumFields = append(enumFields, name)
		}
	}
	if len(required) > 0 {
		g.Schemas[prefix+".required"] = required
	}
	if len(enumFields) > 0 {
		g.Schemas[prefix+".enumFields"] = enumFields
	}
}
