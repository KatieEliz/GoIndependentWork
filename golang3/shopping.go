package main

import (
	"fmt"
)

var (
	userInput    string
	updatedTotal int
)

func main() {
	trolley := []string{"Cheese", "Oranges", "Yoghurt", "Bread"}
	prices := []int{120, 30, 80, 60}

	origionalTotal := calculateItems(trolley, prices)
	fmt.Println("Your current total is: ", origionalTotal)

	askUser(trolley, prices, origionalTotal)

	newItems(updatedTotal, userInput)

}
func calculateItems(trolley []string, prices []int) int {
	origionalTotal := 0

	for i, item := range trolley {
		if item == "Cheese" {
			prices[i] += 120
		}
		if item == "Oranges" {
			prices[i] += 30
		}
		if item == "Yoghurt" {
			prices[i] += 80
		}
		if item == "Bread" {
			prices[i] += 60
		}
		origionalTotal += prices[i]
	}
	return origionalTotal
}
func askUser(trolley []string, prices []int, origionalTotal int) {
	updatedTotal := origionalTotal
	for {
		fmt.Println("Would you like to add anything from this aisle?(Y/N): Pasta, Noodles, Rice, Sauce")
		fmt.Scanf("%v\n", &userInput)
		if userInput == "Y" {
			fmt.Println("What would you like to add from this aisle?")
			fmt.Scanf("%v\n", &userInput)
			updatedTotal = newItems(updatedTotal, userInput)
			fmt.Println("Your updated total is:", updatedTotal)
		} else if userInput == "N" {
			fmt.Println("Do you want to checkout?(Y/N)")
			fmt.Scanf("%v\n", &userInput)
			updatedTotal = checkout(updatedTotal, userInput)
			break
		} else {
			fmt.Println("Enter Y or N")
		}

	}

}
func newItems(updatedTotal int, userInput string) int {
	total := updatedTotal

	if userInput == "Pasta" || userInput == "Rice" {
		total += 30
	}
	if userInput == "Noodles" {
		total += 35
	}
	if userInput == "Sauce" {
		total += 50
	}
	return total
}
func checkout(updatedTotal int, userInput string) int {

	if userInput == "Y" {
		fmt.Println("Your total is: ", updatedTotal)
	}
	if userInput == "N" {
		fmt.Println("")
	}
	return updatedTotal
}
