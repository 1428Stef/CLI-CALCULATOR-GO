package main

import (
	"fmt"
<<<<<<< HEAD

	"github.com/fatih/color"
=======
>>>>>>> f047585c7987357528933120d76c179fc5e3a028
)

func main() {
	var firstNum, secondNum int
<<<<<<< HEAD
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

		fmt.Println("Enter fist number & second number:")
		_, err = fmt.Scan(&firstNum, &secondNum)
		if err != nil {
			color.Red("error:", err)
			return
		}

		switch operators {
=======
	var operation string
	for {

		fmt.Println("To exit, enter 'q'")

		if operation == "q" {
			break
		}

		fmt.Print("Enter two values: ")
		fmt.Scan(&firstNum, &secondNum)

		fmt.Println("Select operation: '+' '-' '/' '*' 'q'")
		fmt.Scan(&operation)

		switch operation {
>>>>>>> f047585c7987357528933120d76c179fc5e3a028
		case "+":
			fmt.Printf("%d + %d = %d\n", firstNum, secondNum, firstNum+secondNum)
		case "-":
			fmt.Printf("%d - %d = %d\n", firstNum, secondNum, firstNum-secondNum)
		case "/":
			if secondNum != 0 {
<<<<<<< HEAD
				result := float64(firstNum) / float64(secondNum)
				fmt.Printf("%d / %d = %f\n", firstNum, secondNum, result)
			} else {
				color.Red("error")
=======
				fmt.Printf("%d / %d = %d\n", firstNum, secondNum, firstNum/secondNum)
			} else {
				fmt.Println("error")
>>>>>>> f047585c7987357528933120d76c179fc5e3a028
			}
		case "*":
			fmt.Printf("%d * %d = %d\n", firstNum, secondNum, firstNum*secondNum)
		default:
<<<<<<< HEAD
			color.Red("unknown operator")
=======
			fmt.Println("unknown operation:", operation)
>>>>>>> f047585c7987357528933120d76c179fc5e3a028
		}
	}
}
