//ff:func feature=gen-ir type=util control=selection
//ff:what opModelName -- Op 에서 모델명 추출 (Get/Post/Put/Delete)

package ir

// opModelName extracts the model name from an Op (Get/Post/Put/Delete).
func opModelName(op *Op) string {
	switch op.Kind {
	case OpGet:
		if op.Get != nil {
			return op.Get.Model
		}
	case OpPost:
		if op.Post != nil {
			return op.Post.Model
		}
	case OpPut:
		if op.Put != nil {
			return op.Put.Model
		}
	case OpDelete:
		if op.Delete != nil {
			return op.Delete.Model
		}
	}
	return ""
}
