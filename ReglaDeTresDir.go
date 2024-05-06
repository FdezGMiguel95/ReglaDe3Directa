package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Println("necesito tres cifras")
		os.Exit(1)
	}

	a, errA := strconv.Atoi(os.Args[1])
	b, errB := strconv.Atoi(os.Args[2])
	c, errC := strconv.Atoi(os.Args[3])

	if errA != nil || errB != nil || errC != nil {
		fmt.Println("Los argumentos deben ser enteros")
	}

	x := DoTheMath(a, b, c)

	fmt.Println(strconv.Itoa(a) + " -> " + strconv.Itoa(b))
	fmt.Println(strconv.Itoa(c) + " -> " + strconv.Itoa(x))
}

func DoTheMath(a, b, c int) int {
	x := (c * b) / a
	return x
}
