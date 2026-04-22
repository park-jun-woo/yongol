//ff:type feature=orchestrator type=model
//ff:what Diagnostic message for the parse/validate/crosscheck phases
package diagnostic

// Diagnostic represents a single diagnostic message from any phase.
type Diagnostic struct {
	File    string // source file path
	Line    int    // line number (0 if unknown)
	Phase   Phase  // parse, validate
	Level   Level  // error, warning
	Message string // main message (Rule-ID + what is wrong)
	Advice  string // → Advice body (how to fix). Hidden when empty.
}
