package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Print("write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("unkown")

	}

	switch time.Now().Weekday() {

	case time.Saturday, time.Sunday:
		fmt.Println("周末")
	default:
		fmt.Println("上班日")
	}
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("上午")
	default:
		fmt.Println("下午")
	}

	typeTest := func(i interface{}) {
		switch i.(type) {
		case bool:
			fmt.Println("Bool")
		case int:
			fmt.Println("int")
		default:
			fmt.Println("unkowen type")

		}
	}

	typeTest(true)
	typeTest(10)
	typeTest("int")
}
