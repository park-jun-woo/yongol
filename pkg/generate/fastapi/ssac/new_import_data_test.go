//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestCollectOpImportBranches — Get/Post/Put/Delete/Publish/VerifyPW/Call/Eval 분기 커버
package ssac

func newImportData() importData {
	return importData{
		Models:  make(map[string]bool),
		ExtPkgs: make(map[string]map[string]bool),
	}
}
