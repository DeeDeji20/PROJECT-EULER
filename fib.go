package main

import(
	"fmt"
)
var xi []int
var ii []int
func main() {
	a:= foo()
	fmt.Println(a)
}
func foo() []int {
	x, y := 1, 1
	xi = append(xi, x, y)
	for i := 0; i <31; i++{
		x, y = x+y, x
		xi = append(xi, x)
	}
for _, x = range xi{
	if x%2 == 0{
		ii = append(ii, x)
	}
}
return ii
}

// Declare a slice of int xi
// Create a function foo which returns a slice of ints
// Declare variables x, y and assign the value one to each
// Append x and y to the xi as part of the fibonacci sequence
// Loop over the index position which gives the fibonacci number less than 4000000
// The logic x+y assigned to x and x assigned to y gives the order of the sequence
// Append the vales of x to the slice of string xi
// Create another slice of int (ii) outside func main
// Range over the slice xi, using conditional if, 
// append the values of multiples of 2 to the slice of int ii
// return the slice int ii, which gives the even numbers in the fibonacii sequence