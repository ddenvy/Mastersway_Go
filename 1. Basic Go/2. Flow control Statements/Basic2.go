package main

import (
	"errors"
	"fmt"
)

// Константа, определяющая стоимость бутылки вина
const winePrice = 100

func main() {
	// Пример 1: Покупка вина с возрастом 20 и достаточным количеством денег
	change, err := buyWine(20, 110)
	if err != nil {
		// Если произошла ошибка, выводим сообщение об ошибке
		fmt.Println("Purchase is unsuccessful:", err.Error())
	} else {
		// Если покупка успешна, выводим сдачу и пожелание хорошего дня
		fmt.Printf("Here's your change - %d. Good day!\n", change)
	}

	// Пример 2: Покупка вина с возрастом 17 (недостаточно для покупки)
	change, err = buyWine(17, 110)
	if err != nil {
		fmt.Println("Purchase is unsuccessful:", err.Error())
	} else {
		fmt.Printf("Here's your change - %d. Good day!\n", change)
	}

	// Пример 3: Покупка вина с возрастом 33, но недостаточным количеством денег
	change, err = buyWine(33, 30)
	if err != nil {
		fmt.Println("Purchase is unsuccessful:", err.Error())
	} else {
		fmt.Printf("Here's your change - %d. Good day!\n", change)
	}
}

// Функция buyWine принимает возраст и количество денег, возвращает сдачу и ошибку
func buyWine(age, moneyAmount int) (int, error) {
	// Проверяем, достаточно ли возраста и денег для покупки вина
	if age >= 18 && moneyAmount >= winePrice {
		// Если условия выполнены, возвращаем сдачу и nil (ошибки нет)
		return moneyAmount - winePrice, nil
	} else if age < 18 {
		// Если возраст меньше 18, возвращаем ошибку "go drink juice"
		return moneyAmount, errors.New("go drink juice")
	} else {
		// Если денег недостаточно, возвращаем ошибку "Not enough money"
		return moneyAmount, errors.New("Not enough money")
	}
}
