package main

import "fmt"

func main() {
	// Создаем переменную name и присваиваем ей значение "Danny"
	name := "Danny"
	fmt.Println(name) // Выводим значение переменной name

	// Вызываем функцию setName, передавая указатель на переменную name
	setName(&name) // &name - это указатель на переменную name

	// После вызова функции setName, значение переменной name изменилось
	fmt.Println(name) // Выводим новое значение переменной name
}

// Функция setName принимает указатель на строку (string)
func setName(name *string) {
	// Изменяем значение переменной, на которую указывает указатель
	*name = "Roma" // Разыменовываем указатель и присваиваем новое значение
}
