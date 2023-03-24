package main

import "fmt"

// 1.5 Переменные и арифметические операции, ввод/вывод данных

func Tasks15() {
  var userChoice string

  fmt.Print("Select the task number: ")
  fmt.Scan(&userChoice)

  switch userChoice {
  case "1":
    SequenceOperationWithNumber()
  case "2":
    FindTheSquareOfNumber()
  case "3":
    PrintTheLastDigitOfNumber()
  case "4":
    FindTheNumberOfTens()
  case "5":
    NumberOfHoursAndMinutesInDegrees()
  default:
    fmt.Print("Introduced a non-existent variant! Please try again.")
    Tasks15()
  }
}

//Напишите программу, которая последовательно делает следующие операции
//с введённым числом: 1. Число умножается на 2; 2.Затем к числу прибавляется 100.

func SequenceOperationWithNumber() {                  
  var result int                  
  var number int 
    
  fmt.Print("Enter your number: ")
  fmt.Scan(&number, "/n")
    
  result = (number * 2) + 100
    
  fmt.Println("Result: ", result)
}

//По данному целому числу, найдите его квадрат.

func FindTheSquareOfNumber() {                  
  var number int                  
  
  fmt.Scan(&number)
        
  fmt.Print("Result: ", number * number)
}

//Дано натруальное число, не превосходящее 1000, выведите его последнюю цифру.

func PrintTheLastDigitOfNumber() {                 
  var number int
  var result int

  fmt.Print("Enter your number: ")
  fmt.Scan(&number)
  
  switch {
  case number >= 100 || number < 100:
    result = number % 100
    fmt.Print("Result: ", result)
  case number < 10:
    result = number
    fmt.Print("Result: ", result)
  default:
    fmt.Print("Program error! Enter the correct number.")
    PrintTheLastDigitOfNumber()     
  }
}

// Дано неотрицательное целое число, не первосходящее 10000.
// Найдите число десятков (то есть вторую цифру справа).

func FindTheNumberOfTens() {                    
  var number int                               
  var result int

  fmt.Print("Enter your number: ")
  fmt.Scan(&number)
  
  switch {
  case number >= 1000 && number < 10000:
    result = (number % 100) / 10
    fmt.Print("Result: ", result)
  case number < 1000:
    result = (number % 100) / 10
    fmt.Print("Result: ", result)
  case number < 10:
    result = 0
    fmt.Print("Result: ", result)
  default:
    fmt.Print("Program error! Enter the correct number.")
    FindTheNumberOfTens()     
  }
}

//Часовая стрелка повернулась с начала суток на d градусов.
//Определите, сколько сейчас целых часов h и целых минут m.
//На вход программе подается целое число d (0 < d < 360).

func NumberOfHoursAndMinutesInDegrees() {                 
  var numberOfHours int           
  var numberOfMinutes int         
    
  fmt.Print()
  fmt.Scan(&numberOfHours)

  if numberOfHours > 360 {
    fmt.Print("Too large number entered! Please try again.")
    NumberOfHoursAndMinutesInDegrees()
  }
    
  numberOfMinutes = (numberOfHours % 30) * 2
  
  /*Если 360 градусов - это 12 часов (или 12 * 60 = 720 минут), то 1 градус - это две минуты,
  30 градусов - это один час, поэтому количество целых часов будет = d div 30, 
  а минуты получаются из формулы, по которой подсчитывается остаток: d mod 30.
  И этот результат надо умножить на 2, так как один градус - это 2 минуты.*/
  
  numberOfHours = numberOfHours / 30
    
  fmt.Print("It is ", numberOfHours, " hours ", numberOfMinutes, " minutes.")
}