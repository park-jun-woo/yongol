//ff:type feature=validate-contract type=model
//ff:what missingSymbols — PRV-02 카테고리별(Queries/Calls/Fields) 누락 심볼 번들

package contract

// missingSymbols bundles the per-category lists of drifted references a
// preserved file still carries after the SSOT has been reshaped.
type missingSymbols struct {
	Queries []string
	Calls   []string
	Fields  []string
}
