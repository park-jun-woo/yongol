//ff:func feature=rule type=parser control=sequence topic=catalog
//ff:what rulebookParseState.resetTable — 테이블 경계를 넘을 때 in/header 플래그를 초기화

package catalog

func (st *rulebookParseState) resetTable() {
	st.inTable = false
	st.sawHeader = false
}
