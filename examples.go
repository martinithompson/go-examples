package main

import (
	"fmt"
	"strings"
)

func main() {
	string1 := "Hello, World"
	string2 := "hello, world"

	// case insensitive match
	fmt.Println(strings.EqualFold(string1, string2))

	// index
	wIndex := strings.Index(string1, "W")
	fmt.Println(wIndex)

	// replace (1 replacement)
	fmt.Println(strings.Replace(string1, "World", "Martin", 1))

	temperatureC := 32.0
	// temperatureK := 0.0

	temperatureF := temperatureC*9/5 + 32
	fmt.Println(temperatureF)

	// call pointer example
	myInt := 5
	changeVal(&myInt)
	fmt.Println("changed val via pointer", myInt)

	// more pointer examples
	f4 := 10
	var f4Ptr *int = &f4

	fmt.Println("f4 address :", f4Ptr)
	fmt.Println("f4 value :", *f4Ptr)

	fmt.Println("f4 before func :", f4)
	changeVal(&f4)
	fmt.Println("f4 after func :", f4)

	// double array in place example
	pArr := [4]int{1, 2, 3, 4}
	dblArrayVals(&pArr)
	fmt.Println(pArr)

	// slice varadic example
	iSlice := []float64{11, 13, 17}
	fmt.Println(getAverage(iSlice...))

	// map example
	// var myMap map [keyType]valueType

	var heroes map[string]string
	heroes = make(map[string]string)
	villains := make(map[string]string)

	heroes["Batman"] = "Bruce Wayne"
	heroes["Superman"] = "Clark Kent"
	heroes["The Flash"] = "Barry Allen"
	villains["Lex Luthor"] = "Lex Luthor"

	superPets := map[int]string{1: "Krypto", 2: "Bat hound"}

	fmt.Printf("batman is: %v\n", heroes["Batman"])
	fmt.Printf("Chip is: %v\n", superPets[3])

	_, ok := superPets[3]
	fmt.Println("Is there a 3rd pet: ", ok)

	// range over a map
	for k, v := range heroes {
		fmt.Printf("%s is %s\n", k, v)
	}
	// delete from a map
	delete(heroes, "The Flash")

	// generics example
	fmt.Println("5 + 4 = ", getSumGen(5, 4))
	fmt.Println("5.7 + 4.8 = ", getSumGen(5.7, 4.8))

	// struct example
	var tS customer
	tS.name = "Tom Smith"
	tS.address = "5 main street"
	tS.bal = 2345.43

	// call tsp to ML using defined types
	fmt.Printf("3 tsp = %.2f mL\n", tspToML(3))

	// using associate method - better way
	tsp1 := Tsp(3)
	fmt.Printf("3 tsp = %.2f mL\n", tsp1.ToMLs())

	// using interface example
	var kitty Animal
	kitty = Cat("Kitty")
	kitty.AngrySound()

}

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

// generics
type MyContstraint interface {
	int | float64
}

func getSumGen[T MyContstraint](x T, y T) T {
	return x + y
}

// structs - like classes
type customer struct {
	name    string
	address string
	bal     float64
}

// defined types
type Tsp float64
type TBs float64
type ML float64

func tspToML(tsp Tsp) ML {
	return ML(tsp * 4.92)
}

// better solution to above function - using associate methods:
func (tsp Tsp) ToMLs() ML {
	return ML(tsp * 4.92)
}

// interface example
type Animal interface {
	AngrySound()
	HappySound()
}

type Cat string

func (c Cat) Attack() {
	fmt.Println("cat attacks its prey")
}

func (c Cat) Name() string {
	return string(c)
}

func (c Cat) AngrySound() {
	fmt.Println("cat says hiss")
}

func (c Cat) HappySound() {
	fmt.Println("cat says meow")
}
