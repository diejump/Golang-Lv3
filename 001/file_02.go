package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	max := 100
	rand.Seed(time.Now().UnixNano())
	secret_number := rand.Intn(max)
	println("Please input your guess")
	var guess_number, count int = 0, 3
	for count > 0 {
		_, err := fmt.Scanln(&guess_number)
		if guess_number == 114514 {
			println("The secret number is ", secret_number) //作弊模式
		}
		if err != nil {
			fmt.Println("Invalid input. Please enter an integer value")
			count--
			println(count, "times only")
			continue
		}
		if guess_number == secret_number {
			println("Your guess is", guess_number, ", you win!")
			break
		} else {
			count--
			println("Your guess is", guess_number, ", wrong answer!")
			println(count, "times only")
			if count == 0 {
				println("You lose, the number is ", secret_number)
			} else {
				println("Please input your guess")
			}

		}
	}

}
