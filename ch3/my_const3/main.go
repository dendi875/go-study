package main

type Weekday int

func main()  {
	const (
		One Weekday = iota
		Two
		Three
		Four
		Five
		Six
		Seven
	)

	println(One, Two, Three, Four, Five, Six, Seven)
}
