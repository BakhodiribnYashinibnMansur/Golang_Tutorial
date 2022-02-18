package bigInt

import (
	"errors"
	"regexp"
	"strconv"
)

type Bigint struct {
	Value string
}

var ErrorBadInput = errors.New("bad input , Please only number ")

func checkInput(num string) (bool, error) {
	match, err := regexp.MatchString(`^[+-]?[0-9]*$`, num)
	if err != nil {

		return false, err
	} else {
		return match, nil
	}
}

func NewInt(num string) (Bigint, error) {
	flag, err := checkInput(num)
	if err != nil {
		return Bigint{}, err
	}
	if !flag {
		return Bigint{}, ErrorBadInput
	}
	return Bigint{Value: num}, nil
}

func (bgt *Bigint) SetInt(num string) error {
	flag, err := checkInput(num)
	if err != nil {
		return err
	}
	if !flag {
		return ErrorBadInput
	}
	bgt.Value = num
	return nil
}

//Add two numbers
func Add(a, b Bigint) Bigint {
	int1, _ := strconv.Atoi(a.Value)
	int2, _ := strconv.Atoi(b.Value)
	add := int1 + int2
	strParse := strconv.Itoa(add)
	return Bigint{Value: strParse}
}

//Subtraction two numbers
func Sub(a, b Bigint) Bigint {
	int1, _ := strconv.Atoi(a.Value)
	int2, _ := strconv.Atoi(b.Value)
	sub := int1 - int2
	strParse := strconv.Itoa(sub)
	return Bigint{Value: strParse}
}

//Multiply two numbers
func Multiply(a, b Bigint) Bigint {
	int1, _ := strconv.Atoi(a.Value)
	int2, _ := strconv.Atoi(b.Value)
	mult := int1 * int2
	strParse := strconv.Itoa(mult)
	return Bigint{Value: strParse}
}

// Find mod two numbers divided
func Mod(a, b Bigint) Bigint {
	int1, _ := strconv.Atoi(a.Value)
	int2, _ := strconv.Atoi(b.Value)
	mod := int1 % int2
	strParse := strconv.Itoa(mod)
	return Bigint{Value: strParse}
}
