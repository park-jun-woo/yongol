//ff:type feature=gen-ir type=model
//ff:what OpKind -- 13종 SSaC 시퀀스에 대응하는 IR 연산 종류 enum

package ir

// OpKind enumerates the 13 abstract operation types that correspond to SSaC
// sequence annotations.
type OpKind int

const (
	OpGet            OpKind = iota // @get
	OpPost                         // @post
	OpPut                          // @put
	OpDelete                       // @delete
	OpEmpty                        // @empty
	OpExists                       // @exists
	OpAuth                         // @auth
	OpState                        // @state
	OpCall                         // @call
	OpEval                         // @eval
	OpPublish                      // @publish
	OpVerifyPassword               // @verify-password
	OpResponse                     // @response
)
