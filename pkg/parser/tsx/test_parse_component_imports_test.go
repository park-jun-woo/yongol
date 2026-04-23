//ff:func feature=tsx-parser type=test control=iteration dimension=1
//ff:what Parse — component_imports.tsx 의 로컬 component 만 수집, npm 패키지는 필터링

package tsx

import "testing"

func TestParse_ComponentImports(t *testing.T) {
	if !swcAvailable(t) {
		return
	}
	got, err := Parse("testdata/component_imports.tsx")
	if err != nil {
		t.Fatalf("Parse: %v", err)
	}
	want := map[string]string{
		"Button":       "@/components/ui/Button",
		"Card":         "@/components/ui/Card",
		"CreditsGauge": "./components/CreditsGauge",
	}
	gotImps := map[string]string{}
	for _, i := range got.Imports {
		gotImps[i.Name] = i.Path
	}
	for name, path := range want {
		if gotImps[name] != path {
			t.Errorf("import %s: want %q, got %q", name, path, gotImps[name])
		}
	}
	// react / clsx must NOT be in Imports (npm packages).
	for _, i := range got.Imports {
		if i.Path == "react" || i.Path == "clsx" {
			t.Errorf("npm package %q should be filtered, got in Imports", i.Path)
		}
	}
}
