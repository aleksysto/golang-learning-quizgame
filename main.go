package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

type question struct {
	q string
	a string
	p int
}

func removeElement(qArray []question, index int) []question {
	qArray[index] = qArray[len(qArray)-1]
	return qArray[:len(qArray)-1]
}

var questions = []question{
	{"What year marks the time of UNIX? \n", "1970", 10},
	{"Which company created the Macintosh? \n", "apple", 5},
	{"Which programming language is named after a snake? \n", "python", 5},
	{"What is love? \n", "baby don't hurt me", 25},
	{"Which programming language is named after a jewel? \n", "ruby", 5},
	{"What is the capital of Poland? \n", "warsaw", 50},
}

func main() {
	qArray := make([]question, len(questions))
	copy(qArray, questions)
	var name string
	var age int
	var score int
	fmt.Println("Welcome to the quiz game")
	fmt.Println("Enter your name and age: ")

	fmt.Scanf("%v %d", &name, &age)

	if age < 18 {
		fmt.Println("Only users of age are allowed to play")
		return
	}
	if name == "" {
		fmt.Println("Please enter a proper name")
		return
	}

	fmt.Printf("%v, aged %v, welcome to the game \n", name, age)
	scanner := bufio.NewScanner(os.Stdin)
	for true {
		if len(qArray) == 0 {
			var percent float64 = (float64(score) / 50) * 100
			fmt.Printf("Finished. \n Your final score is: %d or %.0f %% \n", score, percent)
			fmt.Println("Want to play again? y/n")
			var restart string
			fmt.Scan(&restart)
			if strings.ToLower(restart) == "y" {
				qArray = make([]question, len(questions))
				copy(qArray, questions)
				score = 0
				fmt.Println(qArray, questions)
			} else {
				fmt.Println("See you next time!")
				break
			}
		}
		random := rand.Intn(len(qArray))
		q := qArray[random].q
		a := qArray[random].a
		p := qArray[random].p
		qArray = removeElement(qArray, random)
		fmt.Println(q)
		var answer string
		for scanner.Scan() {
			answer = scanner.Text()
			if answer != "" {
				break
			}
		}
		if strings.ToLower(answer) == a {
			score += p
			fmt.Printf("Correct answer! \nCurrent score: %d \n", score)
		} else {
			fmt.Printf("Wrong answer :( \nCurrent score: %d \n", score)
		}
	}
	return
}
