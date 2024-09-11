package basics
import "fmt"

func FindOddOrEven(){
    var number int
    fmt.Print("Please Enter a number: ")
    fmt.Scanln(&number)

    if number % 2 == 0 {
        fmt.Printf("%d is Even number", number)
    }else{
        fmt.Printf("%d is Odd number", number)
    }
}