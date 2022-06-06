//  9/80
package main

import "fmt"

func main() {
	s := make([]string, 4, 10)
	fmt.Println(s)
	s[0] = "1"
	s[1] = "2"
	s[2] = "3"
	s[3] = "4"
	fmt.Println("set:", s)
	fmt.Println("get:", s)
	fmt.Println(len(s))
	fmt.Println(cap(s))

	s = append(s, "d")
	s = append(s, "e", "f")

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	l := s[2:5]
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("2d:", twoD)

}
