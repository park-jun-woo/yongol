//ff:type feature=statemachine type=model topic=states
//ff:what Mermaid stateDiagram 파싱 결과를 담는 구조체
package statemachine

// StateDiagram represents a parsed Mermaid stateDiagram.
//
// ID keeps the lowercase filename stem because SSaC `@state <id>` and
// `states/<id>.md` match on that. Symbol is the PascalCase form and is
// the only field consumers should use when emitting exported Go
// identifiers (`<Symbol>Transitions`, `<Symbol>CanTransition`,
// `<Symbol>NextState`) — splicing ID directly into generated code
// produces unexported identifiers that break cross-package references.
type StateDiagram struct {
	ID           string       // filename stem (e.g. "course") — SSaC @state <id> 매칭용
	Symbol       string       // PascalCase of ID (e.g. "Course") — exported Go identifier 용
	File         string       // source file path (e.g. "specs/states/course.md")
	InitialState string       // state after [*] -->
	States       []string     // all unique state names
	Transitions  []Transition // all state transitions
}
