package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func main() {
	for true {
		clearConsole()
		fmt.Println("wpisz swoje kombinacje w formacie (1/0)-(1/0) ile bloków z lewej i ile z prawej strony")

		finalResult := 0
		for i := 0; i < 3; i++ {
			fmt.Println("kombinacja")

			reader := bufio.NewReader(os.Stdin)
			text, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Error reading input:", err)
				continue
			}

			left, right, err := extractNumbers(text)
			if err != nil {
				fmt.Println("Error extracting numbers:", err)
				continue
			}

			leftResult, err := convertPental(left)
			if err != nil {
				fmt.Println("Error converting left number:", err)
				continue
			}

			rightResult, err := convertPental(right)
			if err != nil {
				fmt.Println("Error converting right number:", err)
				continue
			}
			finalResult = finalResult + leftResult - rightResult
			fmt.Println("converted left", leftResult, "converted right ", rightResult, "current final result ", finalResult)
		}
		fmt.Println("finalResult ", finalResult)

		fmt.Println("Wciśnij coś aby zacząć od nowa...")
		bufio.NewReader(os.Stdin).ReadString('\n')
	}
}

func extractNumbers(text string) (int, int, error) {
	values := strings.Split(text, "-")
	if len(values) != 2 {
		return 0, 0, errors.New("input must be in format (1/0)-(1/0)")
	}

	left, err := parseNumber(values[0])
	if err != nil {
		return 0, 0, fmt.Errorf("error parsing left number: %w", err)
	}

	right, err := parseNumber(values[1])
	if err != nil {
		return 0, 0, fmt.Errorf("error parsing right number: %w", err)
	}

	return left, right, nil
}

func parseNumber(text string) (int, error) {
	text = strings.TrimSpace(text)

	x, err := strconv.ParseInt(text, 10, 32)
	if err != nil {
		return 0, fmt.Errorf("invalid number: %w", err)
	}

	return int(x), nil
}

func convertPental(pentalNumber int) (int, error) {
	length := 0
	decimalNumber := 0
	for pentalNumber > 0 {
		digit := getLastDigit(pentalNumber)
		power := length
		if digit >= 5 {
			return 0, errors.New("pental number contains a digit >= 5, which is invalid")
		}
		numberToAppend := digit * powerOf(5, power)
		decimalNumber += numberToAppend

		pentalNumber /= 10
		length++

	}

	return decimalNumber, nil
}

func getLastDigit(number int) int {
	return number % 10
}

func powerOf(number int, power int) int {
	solution := 1
	for power > 0 {
		solution = solution * number
		power--

	}
	return solution
}

func clearConsole() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
