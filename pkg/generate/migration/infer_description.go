//ff:func feature=migration type=util control=sequence
//ff:what InferDescription — ops 리스트에서 파일명용 snake_case 요약 생성
package migration

// InferDescription returns a snake_case string used as the <desc> in
// artifacts/db/migrations/NNNN_<desc>.sql. Rules:
//   - empty ops        -> "noop"
//   - single op        -> mnemonic derived from op.Description()
//   - all CreateTable  -> "initial"
//   - otherwise        -> "<mnemonic-of-first>_and_more"
func InferDescription(ops []Operation) string {
	if len(ops) == 0 {
		return "noop"
	}
	mnemonic := mnemonicFor(ops[0])
	if len(ops) == 1 {
		return mnemonic
	}
	if allCreateTableOps(ops) {
		return "initial"
	}
	return mnemonic + "_and_more"
}
