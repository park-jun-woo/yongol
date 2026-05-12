//ff:type feature=gen-react type=generator
package react

// semanticColorNames lists the DESIGN.md color keys that are already handled
// by the fixed semantic slots (primary, secondary, accent, destructive,
// muted, background, foreground, border) and their -foreground variants.
var semanticColorNames = map[string]bool{
	"primary":                true,
	"primary-foreground":     true,
	"secondary":              true,
	"secondary-foreground":   true,
	"accent":                 true,
	"accent-foreground":      true,
	"destructive":            true,
	"destructive-foreground": true,
	"muted":                  true,
	"muted-foreground":       true,
	"background":             true,
	"foreground":             true,
	"border":                 true,
}
