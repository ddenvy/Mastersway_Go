package main

import "fmt"

// Определяем структуру contact, которая содержит информацию о контакте
type contact struct {
	firstName   string // Имя
	lastName    string // Фамилия
	phoneNumber string // Номер телефона
	email       string // Электронная почта
	address     string // Адрес
	jobTitle    string // Должность
}

// Метод printInfo для структуры contact, который выводит информацию о контакте
func (c contact) printInfo() {
	fmt.Printf("Имя: %s\nФамилия: %s\nНомер телефона:  %s\nЭл почта:  %s\nАдресс:  %s\nРабота:  %s\n",
		c.firstName, c.lastName, c.phoneNumber, c.email, c.address, c.jobTitle)
}

// Метод setName для структуры contact, который изменяет имя контакта
// Используется указатель (*contact), чтобы изменить оригинальный объект
func (c *contact) setName(name string) {
	c.firstName = name
}

func main() {
	// Создаем экземпляр структуры contact с начальными данными
	c1 := contact{
		firstName:   "Вася",
		lastName:    "Пупкин",
		phoneNumber: "13322345678",
		email:       "nagibator@yandex.ru",
		address:     "Ул Кукуева д5, г.Иваново",
		jobTitle:    "Сантехник",
	}

	// Используем метод setName для изменения имени контакта
	c1.setName("Петя")

	// Выводим информацию о контакте с помощью метода printInfo
	c1.printInfo()
}
