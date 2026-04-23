//ff:type feature=projectconfig type=model
//ff:what FrontendTheme — shadcn semantic token 기반 tailwind theme.extend.colors 모델

package manifest

// FrontendTheme mirrors shadcn/ui's canonical semantic tokens. All fields
// are optional; unset tokens inherit shadcn defaults in the generated
// tailwind.config.js. Colors accept any CSS color string (hex, rgb, hsl).
type FrontendTheme struct {
	Primary     string `yaml:"primary,omitempty"`
	Secondary   string `yaml:"secondary,omitempty"`
	Accent      string `yaml:"accent,omitempty"`
	Destructive string `yaml:"destructive,omitempty"`
	Muted       string `yaml:"muted,omitempty"`
	Background  string `yaml:"background,omitempty"`
	Foreground  string `yaml:"foreground,omitempty"`
	Border      string `yaml:"border,omitempty"`
	Radius      string `yaml:"radius,omitempty"` // e.g. "0.5rem"
}
