package main

import "fmt"

func main() {
	var name string
	var age int
	var weight float32
	var height float32

	name = "Danny"
	age = 32
	weight = 79.2
	height = 179.8

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(weight)
	fmt.Println(height)
}

func printPersonInfo(name string, age int, weight float32, height float32) {
	fmt.Printf("Name: %s\n Age: %d\n Weight: %f\n Height: %f\n", name, age, weight, height)

}
