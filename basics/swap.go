package basics
import "fmt"

func SwapNumbers() {
	var num1 int
	var num2 int
	var temp int
	fmt.Println("Please enter first number: ")
	fmt.Scanln(&num1)
	fmt.Println("Please enter second number: ")
	fmt.Scanln(&num2)
	fmt.Printf("Entered numbers are %d and %d \n", num1, num2)
	temp = num1
	num1 = num2
	num2 = temp
	fmt.Printf("After swaping num1 is %d and num2 is %d", num1, num2)
}