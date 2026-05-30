//ff:func feature=gen-splitter type=test control=iteration dimension=1
//ff:what snake / needsSnakeBreak / isVersionSuffix / tailSegment / summariseDoc 순수 헬퍼

package splitter

import "testing"

func TestSnake(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{"CamelCase", "camel_case"},
		{"HTTPServer", "http_server"},
		{"APIKey", "api_key"},
		{"simple", "simple"},
		{"ID", "id"},
		{"UserID", "user_id"},
		{"", ""},
	}
	for _, c := range cases {
		if got := snake(c.in); got != c.want {
			t.Errorf("snake(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}

func TestNeedsSnakeBreak(t *testing.T) {
	cases := []struct {
		s    string
		i    int
		want bool
	}{
		{"CamelCase", 0, false},  // i==0
		{"CamelCase", 5, true},   // l->C boundary (Camel|Case)
		{"HTTPServer", 4, true},  // acronym end: P->S where S precedes lowercase e
		{"HTTPServer", 1, false}, // T after H, both upper, next upper -> no break
		{"abc", 1, false},        // not upper
		{"x9Y", 2, true},         // digit before upper
	}
	for _, c := range cases {
		if got := needsSnakeBreak([]rune(c.s), c.i); got != c.want {
			t.Errorf("needsSnakeBreak(%q,%d) = %v, want %v", c.s, c.i, got, c.want)
		}
	}
}

func TestIsVersionSuffix(t *testing.T) {
	cases := []struct {
		in   string
		want bool
	}{
		{"v2", true},
		{"v10", true},
		{"v", false},
		{"x2", false},
		{"v2a", false},
		{"", false},
	}
	for _, c := range cases {
		if got := isVersionSuffix(c.in); got != c.want {
			t.Errorf("isVersionSuffix(%q) = %v, want %v", c.in, got, c.want)
		}
	}
}

func TestTailSegment(t *testing.T) {
	if got := tailSegment("/a/b/c.go"); got != "c.go" {
		t.Errorf("tailSegment = %q, want c.go", got)
	}
	if got := tailSegment("plain.go"); got != "plain.go" {
		t.Errorf("tailSegment = %q, want plain.go", got)
	}
}

func TestSummariseDoc(t *testing.T) {
	cases := []struct {
		name     string
		doc      string
		fallback string
		want     string
	}{
		{"first non-blank", "// first line\n// second", "fb", "first line"},
		{"skip leading blank", "\n\n// real", "fb", "real"},
		{"empty uses fallback", "", "MyFunc", "MyFunc"},
		{"empty no fallback", "   \n  ", "", "generated"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := summariseDoc(c.doc, c.fallback); got != c.want {
				t.Errorf("summariseDoc(%q,%q) = %q, want %q", c.doc, c.fallback, got, c.want)
			}
		})
	}
}
