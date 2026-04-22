//ff:func feature=rule type=loader control=sequence
//ff:what registerSSaCModelRef — SSaC get/post/put/delete 시퀀스의 model + DDL table 이름 등록
package ground

import (
	"strings"

	"github.com/jinzhu/inflection"

	"github.com/park-jun-woo/yongol/pkg/rule"
)

// registerSSaCModelRef extracts the model name from "Model.field" form and
// registers both the Go model name and the lowercase plural (= DDL table name).
func registerSSaCModelRef(seqModel string, modelRefs rule.StringSet) {
	idx := strings.IndexByte(seqModel, '.')
	if idx <= 0 {
		return
	}
	model := seqModel[:idx]
	modelRefs[model] = true
	// DDL table name = lowercase plural of model name
	modelRefs[strings.ToLower(inflection.Plural(model))] = true
}
