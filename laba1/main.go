package main

import (
	"errors"
	"fmt"
	"time"
)

// Не совсем точная постановка задачи касательно времени и даты
func main() {
	//region #1 Exercise
	fmt.Println("current date + time:")
	fmt.Println(time.Now().Local().Format(time.RFC1123))

	//region #2-3 Exercise
	someInt, someFloat64, someString, someBool := 13, 12.28, "nigger", false
	fmt.Println(someInt, someFloat64, someString, someBool)

	//region #4-5 Exercise
	val1, _ := doSomeOperation(3, 2, "+")
	fmt.Println(val1)

	val2, _ := doSomeOperation(3, 2, "-")
	fmt.Println(val2)

	val3, _ := doSomeOperation(3, 2, "*")
	fmt.Println(val3)

	val4, err := doSomeOperation(3, 2, "/")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(val4)

	//Проверка на ноль, могу сделать return, но прервёт выполнение кода
	val5, err := doSomeOperation(3, 0, "/")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(val5)

	//region #6 Exercise

	fmt.Println("bla")
	var num1 int = 3
	var num2 int = 2
	var num3 int = 5

	fmt.Println(float64(num1+num2+num3) / 3)
}

func doSomeOperation(num1, num2 int, opertaion string) (float64, error) {
	switch opertaion {
	case "-":
		return float64(num1 - num2), nil
	case "+":
		return float64(num1 + num2), nil
	case "*":
		return float64(num1 * num2), nil
	case "/":
		if num2 != 0 {
			return float64(num1) / float64(num2), nil
		}
		return -1, errors.New("не дели на 0")
	}

	return -1, errors.New("что-то пошло не так")
}

func floadOpers(num1, num2 float64) (float64, float64) {
	return num1 + num2, num1 - num2
}
