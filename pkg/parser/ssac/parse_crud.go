//ff:func feature=ssac-parse type=parser control=sequence
//ff:what @get/@post/@put/@delete CRUD 시퀀스 파싱
package ssac

import "strings"

// parseCRUD는 @get/@post/@put/@delete를 파싱한다.
// hasResult=true: Type var = Model.Method({Key: val, ...})
// hasResult=false: Model.Method({Key: val, ...})
func parseCRUD(seqType, rest string, hasResult bool) (*Sequence, error) {
	rest = strings.TrimSpace(rest)
	seq := &Sequence{Type: seqType}

	if hasResult {
		if err := parseCRUDWithResult(rest, seq); err != nil {
			return nil, err
		}
		if seq.Result == nil {
			return nil, nil
		}
	} else {
		model, inputs, _, err := parseCallExprInputs(rest)
		if err != nil {
			return nil, err
		}
		seq.Package, seq.Model = splitPackagePrefix(model)
		seq.Inputs = inputs
	}

	return seq, nil
}
