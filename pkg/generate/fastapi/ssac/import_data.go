//ff:type feature=gen-fastapi type=model
//ff:what importData — feature service 파일 import 수집 결과

package ssac

// importData holds the collected import information for a feature service file.
type importData struct {
	UsesSelect bool
	UsesUpdate bool
	UsesDelete bool
	HasPublish bool
	HasAuth    bool
	Models     map[string]bool
	ExtPkgs    map[string]map[string]bool // pkg → set of functions
}
