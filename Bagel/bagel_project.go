package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"math"
	"math/rand/v2"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

func generateNumber(size int) string {
	min := int(math.Pow10(size - 1))
	return strconv.Itoa(rand.IntN(9*min) + min)
}

func validateGuess(guess string, size int) (bool, string) {
	errorMessage := fmt.Sprintf("Must enter a valid number of size %d", size)
	if len(guess) != size {
		return false, errorMessage
	}
	if !regexp.MustCompile(`^\d+$`).MatchString(guess) {
		return false, errorMessage
	}
	return true, ""
}

func startGameMessage(oneDigitWrongPosition, oneDigitRightSpot, allWrong string, maxNumberofGuesses int) {
	fmt.Println("I am thinking of a 3-digit number. Try to guess what it is.")
	fmt.Println("Here are some clues:")
	fmt.Println("When I say:   That means:")
	fmt.Printf("%s   One digit is correct but in the wrong position.\n", oneDigitWrongPosition)
	fmt.Printf("%s  One digit is correct and in the right position.\n", oneDigitRightSpot)
	fmt.Printf("%s No digit is correct.\n", allWrong)
	fmt.Printf("You have %d guesses to get it.\n", maxNumberofGuesses)
}
func mapCharIndexes(secretNumber string) map[rune]int {
	charIndexMap := make(map[rune]int)
	for index, char := range secretNumber {
		charIndexMap[char] = index
	}
	return charIndexMap
}
func printFeedback(guess string, charIndexMap map[rune]int) {
	feedback := make([]string, len(guess))
	isBagel := true
	oneDigitRightSpot := os.Getenv("ONE_DIGIT_CORRECT_RIGHT_SPOT")
	oneDigitWrongSpot := os.Getenv("ONE_DIGIT_CORRECT_WRONG_SPOT")
	for i, char := range guess {
		if index, ok := charIndexMap[char]; ok {
			isBagel = false
			if index == i {
				feedback[i] = oneDigitRightSpot
			} else {
				feedback[i] = oneDigitWrongSpot
			}
		}
	}

	if isBagel {
		fmt.Println(os.Getenv("NO_DIGITS_CORRECT"))
	} else {
		fmt.Println(strings.Join(feedback, " "))
	}
}

func createGame() error {
	size, err := strconv.Atoi(os.Getenv("SIZE"))
	if err != nil {
		return fmt.Errorf("invalid size from env: %v", err)
	}
	maxNumberofGuesses, err := strconv.Atoi(os.Getenv("MAX_NUMBER_GUESSES"))
	if err != nil {
		return fmt.Errorf("invalid max number of guesses from env: %v", err)
	}

	secretNumber := generateNumber(size)
	startGameMessage(os.Getenv("ONE_DIGIT_CORRECT_WRONG_SPOT"), os.Getenv("ONE_DIGIT_CORRECT_RIGHT_SPOT"), os.Getenv("NO_DIGITS_CORRECT"), maxNumberofGuesses)

	guesses := 0
	charIndexMap := mapCharIndexes(secretNumber)

	for guesses < maxNumberofGuesses {
		fmt.Printf("Guess #%d: ", guesses+1)
		var guess string
		fmt.Scan(&guess)
		isValid, errorMsg := validateGuess(guess, size)
		if !isValid {
			fmt.Println(errorMsg)
			continue
		}
		if guess == secretNumber {
			fmt.Println("Congratulations! You guessed the number correctly!")
			return nil
		}
		printFeedback(guess, charIndexMap)
		guesses++
	}
	fmt.Println("You ran out of guesses.")
	fmt.Println("The answer was", secretNumber)
	return nil
}

func main() {
	for {
		if err := createGame(); err != nil {
			log.Fatal(err)
		}
		fmt.Println("Do you want to play again? (yes or no)")
		var playAgain string
		fmt.Scanln(&playAgain)
		if strings.ToLower(playAgain) != "yes" && strings.ToLower(playAgain) != "y" {
			fmt.Println("Thanks for playing")
			return
		}
	}
}
