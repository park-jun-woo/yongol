//ff:func feature=gen-ir type=util control=iteration dimension=1
//ff:what annotateGetGuards -- @get 직후 @empty/@exists 가드 감지 → GetOp.FollowedByGuard 주입

package ir

// annotateGetGuards scans the Op list for @get ops that are immediately
// followed by @empty or @exists targeting the same result variable. When
// found, it sets GetOp.FollowedByGuard so the renderer can emit
// errors.Is(err, pgx.ErrNoRows) tolerance instead of plain error propagation.
func annotateGetGuards(ops []Op) {
	for i := 0; i < len(ops)-1; i++ {
		if ops[i].Kind != OpGet || ops[i].Get == nil {
			continue
		}
		varName := ops[i].Get.VarName
		if varName == "" || varName == "_" {
			continue
		}
		guard := matchFollowingGuard(ops[i+1], varName)
		if guard != OpGet {
			ops[i].Get.FollowedByGuard = guard
		}
	}
}
