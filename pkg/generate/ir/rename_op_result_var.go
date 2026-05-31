//ff:func feature=gen-ir type=util control=selection
//ff:what renameOpResultVar -- Op 종류별 result 변수명을 declared 집합과 대조해 충돌 해소 + rename 등록

package ir

// renameOpResultVar resolves the result variable name declared by op against
// the declared set, recording any rename in renames. The op kind selects which
// result-variable field is examined.
func renameOpResultVar(op *Op, declared map[string]bool, renames map[string]string) {
	switch op.Kind {
	case OpGet:
		if op.Get != nil {
			assignResolvedVar(&op.Get.VarName, declared, renames)
		}
	case OpPost:
		if op.Post != nil {
			assignResolvedVar(&op.Post.VarName, declared, renames)
		}
	case OpCall:
		if op.Call != nil {
			assignResolvedVar(&op.Call.ResultVar, declared, renames)
		}
	case OpVerifyPassword:
		if op.VerifyPW != nil {
			assignResolvedVar(&op.VerifyPW.ResultVar, declared, renames)
		}
	}
}
