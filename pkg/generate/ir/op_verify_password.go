//ff:type feature=gen-ir type=model
//ff:what VerifyPasswordOp -- @verify-password 시퀀스의 IR 표현 (timing-safe 비밀번호 검증)

package ir

// VerifyPasswordOp represents a @verify-password sequence: a timing-safe
// password verification that looks up a user by email, compares the password
// hash, and binds the result on success.
type VerifyPasswordOp struct {
	// Model is the sqlc table/model name (e.g. "User").
	Model string

	// Method is the sqlc lookup method (e.g. "FindByEmail").
	Method string

	// EmailCol is the email column name (e.g. "Email").
	EmailCol string

	// EmailExpr is the expression resolving the email value.
	EmailExpr string

	// HashCol is the password hash column name (e.g. "PasswordHash").
	HashCol string

	// PasswordExpr is the expression resolving the plain-text password.
	PasswordExpr string

	// ResultVar is the bound variable on success (e.g. "user").
	ResultVar string

	// ResultType is the result type name (e.g. "User").
	ResultType string

	// ErrStatus is the HTTP status code on failure.
	ErrStatus int

	// Message is the error message on verification failure.
	Message string
}
