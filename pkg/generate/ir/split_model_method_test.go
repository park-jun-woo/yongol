//ff:func feature=gen-ir type=test control=sequence
//ff:what isCRUDSeq/parseDigits/splitModelMethod/planNeedsTransaction/csrfIsActive 순수 헬퍼
package ir

import (
	"testing"
)

func TestSplitModelMethod(t *testing.T) {
	model, method := splitModelMethod("Course.FindByID")
	if model != "Course" || method != "FindByID" {
		t.Errorf("got (%q,%q)", model, method)
	}
	model, method = splitModelMethod("NoDot")
	if model != "" || method != "NoDot" {
		t.Errorf("no dot got (%q,%q)", model, method)
	}
}
