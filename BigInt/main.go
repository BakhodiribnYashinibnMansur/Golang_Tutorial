package main

import (
	bigInt "bigint/bigInt"
	"fmt"
)

func main() {
	a, err := bigInt.NewBigInt("777777777777777777777777")
	if err != nil {
		panic(err)
	}
	b, err := bigInt.NewBigInt("44444444444444444444")
	if err != nil {
		panic(err)
	}

	err = b.Set("1")
	if err != nil {
		panic(err)
	}
	c := bigInt.Add(a, b)

	d := bigInt.Sub(a, b)

	e := bigInt.Multiply(a, b)

	fmt.Println(a, b, c, d, e)
}
