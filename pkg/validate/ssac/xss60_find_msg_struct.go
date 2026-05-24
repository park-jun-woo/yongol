//ff:func feature=validate type=util control=iteration dimension=1 topic=ssac-structural
//ff:what xss60FindMsgStruct — subscriber 함수에서 Param.TypeName에 해당하는 StructInfo 검색

package ssac

import parsessac "github.com/park-jun-woo/yongol/pkg/parser/ssac"

// xss60FindMsgStruct finds the message struct matching fn.Param.TypeName.
func xss60FindMsgStruct(fn parsessac.ServiceFunc) *parsessac.StructInfo {
	for i := range fn.Structs {
		if fn.Structs[i].Name == fn.Param.TypeName {
			return &fn.Structs[i]
		}
	}
	return nil
}
