//ff:func feature=ssac-parse type=test control=sequence
//ff:what parseResult/parseGuard/parseCRUD(no-result) 파싱 검증
package ssac

import (
	"testing"
)

func TestParseCRUD_NoResult(t *testing.T) {
	seq, err := parseCRUD("delete", "Reservation.Delete({ID: request.ID})", false)
	if err != nil {
		t.Fatalf("err = %v", err)
	}
	if seq.Type != "delete" || seq.Model != "Reservation.Delete" {
		t.Errorf("seq = %+v", seq)
	}
	if seq.Inputs["ID"] != "request.ID" {
		t.Errorf("inputs = %v", seq.Inputs)
	}
}
