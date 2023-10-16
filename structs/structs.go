package main

import (
	"fmt"
)

type Measure struct {
	height float32
	weight int
}
type employee struct {
	firstName, lastName string
	age                 int
	measure             Measure
}

func main() {
	p := &employee{
		firstName: "Katie",
		lastName:  "Morton",
		age:       20,
		measure: Measure{
			height: 167.64,
			weight: 66,
		},
	}
	fmt.Println("first name : ", (*p).firstName)
	fmt.Println("age : ", (*p).age)
	fmt.Println("Height: ", (*p).measure.height)
}
