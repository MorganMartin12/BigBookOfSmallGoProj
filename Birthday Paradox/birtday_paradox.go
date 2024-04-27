package main

import (
	"fmt"
	"math/rand/v2"
)

func generateBirthdays(days int) []int {
	birthdays := make([]int, days)
	for i := 0; i < days; i++ {
		dayOfYear := rand.IntN(365) + 1
		birthdays[i] = dayOfYear
	}
	return birthdays
}

func shareBirthdays(birthdays []int, resch chan bool) {
	bitVector := make([]bool, 366)
	for _, day := range birthdays {
		if bitVector[day] {
			resch <- true
			return
		}
		bitVector[day] = true
	}
	resch <- false
}
func main() {
	count := 0
	resultChan := make(chan bool)
	for i := 0; i < 10000000; i++ {
		go shareBirthdays(generateBirthdays(30), resultChan)
	}
	fmt.Println("started processing results")
	for i := 0; i < 10000000; i++ {
		if <-resultChan == true {
			count += 1
		} else {
			count -= 1
		}
	}

	fmt.Println(float64(count) / float64(10000000))
}
