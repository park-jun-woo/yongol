//ff:type feature=ssac-parse type=model
//ff:what ServiceFunc — type representing a single service function declaration
package ssac

// ServiceFunc represents a single service function declaration.
type ServiceFunc struct {
	Name         string         // function name (e.g. "GetCourse")
	FileName     string         // source file name
	Line         int            // line number of the function declaration (1-based; 0 = unknown)
	Feature      string         // feature folder name (e.g. "auth"; empty string if absent)
	Sequences    []Sequence     // sequence list
	Imports      []string       // Go import paths
	Subscribe    *SubscribeInfo // nil means HTTP trigger
	Param        *ParamInfo     // function parameter (for subscribe functions)
	Structs      []StructInfo   // Go structs declared in the .ssac file
	NoPagination bool           // true when a // @no-pagination comment is present — exempts from S-63
	StateNeutral bool           // true when a // @state-neutral comment is present — exempts from XSM-27
}
