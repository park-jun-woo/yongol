//ff:func feature=validate type=rule control=sequence topic=stml-openapi
//ff:what tm53Diag — TM-53 WARNING 진단 한 건을 구성한다

package stml_openapi

import "github.com/park-jun-woo/yongol/pkg/diagnostic"

// tm53Diag builds a single TM-53 WARNING diagnostic.
func tm53Diag(file, message, advice string) []diagnostic.Diagnostic {
	return []diagnostic.Diagnostic{{
		File:    file,
		Phase:   diagnostic.PhaseValidate,
		Level:   diagnostic.LevelWarning,
		Message: message,
		Advice:  advice,
	}}
}
