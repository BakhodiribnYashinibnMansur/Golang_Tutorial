package bigInt

import (
	"errors"
	"strconv"
	"unicode"
)

type BigInt struct {
	value int
}

func NewBigInt(num string) (BigInt, error) {
	// 1- method

	for _, symbol := range num {
		if unicode.IsLetter(symbol) {
			return BigInt{}, errors.New("INVALID NUMBER TO CREATE")
		}
	}

	// 2 - method

	// for _, symbol := range num {
	// 	if (symbol < 'a' || symbol > 'z') && (symbol < 'A' || symbol > 'Z') {
	// 		return BigInt{}, nil
	// 	}
	// }

	bigInt, err := strconv.Atoi(num)
	if err != nil {
		panic(err)
	}
	return BigInt{value: bigInt}, nil
}

func (bi *BigInt) Set(num string) error {
	for _, symbol := range num {
		if unicode.IsLetter(symbol) {
			return errors.New("INVALID NUMBER TO UPDATE")
		}
	}
	newBigInt, err := strconv.Atoi(num)
	if err != nil {
		panic(err)
	}
	bi.value = newBigInt
	return nil
}

func Add(a, b BigInt) BigInt {
	newBigInt := BigInt{}
	newBigInt.value = a.value + b.value
	return newBigInt
}

func Sub(a, b BigInt) BigInt {
	newBigInt := BigInt{}
	newBigInt.value = a.value - b.value
	return newBigInt
}

func Multiply(a, b BigInt) BigInt {
	newBigInt := BigInt{}
	newBigInt.value = a.value * b.value
	return newBigInt
}
func Mod(a, b BigInt) BigInt {
	newBigInt := BigInt{}
	newBigInt.value = a.value % b.value
	return newBigInt
}
