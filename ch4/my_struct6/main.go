package main

import "fmt"

type Rect struct {
	x, y int
	w, h float64
}

func main()  {
	r := Rect{x:1, y:2, w:1.1, h:2.2}
	fmt.Printf("area = %g\n", r.Area())
}

func (r *Rect) Area() float64 {
	return  r.w * r.h
}

/**
如果函数名加一对括号，出现在要参数的后面，说明这个方法绑定到这个类型，这个类型有这个方法可以使用
 */
