//ff:func feature=gen-ir type=util control=iteration dimension=1
//ff:what convertInputsToFieldArgs -- SSaC map[string]string Inputs → []FieldArg 변환

package ir

import "sort"

// convertInputsToFieldArgs converts an SSaC Inputs map (key → "source.Field"
// or literal) into a sorted slice of FieldArg. Keys are sorted for
// deterministic output.
func convertInputsToFieldArgs(inputs map[string]string) []FieldArg {
	if len(inputs) == 0 {
		return nil
	}
	keys := make([]string, 0, len(inputs))
	for k := range inputs {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	args := make([]FieldArg, 0, len(keys))
	for _, k := range keys {
		v := inputs[k]
		args = append(args, parseInputValue(k, v))
	}
	return args
}
