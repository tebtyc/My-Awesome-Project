package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	error1 = "Выдача паники, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)."
	error2 = "Выдача паники, так как калькулятор производит операцию только с целыми числами одиноковой системы счисления от 1 до 10 включительно."
	error3 = "Вывод ошибки так как в римской системе счисления нет 0 и отрицательных чисел"
)

var rim = map[string]int{
	"I":        1,
	"II":       2,
	"III":      3,
	"IV":       4,
	"V":        5,
	"VI":       6,
	"VII":      7,
	"VIII":     8,
	"IX":       9,
	"X":        10,
	"XI":       11,
	"XII":      12,
	"XIII":     13,
	"XIV":      14,
	"XV":       15,
	"XVI":      16,
	"XVII":     17,
	"XVIII":    18,
	"XIX":      19,
	"XX":       20,
	"XXI":      21,
	"XXII":     22,
	"XXIII":    23,
	"XXIV":     24,
	"XXV":      25,
	"XXVI":     26,
	"XXVII":    27,
	"XXVIII":   28,
	"XXIX":     29,
	"XXX":      30,
	"XXXI":     31,
	"XXXII":    32,
	"XXXIII":   33,
	"XXXIV":    34,
	"XXXV":     35,
	"XXXVI":    36,
	"XXXVII":   37,
	"XXXVIII":  38,
	"XXXIX":    39,
	"XL":       40,
	"XLI":      41,
	"XLII":     42,
	"XLIII":    43,
	"XLIV":     44,
	"XLV":      45,
	"XLVI":     46,
	"XLVII":    47,
	"XLVIII":   48,
	"XLIX":     49,
	"L":        50,
	"LI":       51,
	"LII":      52,
	"LIII":     53,
	"LIV":      54,
	"LV":       55,
	"LVI":      56,
	"LVII":     57,
	"LVIII":    58,
	"LIX":      59,
	"LX":       60,
	"LXI":      61,
	"LXII":     62,
	"LXIII":    63,
	"LXIV":     64,
	"LXV":      65,
	"LXVI":     66,
	"LXVII":    67,
	"LXVIII":   68,
	"LXIX":     69,
	"LXX":      70,
	"LXXI":     71,
	"LXXII":    72,
	"LXXIII":   73,
	"LXXIV":    74,
	"LXXV":     75,
	"LXXVI":    76,
	"LXXVII":   77,
	"LXXVIII":  78,
	"LXXIX":    79,
	"LXXX":     80,
	"LXXXI":    81,
	"LXXXII":   82,
	"LXXXIII":  83,
	"LXXXIV":   84,
	"LXXXV":    85,
	"LXXXVI":   86,
	"LXXXVII":  87,
	"LXXXVIII": 88,
	"LXXXIX":   89,
	"XC":       90,
	"XCI":      91,
	"XCII":     92,
	"XCIII":    93,
	"XCIV":     94,
	"XCV":      95,
	"XCVI":     96,
	"XCVII":    97,
	"XCVIII":   98,
	"XCIX":     99,
	"C":        100,
}

var a, b *int
var operators = map[string]func() int{
	"+": func() int { return *a + *b },
	"-": func() int { return *a - *b },
	"*": func() int { return *a * *b },
	"/": func() int { return *a / *b },
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Input:")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	words := strings.Fields(text)
	var operator string
	for i, _ := range operators {
		for _, word2 := range words {
			if i == word2 {
				operator += i
			}
		}
	}

	if len(operator) != 1 {
		panic(error1)
	}

	x, _ := strconv.Atoi(words[0])
	y, _ := strconv.Atoi(words[2])

	chek := x > 0 && x < 11 && y > 0 && y < 11

	if result, ok := operators[operator]; ok && chek == true {
		a = &x
		b = &y
		fmt.Println("Output:\n", result())
	} else {
		num1 := words[0]
		num2 := words[2]
		if n1, ok := rim[num1]; ok {
			if n2, ok := rim[num2]; ok {
				chek := n1 > 0 && n1 < 11 && n2 > 0 && n2 < 11
				if result, ok := operators[operator]; ok && chek == true {
					a = &n1
					b = &n2
					intrim(result())
				} else {
					panic(error2)

				}
			} else {
				panic(error2)
			}

		} else {
			panic(error2)
		}
	}
}
func intrim(rimresult int) {

	if rimresult > 0 {
		for i, num := range rim {
			if num == rimresult {
				result := i
				fmt.Println("Outrut:\n", result)
			}
		}
	} else {
		panic(error3)
	}
}
