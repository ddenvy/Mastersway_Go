package main

import "fmt"

func main() {
	// Объявляем переменные для хранения информации о человеке
	var name string    // Имя (строка)
	var age int        // Возраст (целое число)
	var weight float32 // Вес (число с плавающей точкой)
	var height float32 // Рост (число с плавающей точкой)

	// Присваиваем значения переменным
	name = "Danny"
	age = 32
	weight = 79.2
	height = 179.8

	// Выводим значения переменных на экран
	fmt.Println(name)   // Вывод имени
	fmt.Println(age)    // Вывод возраста
	fmt.Println(weight) // Вывод веса
	fmt.Println(height) // Вывод роста

	// Вызываем функцию для вывода информации о человеке
	printPersonInfo(name, age, weight, height)
}

// Функция для вывода информации о человеке
func printPersonInfo(name string, age int, weight float32, height float32) {
	// Используем fmt.Printf для форматированного вывода
	fmt.Printf("Name: %s\n Age: %d\n Weight: %f\n Height: %f\n", name, age, weight, height)
}
