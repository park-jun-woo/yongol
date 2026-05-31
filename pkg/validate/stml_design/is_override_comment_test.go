//ff:func feature=validate type=test control=sequence topic=stml-design
//ff:what TestStmlDesignHelpers — unit tests for the pure stml_design helper functions
package stml_design

import (
	"testing"
)

func TestIsOverrideComment(t *testing.T) {
	if !isOverrideComment(" @override class=\"x\" ") {
		t.Error("expected override comment recognised")
	}
	if isOverrideComment(" normal comment") {
		t.Error("expected non-override")
	}
}
