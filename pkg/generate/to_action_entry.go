//ff:func feature=generate type=util control=iteration dimension=1
//ff:what ActionBlock에서 actionEntry로 변환한다
package generate

import stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"

// toActionEntry converts an ActionBlock to an actionEntry.
func toActionEntry(a stmlparser.ActionBlock) actionEntry {
	names := make([]string, len(a.Fields))
	for i, f := range a.Fields {
		names[i] = f.Name
	}
	return actionEntry{opID: a.OperationID, fieldNames: names}
}
