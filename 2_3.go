package main

import "fmt"

//2.3 Указатели

func Tasks23() {
  var userChoice string

  fmt.Print("Select the task number: ")
  fmt.Scan(&userChoice)

  switch userChoice {
  case "1":
    var (
      number1 int32
      number2 int32
    )

    fmt.Print("Enter two numbers separated by a space: ")
    fmt.Scan(&number1, &number2)
    
		MultiplicationFunctionOfTwoNumbers(&number1, &number2)
  case "2":
    var (
      number1 int
      number2 int
    )

    fmt.Print("Enter two numbers separated by a space: ")
    fmt.Scan(&number1, &number2)
    
    SwapVariables(&number1, &number2)
	default:
    fmt.Println("Introduced a non-existent variant! Please try again.")
    Tasks23()
  }
}

//Напишите функцию, которая умножает значения на которые ссылаются два указателя и 
//выводит получившееся произведение в консоль. Ввод данных уже написан за вас.	

func MultiplicationFunctionOfTwoNumbers(x1 *int32, x2 *int32) {
  fmt.Print("Result: ", *x1 * *x2)
}

//Поменяйте местами значения переменных на которые ссылаются указатели. 
//После этого переменные нужно вывести.

func SwapVariables(x1 *int, x2 *int) {
	n := *x1
    
  *x1 = *x2
  *x2 = n
  fmt.Print(*x1, " ",*x2)
}