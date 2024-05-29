package main

import "fmt"

func basics() {
	fmt.Println("Hello, World!")
	var age = 20        // automatically infer type
	var budget int = 20 // manually set type
	var name string     // manually set type but no value
	name = "roger"      // assign value
	// next line is the same as above
	desc := "this is a description" // create new variable

	const STATIC_VALUE = "Go is awesome" // consts are also possible

	height, weight := 180, 80 // multi-declaration in one line

	var total int = age * budget
	// not using a var causes an error when compiling
	fmt.Println(total, age, budget, name, height, weight)

	// strings are similar to php
	fmt.Printf("%s, length: %d, first letter: s, first ten: %s, last ten: %s\n", desc, len(desc) /*, desc[0] /* this should result in the letter at that position but is a byte for some reason*/, desc[:10], desc[10:])

	// arrays are single-type only
	arr := [...]string{"one", "two", "three"} // create a new array and let go figure out the size
	arr2 := [5]string{"one", "two", "three"}  // create a new array but set length manually
	fmt.Printf("length: %d, length2: %d, elem: %s, elem2: %s, elem2 (not initialized): %s\n", len(arr), len(arr2), arr[1], arr2[1], arr2[4])

	// slices are arrays with variable sizes
	slice := []int{10, 20, 30}
	fmt.Printf("slice elem: %d\n", slice[1])

	// maps are associative arrays:
	// myMap := map[string]int{"patrick": 25}
}

func loops() {
	// go only has 'for'
	for i := 0; i < 5; i++ {
		fmt.Printf("i: %d\n", i)
	}

	j := 0
	for { // or leave out the condition all together!
		fmt.Printf("j: %d\n", j)
		if j == 5 {
			break
		}
		j++
	}

	arr := [...]string{"one", "two", "three"} // create a new array and let go figure out the size
	// use range to loop over arrays
	for i, ar := range arr {
		fmt.Printf("index: %d, ar string: %s\n", i, ar)
	}
	// use underscore to tell go to 'ignore this'
	for _, ar := range arr {
		// i is now undefined
		fmt.Printf("index: $d, ar string: %s\n" /* , i */, ar)
	}
}

func conditionals() {
	i := 10
	// pretty standard if
	if i < 10 {
		fmt.Println("i is smaller than 10")
	} else if 1 > 10 {
		fmt.Println("i is bigger than 10")
	} else {
		fmt.Println("i is exactly than 10")
	}

	// there's also a switch
	switch i {
	case 10:
		fmt.Println("i is exactly than 10")
		// no 'break' needed
	default:
		fmt.Println("i is not exactly than 10")
	}

	// basic operators are as expected
}

// structs are a collection of properties (just like classes)
type Person struct { // using uppercase makes stuff public
	Age  int
	Name string
} // this is now usable just like any other type (eg int, string)

func structs() {
	// now we can use the struct to create new objects
	patrick := Person{25, "Patrick"}
	// OR
	thomas := Person{Name: "Thomas"}
	// access stuff with dot syntax
	fmt.Printf("name: %s, age: %d\n", patrick.Name, patrick.Age)
	fmt.Printf("name: %s, age: %d\n", thomas.Name, thomas.Age)

}

// name(params+type) return type {
func add(a int, b int) int {
	return a + b
}

// accept an undefined (unlimited) amount of parameters
func sum(numbers ...int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

// you can also return multiple values
func arithmetics(a int, b int) (int, int) {
	return a * b, a / b
}

func functions() {
	// call a function with params
	addedNums := add(11035, 31034)
	fmt.Printf("addedNums: %d\n", addedNums)

	sum := sum(1, 2, 3, 4, 5, 6)
	fmt.Printf("sum: %d\n", sum)

	multiplied, divided := arithmetics(10, 5)
	fmt.Printf("multiplied: %d, divided: %d\n", multiplied, divided)

}

func main() {
	basics()
	loops()
	conditionals()
	structs()
	functions()
}
