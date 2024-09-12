package basics
import "fmt"

func Star(){
	rows:=5
    for i := 1; i <= rows; i++ {
        // Print leading spaces
        for j := i; j < rows; j++ {
            fmt.Print(" ")
        }
        // Print stars
        for j := 1; j <= (2*i - 1); j++ {
            fmt.Print("*")
        }
        fmt.Println()
    }
    for i := rows - 1; i >= 1; i-- {
        // Print leading spaces
        for j := rows; j > i; j-- {
            fmt.Print(" ")
        }
        // Print stars
        for j := 1; j <= (2*i - 1); j++ {
            fmt.Print("*")
        }
        fmt.Println()
    }
}