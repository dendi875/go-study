package main

import (
	"bytes"
	"fmt"
)

func main()  {
	fmt.Println(intsToSting([]int{1, 2, 3}))
}

func intsToSting(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')

	for i, v := range values {
		if i > 0 {
			buf.WriteString(" ,")
		}
		fmt.Fprintf(&buf, "%d", v)
	}

	buf.WriteByte(']')
	return buf.String()
}

/**
1, 2, 3 => [1, 2, 3]
 */