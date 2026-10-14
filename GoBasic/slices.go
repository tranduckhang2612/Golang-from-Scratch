package gobasic

import (
	"fmt"
)
func BasicSlices(){
	a := [10]int{1,2,3,4,5,6,7}
	slice1 := a[3:5] // output: [4,5]
	slice1 = append(slice1, 8,12)
	fmt.Println(slice1)
	slice2:= make([]any, 5)
	slice2= append(slice2, 1, "hi", true)
	slice3:= make([]int, 6)
	slice3 = append(slice3, 1,23,3)
	fmt.Println(len(slice3), cap(slice3))

}

