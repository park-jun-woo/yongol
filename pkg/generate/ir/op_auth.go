//ff:type feature=gen-ir type=model
//ff:what AuthOp -- @auth 시퀀스의 IR 표현 (OPA Rego 정책 평가)

package ir

// AuthOp represents an @auth sequence: an OPA Rego policy evaluation that
// gates handler execution.
type AuthOp struct {
	// Action is the policy action (e.g. "delete", "AcceptProposal").
	Action string

	// Resource is the policy resource (e.g. "project", "gig").
	Resource string

	// Inputs are the key-value pairs passed to the policy evaluation.
	Inputs []FieldArg

	// Message is the error message returned on authorization failure.
	Message string

	// StatusCode is the HTTP status code on failure (default 403).
	StatusCode int

	// Ownership carries the ownership lookup metadata extracted from
	// Rego @ownership annotations. Nil when no ownership mapping applies.
	Ownership *OwnershipInfo
}
