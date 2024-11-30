package main

import (
	"errors"
	"fmt"
)

func dataType(run bool) {
	if !run {
		return
	}

	fmt.Println("=================== Data Type ===================")
	// #1
	var name string = "Nuzzy"
	// #2
	var age = 14
	// #3
	pi := 3.14
	// #4
	const score int = 79

	fmt.Println("String", name)
	fmt.Printf("String: %s, Integer: %d, Float: %f\n", name, age, pi)

	pi = 4
	// can't var pi = 4 or pi := 4
}

func controlStructure(run bool) {
	if !run {
		return
	}

	fmt.Println("=================== Control Structure ===================")
	pi := 3.14
	const score int = 79
	var age = 14

	if pi < 3 {
		pi = 3.14
	}

	var grade string
	switch {
	case score >= 80:
		grade = "A"
	case score >= 70:
		grade = "B"
	case score >= 60:
		grade = "C"
	default:
		grade = "F"
	}
	fmt.Printf("Grade: %s, Score: %d\n", grade, score)

	if sum := age + score; sum > 10 {
		fmt.Println("Sum > 10")
	}

	fmt.Printf("Number (For loop) :")
	for i := 1; i <= 10; i++ {
		fmt.Printf(" %d", i)
	}
	fmt.Printf("\n")

	fmt.Printf("Number (While loop) :")
	i := 1
	for i <= 10 {
		fmt.Printf(" %d", i)
		i++
	}
	fmt.Printf("\n")
}

func dataStructure(run bool) {
	if !run {
		return
	}

	fmt.Println("=================== Data Structure ===================")
	fmt.Println("----- Array -----")
	var array [3]int // -> Array
	array[2] = 4
	fmt.Println(array)
	fmt.Println(array[2])

	fmt.Println("----- Slice -----")
	slice := []int{10, 20, 30, 40, 50} // -> Slice
	slice = append(slice, 50)
	fmt.Println(slice)

	// Converting array to slice
	slice = array[:]

	fmt.Println("----- Map -----")
	myMap := make(map[string]int)

	// Add key-value pairs to the map
	myMap["apple"] = 5
	myMap["banana"] = 10
	myMap["orange"] = 8
	fmt.Println(myMap)
	fmt.Println(myMap["apple"])

	delete(myMap, "orange")
	for key, value := range myMap {
		fmt.Printf("%s -> %d\n", key, value)
	}

	val, ok := myMap["pear"]
	if ok {
		fmt.Println("Pear's value:", val)
	} else {
		fmt.Println("Pear not found in map")
	}

	fmt.Println("----- Struct -----")
	type Student struct {
		Name   string
		Weight int
		Height int
		Grade  string
	}

	var students [3]Student
	students[0].Name = "Mikelopster"
	students[0].Weight = 60
	students[0].Height = 180
	students[0].Grade = "F"

	students[1] = Student{
		Name:   "Alice",
		Weight: 55,
		Height: 165,
		Grade:  "A",
	}

	fmt.Println(students)

	fmt.Println("------- Function -------")
	sayHi("Nuzzy")
	sayHi("guys")

	add := func(a int, b int) int {
		return a + b
	}

	fmt.Printf("Add function %d\n", add(5, 3))

	fmt.Println("----- Receiver method -----")

	rect := Rectangle{Length: 10, Width: 5}

	fmt.Println(rect)
	fmt.Println("Area of rectangle:", rect.Area())

	fmt.Println("------- Interface -------")
	dog := Dog{Name: "Bif"}
	person := Person{Name: "Ali"}

	makeSound(dog)
	makeSound(person)

	fmt.Println("-------- Pointer --------")
	employee := Employee{Name: "Somsri", Salary: 400000}
	fmt.Println(employee)

	raiseSalary(&employee, 300000)
	fmt.Println(employee)

	fmt.Println("-------- Error handling --------")
	result, err := divideNum(5, 0)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		fmt.Printf("Result: %d\n", result)
	}

	result2, err2 := divideNum(5, 2)
	if err2 != nil {
		fmt.Printf("Error: %s\n", err2)
	} else {
		fmt.Printf("Result: %d\n", result2)
	}

}

// ------------ Function ------------ //
func sayHi(name string) {
	fmt.Printf("Hi from %s\n", name)
}

// ----------------------------------- //

// ------------ Receiver method ------------ //
type Rectangle struct {
	Length float64
	Width  float64
}

// Method with a receiver of type Rectangle
func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

// ----------------------------------------- //

// ------------ Interface ------------ //
type Speaker interface {
	Speak() string
}

type Dog struct {
	Name string
}

func (d Dog) Speak() string {
	return "Woof!"
}

type Person struct {
	Name string
}

func (p Person) Speak() string {
	return "Hey Yo!"
}

func makeSound(s Speaker) {
	fmt.Println(s.Speak())
}

// ------------------------------------ //

// ------------ Pointer ------------ //

type Employee struct {
	Name   string
	Salary int
}

func raiseSalary(e *Employee, raise int) {
	e.Salary += raise
}

// --------------------------------- //

// ------------ Pointer ------------ //
func divideNum(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Can not divide by zero")
	}
	return a / b, nil // nil = NULL
}

// --------------------------------- //

func main() {
	dataType(false)
	controlStructure(false)
	dataStructure(true)
}
