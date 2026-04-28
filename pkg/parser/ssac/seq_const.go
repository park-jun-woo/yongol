//ff:type feature=ssac-parse type=model
//ff:what sequence type constants and validity map
package ssac

// sequence type constants
const (
	SeqGet      = "get"
	SeqPost     = "post"
	SeqPut      = "put"
	SeqDelete   = "delete"
	SeqEmpty    = "empty"
	SeqExists   = "exists"
	SeqState    = "state"
	SeqAuth     = "auth"
	SeqCall           = "call"
	SeqEval           = "eval"
	SeqPublish        = "publish"
	SeqResponse       = "response"
	SeqVerifyPassword = "verify-password"
)

var ValidSequenceTypes = map[string]bool{
	SeqGet:      true,
	SeqPost:     true,
	SeqPut:      true,
	SeqDelete:   true,
	SeqEmpty:    true,
	SeqExists:   true,
	SeqState:    true,
	SeqAuth:     true,
	SeqCall:           true,
	SeqEval:           true,
	SeqPublish:        true,
	SeqResponse:       true,
	SeqVerifyPassword: true,
}
