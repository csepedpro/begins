package main

import "fmt"

//2.3 Указатели

func Tasks23() {
  var userChoice string

  fmt.Print("Select the task number: ")
  fmt.Scan(&userChoice)

  switch userChoice {
  case "1":
		FunctionF()
  case "2":
    MinimumFromFour()
	default:
    fmt.Println("Introduced a non-existent variant! Please try again.")
    Tasks21()
  }
}

//Напишите функцию, которая умножает значения на которые ссылаются два указателя и 
//выводит получившееся произведение в консоль. Ввод данных уже написан за вас.	

func MultiplicationFunctionOfTwoNumbers() {
	
}

//Поменяйте местами значения переменных на которые ссылаются указатели. 
//После этого переменные нужно вывести.

func SwapVariables() {
	
}