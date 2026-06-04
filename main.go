package main

import (
	"errors"
	"fmt"
	"math"
)

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

	var dec int

	for i, number := range input {
		if number != '0' && number != '1' {
			err = errors.New("only 0s and 1s are allowed")
			fmt.Println(err)
			return

		}
		var byteChar = byte(number) - '0'

		exponent := len(input) - 1 - i
		dec = dec + (int(byteChar) * int(math.Pow(2, float64(exponent))))
	}

	fmt.Println(dec)
}
