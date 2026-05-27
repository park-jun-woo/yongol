//ff:func feature=gen-ir type=util control=selection
//ff:what collectFieldArgSlices -- Op 내 모든 FieldArg 슬라이스 포인터를 수집 (enrichment 패스 공용)

package ir

// collectFieldArgSlices returns pointers to all FieldArg slices within an Op
// so enrichment passes can iterate uniformly.
func collectFieldArgSlices(op *Op) []*[]FieldArg {
	switch op.Kind {
	case OpGet:
		if op.Get != nil {
			return []*[]FieldArg{&op.Get.Args, &op.Get.PaginationArgs}
		}
	case OpPost:
		if op.Post != nil {
			return []*[]FieldArg{&op.Post.Args}
		}
	case OpPut:
		if op.Put != nil {
			return []*[]FieldArg{&op.Put.Args}
		}
	case OpDelete:
		if op.Delete != nil {
			return []*[]FieldArg{&op.Delete.Args}
		}
	case OpAuth:
		if op.Auth != nil {
			return []*[]FieldArg{&op.Auth.Inputs}
		}
	case OpState:
		if op.State != nil {
			return []*[]FieldArg{&op.State.Inputs}
		}
	case OpCall:
		if op.Call != nil {
			return []*[]FieldArg{&op.Call.Args}
		}
	case OpEval:
		if op.Eval != nil {
			return []*[]FieldArg{&op.Eval.Args}
		}
	case OpPublish:
		if op.Publish != nil {
			return []*[]FieldArg{&op.Publish.Payload, &op.Publish.Options}
		}
	}
	return nil
}
