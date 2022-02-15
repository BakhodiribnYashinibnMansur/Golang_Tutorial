package main

import (
	"fmt"
	bigint "tutorial/bigInt"
)

func main() {
	a, err := bigint.NewBigInt("777777777777777777777777")
	if err != nil {
		panic(err)
	}
	b, err := bigint.NewBigInt("44444444444444444444")
	if err != nil {
		panic(err)
	}

	err = b.Set("1")
	if err != nil {
		panic(err)
	}
	c := bigint.Add(a, b)

	d := bigint.Sub(a, b)

	e := bigint.Multiply(a, b)

	fmt.Println(a, b, c, d, e)
}
