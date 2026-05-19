//ff:type feature=features type=model
//ff:what FeaturesFile — features.yaml 최상위 구조 (features 리스트)
package features

// FeaturesFile is the top-level structure of features.yaml.
type FeaturesFile struct {
	Features []Feature `yaml:"features"`
}
