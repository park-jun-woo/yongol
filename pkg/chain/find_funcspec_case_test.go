//ff:type feature=chain type=model
//ff:what findFuncSpecCase — TestFindFuncSpec 의 단일 테이블 케이스
package chain

// findFuncSpecCase is a single table case for TestFindFuncSpec.
type findFuncSpecCase struct {
	name      string
	pkg       string
	funcName  string
	wantOK    bool
	wantFile  string
	wantSumm  string
	wantLineP bool // whether to assert a positive line number
}
