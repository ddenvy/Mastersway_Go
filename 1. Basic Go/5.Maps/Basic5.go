package main

import "fmt"

func main() {
	// Создаем карту (map) с названиями животных в качестве ключей и их идентификаторами в качестве значений
	animalList := map[string]string{
		"Dog":   "0001",
		"Cat":   "0002",
		"Cow":   "0003",
		"Horse": "0004",
	}

	// Используем цикл for с range для перебора элементов карты
	// range возвращает ключ (name) и значение (idAnimal) для каждой пары в карте
	for name, idAnimal := range animalList {
		// Выводим ключ (название животного) и значение (идентификатор животного)
		fmt.Println(name, idAnimal)
	}

	// Удаляем элемент с ключом "Cow" из карты
	delete(animalList, "Cow")

	// Выводим сообщение о том, что элемент был удален
	fmt.Println("After Delete()")

	// Снова используем цикл for с range для перебора элементов карты
	// После удаления элемента "Cow" он больше не будет отображаться
	for name, idAnimal := range animalList {
		fmt.Println(name, idAnimal)
	}
}
