package main

import "fmt"

func main() {
	// Создаем срез (slice) с названиями продуктов
	productList := []string{"Phone", "Keyboard", "Mouse", "Dynamic", "Monitor", "Computer", "Laptop", "Microphone", "Modem", "Camera"}

	// Выводим заголовок для списка продуктов
	fmt.Println("List of Products:")

	// Используем цикл for с range для перебора элементов среза
	// range возвращает индекс (index) и значение (value) для каждого элемента среза
	for index, value := range productList {
		// Выводим номер продукта (индекс + 1, так как индексы начинаются с 0) и его название
		fmt.Printf("%d. %s\n", index+1, value)
	}

	// Добавляем новые элементы в срез с помощью функции append
	productList = append(productList, "HDD", "Earphones")

	// Используем цикл for для перебора элементов среза
	// len(productList) возвращает длину среза
	for i := 0; i < len(productList); i++ {
		// Выводим каждый элемент среза через пробел
		fmt.Printf("%s ", productList[i])
	}
}
