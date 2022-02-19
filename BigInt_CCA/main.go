package main

import (
	"bigint-cca/bigInt"
	"fmt"
)

func main() {
	a, err := bigInt.NewInt("+7878946578574798745647")
	if err != nil {
		panic(err)
	}

	b, err := bigInt.NewInt("+4576545775645675675457")
	if err != nil {
		panic(err)
	}

	err = b.SetInt("+1")
	if err != nil {
		panic(err)
	}

	c := bigInt.Add(a, b)
	d := bigInt.Sub(a, b)
	e := bigInt.Multiply(a, b)
	f := bigInt.Mod(a, b)
	fmt.Println(a, b, c, d, e, f)

}
