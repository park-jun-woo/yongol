//ff:type feature=projectconfig type=model
//ff:what Frontend — frontend 섹션 모델 (lang / framework / bundler / theme)

package manifest

// Frontend holds the React/TSX scaffold configuration. Theme maps directly
// to the tailwind.config.js `theme.extend.colors` block that yongol emits;
// absent fields fall back to shadcn defaults.
type Frontend struct {
	Lang      string         `yaml:"lang"`
	Framework string         `yaml:"framework"`
	Bundler   string         `yaml:"bundler"`
	Name      string         `yaml:"name"`
	Theme     *FrontendTheme `yaml:"theme,omitempty"`
}
