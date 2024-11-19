package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type PasswordConfig struct {
	length         int
	IncludeNumbers bool
	IncludeSymbols bool
	IncludeUpper   bool
}

func main() {
	val := PasswordConfig{}
	val.length = 0
	val.IncludeNumbers = false
	val.IncludeSymbols = false
	val.IncludeUpper = false
	var err error

	if len(os.Args) < 2 {
		fmt.Println("The program requires at least 1 number for the desired length of password")
		return
	} else if len(os.Args) > 5 {
		fmt.Println("The program can only take 4 arguments")
		return
	}

	if len(os.Args) >= 2 { // handle cases where only the desired length is specified
		val.length, err = strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println("Error while converting to int! Only digits are accepted as the first argument")
			return
		}

		if len(os.Args) >= 3 { // handle cases where only the desired length and and symbol fields are specified
			if os.Args[2] == "Y" {
				val.IncludeSymbols = true
			} else if os.Args[2] == "N" {
				val.IncludeSymbols = false
			} else {
				fmt.Println("only Y and N are the accepted letters!")
				return
			}
		}

		if len(os.Args) >= 4 {
			if os.Args[3] == "Y" {
				val.IncludeNumbers = true
			} else if os.Args[3] == "N" {
				val.IncludeNumbers = false
			} else {
				fmt.Println("only Y and N are the accepted letters!")
				return
			}
		}

		if len(os.Args) == 5 {
			if os.Args[4] == "Y" {
				val.IncludeUpper = true
			} else if os.Args[4] == "N" {
				val.IncludeUpper = false
			} else {
				fmt.Println("only Y and N are the accepted letters!")
				return
			}
		}
	}

	password := val.GeneratePassword()
	fmt.Println(password)
}

// Method to generate a random password containing the specified fields
func (val PasswordConfig) GeneratePassword() string {
	defaultSet := "abcdefghijklmnopqrstuvwxyz"

	if val.IncludeSymbols {
		defaultSet += "!@#$%^&*()"
	} else {
		defaultSet += ""
	}

	if val.IncludeNumbers {
		defaultSet += "0123456789"
	} else {
		defaultSet += ""
	}
	if val.IncludeUpper {
		defaultSet += "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	} else {
		defaultSet += ""
	}

	//initialize generator with a unique seed
	rand.Seed(time.Now().UnixNano())
	output := make([]byte, val.length)

	for i := range output {
		output[i] = defaultSet[rand.Intn(len(defaultSet))]
	}
	return string(output)
}
