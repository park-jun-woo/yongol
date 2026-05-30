//ff:func feature=ssac-parse type=test control=iteration dimension=1
//ff:what parseArg / parseArgs / filterSubscribe / buildSubscribeInfo 단위 검증

package ssac

import "testing"

func TestParseArg(t *testing.T) {
	cases := []struct {
		name string
		in   string
		want Arg
	}{
		{"quoted literal", `"admin"`, Arg{Literal: "admin", IsQuoted: true}},
		{"numeric literal", "42", Arg{Literal: "42"}},
		{"float literal", "3.14", Arg{Literal: "3.14"}},
		{"source field", "request.CourseID", Arg{Source: "request", Field: "CourseID"}},
		{"bare variable", "course", Arg{Source: "course"}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := parseArg(c.in)
			if got != c.want {
				t.Errorf("parseArg(%q) = %+v, want %+v", c.in, got, c.want)
			}
		})
	}
}

func TestParseArgs(t *testing.T) {
	got := parseArgs(`request.ID, "admin" , 7`)
	if len(got) != 3 {
		t.Fatalf("len = %d, want 3 (%+v)", len(got), got)
	}
	if got[0].Source != "request" || got[0].Field != "ID" {
		t.Errorf("arg0 = %+v", got[0])
	}
	if got[1].Literal != "admin" || !got[1].IsQuoted {
		t.Errorf("arg1 = %+v", got[1])
	}
	if got[2].Literal != "7" {
		t.Errorf("arg2 = %+v", got[2])
	}
	// empty parts skipped
	if g := parseArgs("  ,  "); g != nil {
		t.Errorf("empty args = %+v, want nil", g)
	}
}

func TestBuildSubscribeInfo(t *testing.T) {
	t.Run("with param", func(t *testing.T) {
		si := buildSubscribeInfo("order.completed", &ParamInfo{TypeName: "OnOrderCompletedMessage"})
		if si.Topic != "order.completed" || si.MessageType != "OnOrderCompletedMessage" {
			t.Errorf("si = %+v", si)
		}
	})
	t.Run("nil param", func(t *testing.T) {
		si := buildSubscribeInfo("x.y", nil)
		if si.Topic != "x.y" || si.MessageType != "" {
			t.Errorf("si = %+v", si)
		}
	})
}

func TestFilterSubscribe(t *testing.T) {
	sf := &ServiceFunc{Param: &ParamInfo{TypeName: "Msg"}}
	seqs := []Sequence{
		{Type: "subscribe", Topic: "order.done"},
		{Type: "get", Model: "Course.FindByID"},
	}
	filtered := filterSubscribe(sf, seqs)
	if len(filtered) != 1 || filtered[0].Type != "get" {
		t.Errorf("filtered = %+v", filtered)
	}
	if sf.Subscribe == nil || sf.Subscribe.Topic != "order.done" || sf.Subscribe.MessageType != "Msg" {
		t.Errorf("subscribe = %+v", sf.Subscribe)
	}
}
