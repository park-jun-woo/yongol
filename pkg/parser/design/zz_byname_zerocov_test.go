//ff:func feature=design-parse type=test control=sequence
//ff:what TestByName_ZeroCov — design 파서 헬퍼들을 이름으로 직접 호출해 커버리지 귀속

package design

import "testing"

func TestByNameDesignHelpers_ZeroCov(t *testing.T) {
	data := []byte("---\nversion: \"1.0\"\nname: theme\n---\n## Colors\n\n## Typography\n")
	yamlPart, body, err := parseFrontMatter(data)
	if err != nil {
		t.Fatalf("parseFrontMatter: %v", err)
	}
	if len(yamlPart) == 0 {
		t.Errorf("parseFrontMatter yamlPart empty")
	}

	headings := parseHeadings(body)
	if len(headings) != 2 {
		t.Errorf("parseHeadings = %v, want 2", headings)
	}

	// no front matter path.
	if _, _, err := parseFrontMatter([]byte("no front matter\n")); err == nil {
		t.Errorf("parseFrontMatter without delimiter should error")
	}
}
