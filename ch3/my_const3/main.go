package main

type Num int

func main()  {
	const (
		Zero Num = iota
		One
		Two
		Three
		Four
		Five
		Six
		Seven
	)

	println(Zero, One, Two, Three, Four, Five, Six, Seven)
}
