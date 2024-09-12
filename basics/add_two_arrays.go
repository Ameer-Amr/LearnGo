package basics
import "fmt"

func AddTwoArrays() {
	var size int
	fmt.Println("Enter the array size")
	fmt.Scan(&size)
	array1 := make([]int, size)
	array2 := make([]int, size)
	sum_array := make([]int, size)

	fmt.Printf("Enter %d elements to Array1: \n", size)
	for i:=0; i<size; i++ {
		fmt.Scan(&array1[i])
	}
	fmt.Printf("Array 1: %d \n", array1)


	fmt.Printf("Enter %d elements to Array2: \n", size)
	for i:=0; i<size; i++ {
		fmt.Scan(&array2[i])
	}
	fmt.Printf("Array 2: %d \n", array2)

	for i:=0; i<size; i++ {
		sum_array[i] = array1[i] + array2[i] 
	}
	fmt.Printf("Sum of Two Array : %d \n", sum_array)
}