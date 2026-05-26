//ff:type feature=gen-ir type=model
//ff:what Op -- 13종 SSaC 시퀀스의 태그드 유니온 IR (Kind + 종류별 포인터)

package ir

// Op is a tagged union representing one abstract operation in a ServicePlan.
// Exactly one of the pointer fields is non-nil, matching Kind.
type Op struct {
	Kind     OpKind
	Get      *GetOp
	Post     *PostOp
	Put      *PutOp
	Delete   *DeleteOp
	Empty    *EmptyOp
	Exists   *ExistsOp
	Auth     *AuthOp
	State    *StateOp
	Call     *CallOp
	Eval     *EvalOp
	Publish  *PublishOp
	VerifyPW *VerifyPasswordOp
	Response *ResponseOp
}
