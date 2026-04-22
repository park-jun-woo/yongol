//ff:type feature=ssac-parse type=model
//ff:what Sequence — type representing a single SSaC sequence line
package ssac

// Sequence represents a single sequence line.
type Sequence struct {
	Type string // "get", "post", "put", "delete", "empty", "exists", "state", "auth", "call", "response"
	Line int    // comment line number where the sequence begins (1-based; 0 = unknown)

	// get/post/put/delete/call common: function call
	Package string // "session" (package prefix; empty string if absent)
	Model   string // "Course.FindByID" or "auth.VerifyPassword"
	Args    []Arg  // call arguments

	// get/post/call: assignment
	Result *Result // result binding (nil means no assignment)

	// empty/exists: guard
	Target string // "course" or "course.InstructorID"

	// state: state transition
	DiagramID  string            // "reservation"
	Inputs     map[string]string // {status: "reservation.Status"}
	Transition string            // "cancel"

	// publish: event publishing
	Topic   string            // "order.completed"
	Options map[string]string // {delay: "1800"} (optional)
	// Inputs reused: payload

	// auth: authorization check
	Action   string // "delete"
	Resource string // "project"
	// Inputs reused     // {id: "project.ID", owner: "project.OwnerID"}

	// response: field mapping
	Fields map[string]string // {course: "course", instructor_name: "instructor.Name"}

	// verify-password: login timing-safe bundle (Phase010)
	//   @verify-password <Model>.<EmailCol>=<EmailExpr> <Model>.<HashCol> vs <PasswordExpr>
	//     -> <Result.Var> <ErrStatus> "<Message>"
	// Model/EmailCol/HashCol are the sqlc row type and column names; EmailExpr/PasswordExpr are
	// Go expressions resolved inside the handler. Result is the binding variable on success (reuses existing field).
	EmailCol     string
	EmailExpr    string
	HashCol      string
	PasswordExpr string

	// common
	Message      string // error message
	ErrStatus    int    // error HTTP status code (0 = type default: @call→500, @empty→404, @exists→409, @state→409, @auth→403)
	SuppressWarn bool   // @type! — suppress WARNING
}
