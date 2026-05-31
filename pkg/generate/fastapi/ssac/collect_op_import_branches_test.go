//ff:func feature=gen-fastapi type=test control=iteration dimension=1
//ff:what TestCollectOpImportBranches — 서브테스트 디스패치
package ssac

import "testing"

func TestCollectOpImportBranches(t *testing.T) {
	for _, st := range []struct {
		name string
		fn   func(*testing.T)
	}{
		{"Get", subtestTestCollectOpImportBranchesGet},
		{"GetNil", subtestTestCollectOpImportBranchesGetNil},
		{"Post", subtestTestCollectOpImportBranchesPost},
		{"PostNil", subtestTestCollectOpImportBranchesPostNil},
		{"Put", subtestTestCollectOpImportBranchesPut},
		{"PutNil", subtestTestCollectOpImportBranchesPutNil},
		{"Delete", subtestTestCollectOpImportBranchesDelete},
		{"DeleteNil", subtestTestCollectOpImportBranchesDeleteNil},
		{"Publish", subtestTestCollectOpImportBranchesPublish},
		{"VerifyPassword", subtestTestCollectOpImportBranchesVerifyPassword},
		{"VerifyPasswordNil", subtestTestCollectOpImportBranchesVerifyPasswordNil},
		{"CallExternal", subtestTestCollectOpImportBranchesCallExternal},
		{"CallSelfSkipped", subtestTestCollectOpImportBranchesCallSelfSkipped},
		{"CallNilOrEmptyPackage", subtestTestCollectOpImportBranchesCallNilOrEmptyPackage},
		{"EvalExternal", subtestTestCollectOpImportBranchesEvalExternal},
		{"EvalSelfSkipped", subtestTestCollectOpImportBranchesEvalSelfSkipped},
		{"EvalNilOrEmpty", subtestTestCollectOpImportBranchesEvalNilOrEmpty},
	} {
		t.Run(st.name, st.fn)
	}
}
