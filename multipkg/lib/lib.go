package lib

type abc struct {
	xx string
	yy def
}
type def struct {
	a int
	b bool
}

func NewHello() *abc {
	return &abc{
		xx: "new hello",
		yy: def{
			a: 12,
			b: true,
		},
	}
}
func newHello() *abc {
	return &abc{
		xx: "new hello",
	}
}
func NewHello1() abc {
	return abc{
		xx: "new hello",
	}
}
