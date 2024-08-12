package main

// func main() {
// 	string1 := "Hello, World"
// 	string2 := "hello, world"

// 	// case insensitive match
// 	fmt.Println(strings.EqualFold(string1, string2))

// 	// index
// 	wIndex := strings.Index(string1, "W")
// 	fmt.Println(wIndex)

// 	// replace (1 replacement)
// 	fmt.Println(strings.Replace(string1, "World", "Martin", 1))

// 	temperatureC := 32.0
// 	// temperatureK := 0.0

// 	temperatureF := temperatureC*9/5 + 32
// 	fmt.Println(temperatureF)

// 	// call pointer example
// 	myInt := 5
// 	changeVal(&myInt)
// 	fmt.Println("changed val via pointer", myInt)

// 	// more pointer examples
// 	f4 := 10
// 	var f4Ptr *int = &f4

// 	fmt.Println("f4 address :", f4Ptr)
// 	fmt.Println("f4 value :", *f4Ptr)

// 	fmt.Println("f4 before func :", f4)
// 	changeVal(&f4)
// 	fmt.Println("f4 after func :", f4)

// 	// double array in place example
// 	pArr := [4]int{1, 2, 3, 4}
// 	dblArrayVals(&pArr)
// 	fmt.Println(pArr)

// 	// slice varadic example
// 	iSlice := []float64{11, 13, 17}
// 	fmt.Println(getAverage(iSlice...))

// }

// simple pointer example
func changeVal(myPointer *int) {
	*myPointer = 12
}

// modify array in place using pointers
func dblArrayVals(arr *[4]int) {
	for x := 0; x < 4; x++ {
		arr[x] *= 2
	}
}

// slice
func getAverage(nums ...float64) float64 {
	var sum float64
	for _, val := range nums {
		sum += val
	}
	return sum / float64(len(nums))
}
