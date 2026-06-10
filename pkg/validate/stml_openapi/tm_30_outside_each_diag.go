//ff:func feature=validate type=rule control=sequence topic=stml-openapi
//ff:what TM-30 보조 — data-each 밖 item.* 소스 사용에 대한 ERROR 진단 생성

package stml_openapi

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// tm30OutsideEachDiag builds the TM-30 diagnostic for an item.* param
// source used outside any data-each block — no row is in scope there.
func tm30OutsideEachDiag(file, opID, source string) diagnostic.Diagnostic {
	return diagnostic.Diagnostic{
		File:        file,
		Phase:       diagnostic.PhaseValidate,
		Level:       diagnostic.LevelError,
		Message:     fmt.Sprintf("[TM-30] data-param source %q is only valid inside a data-each block — no row item is in scope here", source),
		Advice:      "Move the action inside the data-each block whose rows it targets, or use a route.<Name> source",
		OperationID: opID,
	}
}
