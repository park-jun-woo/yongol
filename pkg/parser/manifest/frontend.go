//ff:type feature=projectconfig type=model
//ff:what Frontend — frontend 섹션 모델 (lang / framework / bundler / theme / index)

package manifest

// Frontend holds the React/TSX scaffold configuration. Theme maps directly
// to the tailwind.config.js `theme.extend.colors` block that yongol emits;
// absent fields fall back to shadcn defaults. Index names the STML page
// (filename without .html — a page-name reference, not a path) the "/"
// index route redirects to (page-flow Phase009, BUG-114 (3)); TM-34
// validates the reference and TM-35 surfaces the file-name-sort fallback
// when neither frontend.index nor a data-route="/" page declares the index.
type Frontend struct {
	Enabled       *bool          `yaml:"enabled,omitempty"`
	Lang          string         `yaml:"lang"`
	Framework     string         `yaml:"framework"`
	Bundler       string         `yaml:"bundler"`
	Name          string         `yaml:"name"`
	Theme         *FrontendTheme `yaml:"theme,omitempty"`
	Auth          *FrontendAuth  `yaml:"auth,omitempty"`
	Design        string         `yaml:"design,omitempty"`
	DefaultLayout string         `yaml:"defaultLayout,omitempty"`
	Index         string         `yaml:"index,omitempty"`
}
