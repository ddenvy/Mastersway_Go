package main

import "fmt"

func main() {
	// Создаем массив (array) с именем contactsList, который содержит 3 строки (имена контактов)
	contactsList := [3]string{"Tanya", "Victor", "Denis"}

	// Выводим заголовок для списка контактов
	fmt.Println("List of Contacts")

	// Используем цикл for с range для перебора элементов массива
	// range возвращает индекс (index) и значение (value) для каждого элемента массива
	for index, value := range contactsList {
		// Выводим номер контакта (индекс + 1, так как индексы начинаются с 0) и его имя
		fmt.Printf("%d. %s\n", index+1, value)
	}

	// Выводим длину массива (количество элементов) с помощью функции len()
	fmt.Println("Длинна массива", len(contactsList))
}
