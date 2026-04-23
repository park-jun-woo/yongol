//ff:type feature=tsx-parser type=model
//ff:what ComponentImport — 로컬 component import 1건 (name + path + line)

package tsx

// ComponentImport captures a local component import (non-npm).
// Only imports whose source starts with "@/components/" or "./components/"
// are recorded; npm package imports are filtered out before emission.
type ComponentImport struct {
	Name string // imported symbol ("Button")
	Path string // module source ("@/components/ui/Button")
	Line int
}
