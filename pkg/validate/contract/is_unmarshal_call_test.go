//ff:func feature=validate-contract type=test control=selection topic=preserve-safety
//ff:what TestIsUnmarshalCall — call 이 json/yaml/toml/xml Unmarshal 호출인지 판정 검증

package contract

import "testing"

func TestIsUnmarshalCall(t *testing.T) {
	tests := []struct {
		name string
		src  string
		want bool
	}{
		{"json", "json.Unmarshal(b, &v)", true},
		{"yaml", "yaml.Unmarshal(b, &v)", true},
		{"toml", "toml.Unmarshal(b, &v)", true},
		{"xml", "xml.Unmarshal(b, &v)", true},
		{"unknown pkg", "proto.Unmarshal(b, &v)", false},
		{"non unmarshal", "json.Marshal(v)", false},
		{"free function", "Unmarshal(b, &v)", false},
		{"chained receiver", "enc.json.Unmarshal(b)", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isUnmarshalCall(mustCall(t, tt.src)); got != tt.want {
				t.Fatalf("isUnmarshalCall(%q) = %v, want %v", tt.src, got, tt.want)
			}
		})
	}
	if isUnmarshalCall(nil) {
		t.Fatal("nil call should return false")
	}
}
