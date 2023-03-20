package begins

import "fmt"

// 1.5 Переменные и арифметические операции, ввод/вывод данных

func Task15() {
  var userChoice string

  fmt.Print("Select the tusk number: ")
  fmt.Scan(&userChoice)

  switch userChoice {
  case "1":
    Task151()
  case "2":
    Task152()
  case "3":
    Task153()
  case "4":
    Task154()
  case "5":
    Task155()
  default:
    fmt.Print("Introduced a non-existent variant! Please try again.")
    Task15()
  }
}

func Task151() {                  //Напишите программу, которая последовательно делает следующие операции
  var result int                  //с введённым числом: 1. Число умножается на 2; 2.Затем к числу прибавляется 100.
  var number int 
    
  fmt.Print("Enter your number: ")
  fmt.Scan(&number, "/n")
    
  result = (number * 2) + 100
    
  fmt.Println("Result: ", result)
}

func Task152() {                  //По данному целому числу, найдите его квадрат.
  var number int                  
  var result int
    
  fmt.Scan(&number)
    
  result = number * number
    
  fmt.Print("Result: ", result)
}

func Task153() {                 //Дано натруальное число, не превосходящее 1000, выведите его последнюю цифру.
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
    Task153()     
  }
}

func Task154() {                   // Дано неотрицательное целое число, не первосходящее 10000. 
  var number int                   // Найдите число десятков (то есть вторую цифру справа).
  var result int = 0

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
    Task154()     
  }
}

func Task155() {                  //Часовая стрелка повернулась с начала суток на d градусов.
  var numberOfHours int           //Определите, сколько сейчас целых часов h и целых минут m.
  var numberOfMinutes int         //На вход программе подается целое число d (0 < d < 360).
    
  fmt.Print()
  fmt.Scan(&numberOfHours)

  if numberOfHours > 360 {
    fmt.Print("Too large number entered! Please try again.")
    Task155()
  }
    
  numberOfMinutes = (numberOfHours % 30) * 2
  
  /*Если 360 градусов - это 12 часов (или 12 * 60 = 720 минут), то 1 градус - это две минуты,
  30 градусов - это один час, поэтому количество целых часов будет = d div 30, 
  а минуты получаются из формулы, по которой подсчитывается остаток: d mod 30.
  И этот результат надо умножить на 2, так как один градус - это 2 минуты.*/
  
  numberOfHours = numberOfHours / 30
    
  fmt.Print("It is ", numberOfHours, " hours ", numberOfMinutes, " minutes.")
}