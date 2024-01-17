package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	fmt.Println(
		"Это приложение калькулятор\n" +
			"Введите математическое выражение по примеру: '1 + 2'\n" +
			"Вы можете использовать римские(I,V,X) и арабские(1,5,10)\n" +
			"Числа могут быть от 1 до 10 включительно\n" +
			"Типы поддерживаемых операций: '+ - / *'")

	var num1, num2, symbol, close string

	fmt.Scanln(&num1, &symbol, &num2, &close)
	checkBase(num1, symbol, num2, close)
	controller(num1, symbol, num2)
}

func checkBase(num1, symbol, num2, close string) {
	if num2 == "" || close != "" {
		panic("неправильно введины данные")
		return
	}

	if symbol != "+" && symbol != "-" && symbol != "/" && symbol != "*" {
		panic("введина несуществующая операция")
		return
	}
}

func controller(num1, symbol, num2 string) {
	var rim bool
	rim, _ = regexp.MatchString("^[IVX]*$", num1)
	if rim == true {
		rim, _ = regexp.MatchString("^[IVX]*$", num2)
		if rim == true {
			//Римская империя
			var num1, num2 int = convertToInt(num1), convertToInt(num2)

			if checkNumMax(num1, num2) == true {
				var sum int = arithm(num1, num2, symbol)
				fmt.Println("Ответ: ", convertToString(sum))
				return
			}
		}
	}

	if num1, err := strconv.Atoi(num1); err == nil {
		if num2, err := strconv.Atoi(num2); err == nil {
			//Числа
			if checkNumMax(num1, num2) == true {
				fmt.Println("Ответ: ", arithm(num1, num2, symbol))
				return
			}
		}
	}

	panic("Введенные числа не подходят под критерии работы приложения")
}

func arithm(num1, num2 int, symbol string) int {
	switch symbol {
	case "+":
		return num1 + num2
	case "-":
		return num1 - num2
	case "*":
		return num1 * num2
	case "/":
		return num1 / num2
	default:
		panic("Ошибка связанная с символом")
	}
}

func convertToString(sum int) string {
	if sum <= 0 {
		panic("Число ниже нуля")
	}

	var number string
	for sum > 0 {
		if sum >= 100 {
			sum -= 100
			number += "C"
		} else if sum >= 90 {
			sum -= 90
			number += "XC"
		} else if sum >= 50 {
			sum -= 50
			number += "L"
		} else if sum >= 40 {
			sum -= 40
			number += "XL"
		} else if sum >= 10 {
			sum -= 10
			number += "X"
		} else if sum >= 9 {
			sum -= 9
			number += "IX"
		} else if sum >= 5 {
			sum -= 5
			number += "V"
		} else if sum >= 4 {
			sum -= 4
			number += "IV"
		} else if sum >= 1 {
			sum -= 1
			number += "I"
		}
	}
	return number
}

func convertToInt(num string) int {
	var sum, lastNum int = 0, 0
	for i := 0; i < len(num); i++ {
		var nowNum int = 0
		if string(num[i]) == "X" {
			nowNum = 10
			sum += nowNum
		} else if string(num[i]) == "V" {
			nowNum = 5
			sum += nowNum
		} else if string(num[i]) == "I" {
			nowNum = 1
			sum += nowNum
		}
		if lastNum < nowNum {
			sum = sum - lastNum - lastNum
			lastNum = nowNum
		}
	}
	return sum
}

func checkNumMax(num1, num2 int) bool {
	if num1 <= 10 && num2 <= 10 && num1 >= 0 && num2 >= 0 {
		return true
	}
	return false
}
