module gopl.io/ch2/ex02_02

go 1.25.0

replace gopl.io/ch2/tempconv => ./tempconv

replace gopl.io/ch2/lengthconv => ./lengthconv

replace gopl.io/ch2/weightconv => ./weightconv

require (
	gopl.io/ch2/lengthconv v0.0.0-00010101000000-000000000000
	gopl.io/ch2/tempconv v0.0.0-00010101000000-000000000000
	gopl.io/ch2/weightconv v0.0.0-00010101000000-000000000000
)
