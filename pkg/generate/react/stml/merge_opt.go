//ff:func feature=stml-gen type=util control=sequence
//ff:what 두 GenerateOptions를 병합하여 override 값을 우선 적용한다
package stml

func mergeOpt(base, override GenerateOptions) GenerateOptions {
	if override.APIImportPath != "" {
		base.APIImportPath = override.APIImportPath
	}
	base.UseClient = override.UseClient
	base.BearerAuth = override.BearerAuth
	if override.RequestConstraints != nil {
		base.RequestConstraints = override.RequestConstraints
	}
	if override.ResponseArrayItemFields != nil {
		base.ResponseArrayItemFields = override.ResponseArrayItemFields
	}
	if override.NoBodyOps != nil {
		base.NoBodyOps = override.NoBodyOps
	}
	if override.PathParamTypes != nil {
		base.PathParamTypes = override.PathParamTypes
	}
	return base
}
