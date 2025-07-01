package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Guessing Game")
	fmt.Println(
		"Um número aleatòrio será sorteado. Tente acertar. (0-100)",
	)

	x := rand.Int63n(101)
	scanner := bufio.NewScanner(os.Stdin)
	guesses := [10]int64{}

	for i := range guesses {
		fmt.Println("Qual é seu chute?")
		scanner.Scan()
		guess := scanner.Text()
		guess = strings.TrimSpace(guess)

		guessInt, err := strconv.ParseInt(guess, 10, 64)
		if err != nil {
			fmt.Println("O seu chute tem que ser um número inteiro ")
			return
		}

		switch {
		case guessInt < x:
			fmt.Println("Você errou. O número sorteado e maior que", guessInt)
		case guessInt > x:
			fmt.Println("Você errou. O número sorteado e menor que", guessInt)
		case guessInt == x:
			fmt.Printf("Parabens! Você acertou! O número era: %d\nVocê acertou em %d tentativas\n Essas foram as suas tentativas: %v\n", x, i+1, guesses[:i])
			return
		}

		guesses[i] = guessInt
	}

	fmt.Printf("Infelizmente, voce não acertou número, que era: %d\nVocê teve 10 tentativas.\n Essas foram as suas tentativas: %v\n", x, guesses)

}
