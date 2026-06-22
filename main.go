package main

import (
	"errors"
	"fmt"
	"math"
)

type Converter interface {
	Convert(input string) (int, error)
}

type BinaryConverter struct{}

func (b BinaryConverter) Convert(input string) (int, error) {
	var decimal int

	var err error
	for i, number := range input {
		if number != '0' && number != '1' {
			err = errors.New("only 0s and 1s are allowed")
			fmt.Println(err)
			return 0, err
		}
		var byteChar = byte(number) - '0'

		exponent := len(input) - 1 - i
		decimal = decimal + (int(byteChar) * int(math.Pow(2, float64(exponent))))

	}
	return decimal, err
}

func main() {
	fmt.Println("Enter the binary number (up to 8 digits): ")

	var input string
	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println("Error reading input")
		return
	}

	if len(input) > 8 {
		err = errors.New("input cannot be longer than 8 digits")
		fmt.Println(err)
		return
	}

	var binary Converter = BinaryConverter{}
	bin, err := binary.Convert(input)
	if err != nil {
		return
	}
	fmt.Println(bin)

}
