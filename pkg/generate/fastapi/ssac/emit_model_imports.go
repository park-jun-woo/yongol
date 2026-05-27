//ff:func feature=gen-fastapi type=util control=iteration dimension=1
//ff:what emitModelImports — SQLAlchemy 모델 클래스 import 출력

package ssac

import (
	"fmt"
	"sort"
	"strings"
)

// emitModelImports writes the model class imports.
func emitModelImports(b *strings.Builder, models map[string]bool) {
	if len(models) == 0 {
		return
	}
	sorted := make([]string, 0, len(models))
	for m := range models {
		sorted = append(sorted, m)
	}
	sort.Strings(sorted)
	b.WriteString(fmt.Sprintf("from app.models.models import %s\n", strings.Join(sorted, ", ")))
}
