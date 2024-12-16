package main

import (
	"errors"
	"fmt"
)

const winePrice = 100

func main() {
	change, err := buyWine(20, 110)
	if err != nil {
		fmt.Println("Purchase is unsuccessful:", err.Error())
	} else {
		fmt.Printf("Here's your change - %d. Good day!\n", change)
	}

	change, err = buyWine(17, 110)
	if err != nil {
		fmt.Println("Purchase is unsuccessful:", err.Error())
	} else {
		fmt.Printf("Here's your change - %d. Good day!\n", change)
	}

	change, err = buyWine(33, 30)
	if err != nil {
		fmt.Println("Purchase is unsuccessful:", err.Error())
	} else {
		fmt.Printf("You back money - %d. Good day!\n", change)
	}
}

func buyWine(age, moneyAmount int) (int, error) {
	if age >= 18 && moneyAmount >= winePrice {
		return moneyAmount - winePrice, nil
	} else if age < 18 {
		return moneyAmount, errors.New("go drink juice")
	} else {
		return moneyAmount, errors.New("Not enough money")
	}
}
