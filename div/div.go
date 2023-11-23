package main

import (
	"fmt"
	"log"
)

func main() {
	// fmt.Println(div(1, 0))
	fmt.Println(safeDiv(1, 0))
	fmt.Println(safeDiv(7, 2))
}

// func div(a, b int) int {
// 	return a / b
// }

func safeDiv(a, b int) (q int, err error) {
	// q & err are local variables in safeDiv
	defer func() {
		// e's type is any (or interface{}) not error
		if e := recover(); e != nil {
			log.Println("ERROR:", e)
			err = fmt.Errorf("%v", e)
		}
	}()

	return a / b, nil
	// possible, but not recommended
	// q = a / b
	// return
}