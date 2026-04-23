//ff:type feature=gen-react type=model
//ff:what manifestTheme — manifest.FrontendTheme 의 패키지-로컬 alias

package react

import "github.com/park-jun-woo/yongol/pkg/parser/manifest"

// manifestTheme is a package-local alias so the rest of the react package
// doesn't need a direct manifest import.
type manifestTheme = manifest.FrontendTheme
