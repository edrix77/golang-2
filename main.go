package main

import "fmt"

func swap(temp *int, xr *int, yr *int) {
	//var temp int

	*temp = *xr
	*xr = *yr
	*yr = *temp
}

func main() {
	x := 1
	y := 2
	temp := new(int)
	swap(temp, &x, &y)
	fmt.Println("x = ", x, "y = ", y, "temp = ", *temp)
}
