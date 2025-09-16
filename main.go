package main

import (
	"fmt"
	"github.com/fatih/color"
)

func main() {
	var firstNum, secondNum int
	var operators string

	for {
		fmt.Printf("Select operator: '%s', '%s', '%s', '%s', '%s'\n", "+", "-", "/", "*", "quit")
		_, err := fmt.Scan(&operators)
		if err != nil {
			color.Red("error:", err)
			return
		}

		if operators == "quit" {
			color.Blue("Goodbye!")
			break
		}

		fmt.Println("Enter first number & second number:")
		_, err = fmt.Scan(&firstNum, &secondNum)
		if err != nil {
			color.Red("error:", err)
			return
		}

		switch operators {
		case "+":
			fmt.Printf("%d + %d = %d\n", firstNum, secondNum, firstNum+secondNum)
		case "-":
			fmt.Printf("%d - %d = %d\n", firstNum, secondNum, firstNum-secondNum)
		case "/":
			if secondNum != 0 {
				result := float64(firstNum) / float64(secondNum)
				fmt.Printf("%d / %d = %f\n", firstNum, secondNum, result)
			} else {
				color.Red("error: division by zero")
			}
		case "*":
			fmt.Printf("%d * %d = %d\n", firstNum, secondNum, firstNum*secondNum)
		default:
			color.Red("unknown operator")
		}
	}
}
