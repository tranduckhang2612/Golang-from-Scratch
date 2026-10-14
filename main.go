package main
import ("fmt"; 
	"golang-from-scratch/GoBasic")

func main()  {
	fmt.Println("Hello")
	// gobasic.BasicArray(
	// )
	gobasic.BasicSlices()


	// Anonymous Function
	func ()  {
		fmt.Println("Hello Anonymous Function")
	} ()

	// functions.go
	result:= gobasic.ApplyOperation(5, 3, gobasic.Add)
	fmt.Println("5 + 3 = ", result)

	multiplyBy2 := gobasic.CreateMultiplier(2)
	fmt.Println("6 * 2 = ", multiplyBy2(6))
}

