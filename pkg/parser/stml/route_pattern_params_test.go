//ff:func feature=stml-parse type=test control=sequence
//ff:what RoutePatternParams — 세그먼트 이름 추출·optional ? strip·빈/정적 패턴 nil 검증

package stml

import (
	"reflect"
	"testing"
)

func TestRoutePatternParams(t *testing.T) {
	t.Run("required and optional segments, ? stripped", func(t *testing.T) {
		want := []string{"BuildingID", "UnitID", "PhotoID"}
		got := RoutePatternParams("/unit-info/:BuildingID/:UnitID/:PhotoID?")
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("static pattern returns nil", func(t *testing.T) {
		if got := RoutePatternParams("/account/settings"); got != nil {
			t.Errorf("got %v, want nil", got)
		}
	})

	t.Run("empty pattern returns nil", func(t *testing.T) {
		if got := RoutePatternParams(""); got != nil {
			t.Errorf("got %v, want nil", got)
		}
	})

	t.Run("bare colon segment is skipped", func(t *testing.T) {
		if got := RoutePatternParams("/x/:/y"); got != nil {
			t.Errorf("got %v, want nil", got)
		}
	})
}
