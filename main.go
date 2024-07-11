package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func sum(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func multy(a, b int) int {
	return a * b
}

func div(a, b int) int {
	return a / b
}

func romanNumToInt(rome string) (operand int, err error) {

	switch rome {
	case "I":
		operand = 1
	case "II":
		operand = 2
	case "III":
		operand = 3
	case "IV":
		operand = 4
	case "V":
		operand = 5
	case "VI":
		operand = 6
	case "VII":
		operand = 7
	case "VIII":
		operand = 8
	case "IX":
		operand = 9
	case "X":
		operand = 10
	default:
		err = fmt.Errorf("Enter Roman numerals from I to X!")
	}

	return
}

func intToRomanNum(num int) string {

	var romas string = ""
	var numbers = []int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}
	var romasArr = []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}
	index := len(romasArr) - 1

	for num > 0 {
		for numbers[index] <= num {
			romas += romasArr[index]
			num -= numbers[index]
		}
		index -= 1
	}

	return romas
}

func isRomans(num string) bool {

	var romanNumerals = []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}

	for _, l := range  romanNumerals {
		if num == l {
			return true
		}
	}

	return false
}

func getDataType(text string, op string) (a, b int, isRoman bool, err error) {

	input := strings.Split(text, op)

	if len(input) > 2 {

		return a, b, isRoman, fmt.Errorf("Many elements!")
	}

	firstRomeNum := isRomans(input[0])
	secondRomeNum := isRomans(input[1])

	if firstRomeNum != secondRomeNum {
		return a, b, isRoman, fmt.Errorf("Different number systems are used simultaneously!")
	}

	if firstRomeNum && secondRomeNum {

		isRoman = true
		a, err = romanNumToInt(input[0])
		if err != nil {
			panic(err)
		}
		b, err = romanNumToInt(input[1])
		if err != nil {
			panic(err)
		}

	} else {

		a, err = strconv.Atoi(input[0])
		if err != nil {
			panic(err)
		}

		b, err = strconv.Atoi(input[1])
		if err != nil {
			panic(err)
		}
	}

	if a < 0 || a > 10 || b < 0 || b > 10 {

		return a, b, isRoman, fmt.Errorf("a <= 0 or a > 10 or b <= 0 or b > 10")
	}

	return a, b, isRoman, nil

}

func findOperator(text string) (string, error) {

	switch {
	case strings.Contains(text, "+"):

		return "+", nil
	case strings.Contains(text, "-"):

		return "-", nil
	case strings.Contains(text, "*"):

		return "*", nil
	case strings.Contains(text, "/"):

		return "/", nil
	default:
		if len(text) == 1 {
			return "", fmt.Errorf("A string is not a mathematical operation")
		} else {
			return "", fmt.Errorf("operator not find")
		}

	}
}

func calculate(a, b int, operator string) (result int, err error) {

	switch operator {

	case "+":

		result = sum(a, b)
	case "-":

		result = sub(a, b)
	case "*":

		result = multy(a, b)
	case "/":
		if b == 0 {

			err = fmt.Errorf("operand = 0")
		} else {

			result = div(a, b)
		}

	default:
		err = fmt.Errorf("operator not found")
	}

	return
}

func main() {

	reader := bufio.NewReader(os.Stdin)

	for {

		fmt.Println("Введите значение")
		fmt.Println("Введите ex для выхода!")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		text = strings.ReplaceAll(text, " ", "")

		if text == "ex" {
			fmt.Println("Good bye!")
			return
		}

		operator, err := findOperator(text)
		if err != nil {
			panic(err)
		}

		a, b, isRoman, err := getDataType(text, operator)
		if err != nil {
			panic(err)
		}

		result, err := calculate(a, b, operator)
		if err != nil {
			panic(err)
		}

		if isRoman {
			if result < 0 {
				panic("result < 0 !!!")
			}

			res := intToRomanNum(result)
			fmt.Println(res)

		} else {
			fmt.Println(result)
		}
	}
}