package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

var (
	userInput string
	attempts  int = 0
)

func randomNum(min, max int) int {
	secret := rand.Intn(max-min) + min
	return secret
}
func main() {
	numToGuess := randomNum(1, 100)
	for {
		attempts++
		fmt.Println("Guess a number between 1-100😎")
		fmt.Scanf("%v\n", &userInput)

		usersNum, err := strconv.ParseInt(userInput, 10, 64)
		if err != nil {
			fmt.Println("You must enter an integer value")
		} else if numToGuess == int(usersNum) {
			break
		} else if int(usersNum) < numToGuess {
			fmt.Println("Your number is too small, try again")
		} else {
			fmt.Println("Your number is too big, try again")
		}
	}
	fmt.Printf("YES!✔ %v is the correct number, you guessed it after %v attempts\n", numToGuess, attempts)
}
