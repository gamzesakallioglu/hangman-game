package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
	"unicode"
)

func readUserRune() (rune, error) {
	inputReader := bufio.NewReader(os.Stdin)
	char, _, err := inputReader.ReadRune()

	return char, err
}

func checkTheGuess(guess string, word string) bool {

	for _, v := range word {
		if string(v) == guess {
			return true
		}
	}
	return false
}

func checkIfExistsInSlice(guessedChar string, slice []string) bool {

	for _, v := range slice {
		if v == guessedChar {
			return true
		}
	}
	return false
}

func setTheWordArea(word string, rightGuesses []string) string {

	wordArea := ""

	for _, v := range word {

		if string(v) == " " {
			wordArea += " "
		} else {
			sayac := 0
			for _, v2 := range rightGuesses {
				if v2 == string(v) {
					sayac++
					wordArea += string(v)
				}
			}
			if sayac == 0 {
				wordArea += "_"
			}
		}
	}
	return wordArea
}

func setTheHangerArea(wrongGuesses []string) string {

	// To decide which state will be displayed: use the length of wrongGuesses slice

	wrongGuessTotal := len(wrongGuesses)
	fmt.Println("------- THE HANGER ----------")

	stateFile := fmt.Sprintf("example/states/hangman%d", wrongGuessTotal+2)
	state, err := os.ReadFile(stateFile)
	if err != nil {
		panic(err)
	}
	return string(state)
}

func setWrongGuessesArea(wrongGuesses []string) {

	fmt.Println("")
	if len(wrongGuesses) > 0 {
		fmt.Print("WRONG GUESSES SO FAR: ")

		for i, v := range wrongGuesses {
			fmt.Print(string(v))
			if i != len(wrongGuesses)-1 {
				fmt.Print(", ")
			}
		}
	}
	fmt.Println("")
}

func playTheGame(selectedWord string, remainedLives int) bool {
	gameContinues := true
	var playerWon bool
	wrongGuesses := []string{}
	rightGuesses := []string{}

	for gameContinues {
		guessIsValid := true

		fmt.Println()
		fmt.Println("Make Your Guess: ")
		char, err := readUserRune()
		if err != nil {
			fmt.Println(err)
		}
		guessedChar := string(char)

		fmt.Println()
		// If the guess exists in wrong guesses slice
		if checkIfExistsInSlice(guessedChar, wrongGuesses) {
			guessIsValid = false
			fmt.Println("You already made this guess and it was wrong.\nCome on you can do better")
		}

		if !unicode.IsLetter(char) {
			guessIsValid = false
			fmt.Println("Please only type letter. Rules are absolute!")
		}

		// If the guess exists in right guesses slice
		if checkIfExistsInSlice(guessedChar, rightGuesses) {
			guessIsValid = false
			fmt.Println("You already made this guess and it was right.\nGo on!")
		}

		if guessIsValid {
			// check if the word contains the char
			if !checkTheGuess(guessedChar, selectedWord) {
				remainedLives--
				wrongGuesses = append(wrongGuesses, guessedChar)
				fmt.Println("RUTTEN LUCK! TRY AGAIN")
			} else {
				rightGuesses = append(rightGuesses, guessedChar)
				fmt.Println("CORRECT!")
			}

			// set the hanger
			fmt.Println(setTheHangerArea(wrongGuesses))

			// set the word area
			fmt.Println(setTheWordArea(selectedWord, rightGuesses))

			// set the wrong guesses space
			setWrongGuessesArea(wrongGuesses)

			if remainedLives <= 0 {
				playerWon = false
				gameContinues = false
			}

			if checkIfPlayerWon(setTheWordArea(selectedWord, rightGuesses)) {
				playerWon = true
				gameContinues = false
			}
		}
	}
	return playerWon
}

func checkIfPlayerWon(checkStr string) bool {

	for _, v := range checkStr {
		if string(v) == "_" {
			return false
		}
	}
	return true
}

func main() {

	wantToPlay := true
	totalLives := 6
	wordsToGuess := []string{
		"book", "table", "tiger", "hello world", "hurricane", "snowy weather",
	}

	fmt.Println("--- WELCOME TO OUR LITTLE FUN GAME! ---\nYOU HAVE 5 LIVES TO BEGIN WITH\nMAKE A GUESS, LET'S GO!")

	for wantToPlay {
		// For every play; lives starts from total
		remainedLives := totalLives

		// start with selecting the word to be guessed
		rand.Seed(time.Now().Unix())
		selectedWord := fmt.Sprint(wordsToGuess[rand.Intn(len(wordsToGuess))])

		fmt.Println(selectedWord)

		// set the word to guess area
		fmt.Println(setTheWordArea(selectedWord, []string{}))

		// Play the game until win or lose - this function returns if the player win or lose
		playerWon := playTheGame(selectedWord, remainedLives)

		if playerWon {
			fmt.Println("WOW YOU WON!")
		} else {
			fmt.Println("YOU LOST! THE WORD WAS:", selectedWord, "YOU CAN TRY AGAIN TO BEAT ME (IF YOU CAN SOME DAY HAHA!)")
		}

		// ask if the person wants to play again
		fmt.Println()
		fmt.Println("DO YOU WANT TO PLAY AGAIN? TYPE Y (yes) OR N (no)")
		char, err := readUserRune()
		if err != nil {
			fmt.Println(err)
		}
		if strings.ToLower(string(char)) == "n" {
			wantToPlay = false
		}

	}
	// User doesn't want to play
	fmt.Println("")
	fmt.Println("It was nice to play with you. Come back any time.")
}
