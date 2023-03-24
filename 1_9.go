package main

import "fmt"

// 1.9 Условные конструкции 

func Tasks19() {
  var userChoice string

  fmt.Print("Select the task number: ")
  fmt.Scan(&userChoice)

  switch userChoice {
  case "1":
  	PositiveNegativeNumberOrZero()
  case "2":
    TheDigitsOfTheNumberAreSeparate()
  case "3":
    PrintTheFirstDigitOfNumber()
  case "4":
    LuckyTicketOrNot()
  case "5":
    LeapYearOrNot() 
  default:
    fmt.Println("Introduced a non-existent variant! Please try again.")
    Tasks19() 
  }
}

//На ввод подается целое число. Если число положительное - вывести сообщение "Число положительное",
//если число отрицательное - "Число отрицательное". Если подается ноль - вывести сообщение "Ноль" 
//Выводить сообщение без кавычек.

func PositiveNegativeNumberOrZero() {												
  var number int											
	                                      
	fmt.Print("Enter your number: ")			
  fmt.Scan(&number)
    
  switch {
  case number > 0:
    fmt.Println("The number is positive")
  case number == 0: 
    fmt.Println("Null")
  case number < 0:
    fmt.Println("The number is negative")  
  default:
    fmt.Println("Incorrect input! Try again.")
    PositiveNegativeNumberOrZero()
  }
}

//По данному трехзначному числу определите, все ли его цифры различны.

func TheDigitsOfTheNumberAreSeparate() {												
  var number int

  fmt.Println("Enter your number: ")
  fmt.Scan(&number)
  
  switch {
  case (number / 100 != number % 10) && (number / 100 != (number % 100) / 10 ) && (number % 10 != (number % 100) / 10 ):
    fmt.Println("All numbers are different.")
  case (number / 100 == number % 10) && (number / 100 == (number % 100) / 10 ) && (number % 10 == (number % 100) / 10 ):
    fmt.Println("Not all numbers are different.")
  default:
    fmt.Println("Incorrect input! Try again.")
    TheDigitsOfTheNumberAreSeparate()
  }  
}

//Дано неотрицательное целое число. Найдите и выведите первую цифру числа.

func PrintTheFirstDigitOfNumber() {												
  var number int 

  fmt.Print("Enter youe number: ")
  fmt.Scan(&number)
    
  switch {
  case number >= 1000 && number < 10000:
    fmt.Println(number / 1000)
  case number >= 100 && number < 1000:
    fmt.Println(number / 100)
  case number >= 10 && number < 100:
    fmt.Println(number / 10)            
  case number >= 0 && number < 10:
    fmt.Println(number)
  case number < 0:
    fmt.Println("The number is negative")
  case number == 10000:
        fmt.Println(number / 10000)
  default:
    fmt.Println("Incorrect input! Try again.")
    PrintTheFirstDigitOfNumber()
  }
}

//Определите является ли билет счастливым. Счастливым считается билет в шестизначном
//номере которого сумма первых трёх цифр совпадает с суммой трёх последних.

func LuckyTicketOrNot() {												
  var (																	
			numberOfTicket int												
   		sumOfNumber1 int
			sumOfNumber2 int
	)
		
  fmt.Scan(&numberOfTicket)
  
  if numberOfTicket < 100000 || numberOfTicket > 1000000 {
    fmt.Println("Incorrect input! Try again.")
    LuckyTicketOrNot() 
  }
  
  sumOfNumber1 = numberOfTicket / 100000 + (numberOfTicket / 10000) % 10 + ((numberOfTicket / 1000) % 100) % 10
  sumOfNumber2 = numberOfTicket % 10 + (numberOfTicket % 100) / 10 + ((numberOfTicket % 1000) / 100)
  
  if (sumOfNumber1 == sumOfNumber2) {
    fmt.Println("Your ticket is lucky!")
  } else {
    fmt.Println("Your ticket isn't lucky!")  
  }
}

//Требуется определить, является ли данный год високосным, напомним:
//Год является високосным если он соответствует хотя бы одному из
//нижеперечисленных условий: 1) кратен 400; 2) кратен 4, но не кратен 100.

func LeapYearOrNot() {													
  var year int														
																					
  fmt.Scan(&year)

  if year <= 0 || year > 10000 {
    fmt.Println("Incorrect input! Try again.")
    LeapYearOrNot()
  }

  if year % 400 == 0 || (year % 100 != 0 && year % 4 == 0) {
    fmt.Println("This year is a leap year!")
  } else {
    fmt.Println("This year is not a leap year!")
  }
}