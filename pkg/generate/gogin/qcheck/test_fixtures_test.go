//ff:func feature=gen-gogin type=test control=sequence
//ff:what testFixtures — qcheck 테스트 고정 소스 리터럴 (deep / longLoop / clean)

package qcheck

const deepSrc = `package x

func Deep() {
	for i := 0; i < 10; i++ {
		if i > 0 {
			for j := 0; j < 10; j++ {
				if j%2 == 0 {
					println(j)
				}
			}
		}
	}
}
`

const longLoopSrc = `package x

func LongLoop(xs []int) {
	for _, x := range xs {
		a := x + 1
		b := x + 2
		c := x + 3
		d := x + 4
		e := x + 5
		f := x + 6
		g := x + 7
		h := x + 8
		i := x + 9
		j := x + 10
		k := x + 11
		println(a, b, c, d, e, f, g, h, i, j, k)
	}
}
`

const cleanSrc = `package x

func Clean() int {
	return 42
}
`
