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
	var err error
	val.IncludeNumbers = false
	val.IncludeSymbols = false
	val.IncludeUpper = false

	
	if len(os.Args) < 2 {
		fmt.Println("The program requires at least 1 number for the desired length of password")
		return
	} else if len(os.Args) > 5 {
		fmt.Println("The program can only take 4 arguments")
		return
	} else if len(os.Args) == 2 { // handle cases where only the desired length is specified
		val.length, err = strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println("Error while converting to int! Only digits are accepted as the first argument")
			return
		}
	} else if len(os.Args) == 3 { // handle cases where only the desired length and and symbol fields are specified
		val.length, err = strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println("Error while converting to int! Only digits are accepted as the first argument")
			return
		}
		if os.Args[2] == "Y" {
			val.IncludeSymbols = true
		} else if os.Args[2] == "N" {
			val.IncludeSymbols = false
		} else {
			fmt.Println("only Y or N are accepted inputs!")
			return
		}
	} else if len(os.Args) == 4 { // handle cases where the length, symbol, and number fields are specified
		val.length, err = strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println("Error while converting to int! Only digits are accepted as the first argument")
			return
		}
		if os.Args[2] == "Y" {
			val.IncludeSymbols = true
		} else if os.Args[2] == "N" {
			val.IncludeSymbols = false
		} else {
			fmt.Println("only Y or N are accepted inputs!")
			return
		}

		if os.Args[3] == "Y" {
			val.IncludeNumbers = true
		} else if os.Args[3] == "N" {
			val.IncludeNumbers = false
		} else {
			fmt.Println("only Y or N are accepted inputs!")
			return
		}
	} else if len(os.Args) == 5 { // handle cases where all the fields are specified
		val.length, err = strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println("Error while converting to int! Only digits are accepted as the first argument")
			return
		}
		if os.Args[2] == "Y" {
			val.IncludeSymbols = true
		} else if os.Args[2] == "N" {
			val.IncludeSymbols = false
		} else {
			fmt.Println("only Y or N are accepted inputs!")
			return
		}

		if os.Args[3] == "Y" {
			val.IncludeNumbers = true
		} else if os.Args[3] == "N" {
			val.IncludeNumbers = false
		} else {
			fmt.Println("only Y or N are accepted inputs!")
			return
		}

		if os.Args[4] == "Y" {
			val.IncludeUpper = true
		} else if os.Args[4] == "N" {
			val.IncludeUpper = false
		} else {
			fmt.Println("only Y or N are accepted inputs!")
			return
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

	rand.Seed(time.Now().UnixNano())
	output := make([]byte, val.length)

	for i := range output {
		output[i] = defaultSet[rand.Intn(len(defaultSet))]
	}
	return string(output)
}
