//ff:func feature=gen-gogin type=test control=sequence
//ff:what lookupHTTPWhat 단위 테스트 (Summary > Description > 기본값 우선순위)
package ssac

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestLookupHTTPWhat(t *testing.T) {
	sf := ssacparser.ServiceFunc{Name: "GetThing"}

	t.Run("nil fullstack → fallback", func(t *testing.T) {
		if got := lookupHTTPWhat(nil, sf); got != "HTTP handler" {
			t.Errorf("got %q, want HTTP handler", got)
		}
	})
	t.Run("summary preferred", func(t *testing.T) {
		fs := docWithOp(&openapi3.Operation{
			OperationID: "GetThing",
			Summary:     "Fetch a thing\nsecond line",
			Description: "long desc",
		})
		if got := lookupHTTPWhat(fs, sf); got != "Fetch a thing" {
			t.Errorf("got %q, want first line of summary", got)
		}
	})
	t.Run("description fallback", func(t *testing.T) {
		fs := docWithOp(&openapi3.Operation{
			OperationID: "GetThing",
			Description: "Describe the thing",
		})
		if got := lookupHTTPWhat(fs, sf); got != "Describe the thing" {
			t.Errorf("got %q, want description", got)
		}
	})
	t.Run("no match → fallback", func(t *testing.T) {
		fs := docWithOp(&openapi3.Operation{OperationID: "Other"})
		if got := lookupHTTPWhat(fs, sf); got != "HTTP handler" {
			t.Errorf("got %q, want HTTP handler", got)
		}
	})
}
