package main

import (
	"fmt"
	"math/rand"
)

/*
Guess the Number Game

Problem:
The program randomly selects a number between 1 and 100.
The user must guess the number in a limited number of tries (variable), receiving feedback:
- "Too low!" if the guess is below the number
- "Too high!" if the guess is above the number
- "Correct!" if the guess is right

If the user has lost, return "You have lost!" every time a guess is tried
If the user won, return "You have won!" every time a guess is tried

Input: integers from user input (guesses)
Output: feedback strings

Implement the Guess function that would pass the tests, as well as a
command line user interface that would generate a game (with user given
max retries), generate a secret random number, then let the user play the game.
*/

func main() {
	fmt.Println("guess the number")
	var random int = rand.Intn(100) + 1
	fmt.Println("Enter number of tries")
	var tries int
	fmt.Scan(&tries)
	game := NewGame(random, tries)
	var copyRes string
	var ok bool = false
	for {
		fmt.Println("Enter your guess")
		var guess int
		fmt.Scan(&guess)
		result := game.Guess(guess)
		if result == "You have lost!" || result == "You have won!" {
			if !ok {
				ok = true
				copyRes = result
			}
		}
		if !ok {
			fmt.Println(result)
		} else {
			fmt.Println(copyRes)
		}
	}
}
