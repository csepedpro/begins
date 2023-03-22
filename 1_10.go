package main

import "fmt"

// 1.10 Циклы в Go

func Tasks110() {
	var userChoice string

  fmt.Print("Select the task number: ")
  fmt.Scan(&userChoice)

  switch userChoice {
  case "1":
  	SquareOfNumbersFromOneToTen()
  case "2":
    SumOfTheNumbersInTheRange()
  case "3":
    SumOfTwoDigitMultiplesOfEight()
  case "4":
    TheNumberOfMaximumNumbersInTheSequence()
  case "5":
    FirstТumberToAnotherNumberWithTwoMultiples()
	case "6":
		CheckingGivenConditions()
	case "7":
		NumberOfYearsToDeposit()
	case "8":
		PrintTotalNumbers()
  default:
    fmt.Println("Introduced a non-existent variant! Please try again.")
    Tasks110() 
  }
}

 //Напишите программу, которая выводит квадраты натуральных чисел от 1 до 10.
 //от 1 до 10. Квадрат каждого числа должен выводится в новой строке.

func SquareOfNumbersFromOneToTen() {											 
  for numberIndex:= 1; numberIndex < 11; numberIndex++ { 
    fmt.Println(numberIndex * numberIndex)
  } 
}

//Написать программу, при выполнении которой с клавиатуры считываются два натуральных числа
// A и B (каждое не более 100, A < B). Вывести сумму всех чисел от A до B  включительно.

func SumOfTheNumbersInTheRange() {												
  var summ int = 0											  
  var number1 int												
  var number2 int

  fmt.Print("Enter the first number: ")
  fmt.Scan(&number1)

  fmt.Print("Enter the second number: ")
  fmt.Scan(&number2)
    
  for index := number1 ; index < (number2 + 1); index++ {
    summ = summ + index
  }
    
  fmt.Println(summ)    
}

//Напишите программу, которая в последовательности чисел находит сумму двузначных чисел,
//кратных 8. Программа в первой строке получает на вход число n - количество чисел в
// последовательности, во второй строке - n чисел, входящих в данную последовательность.

func SumOfTwoDigitMultiplesOfEight() {													
  var (																		
		amountOfNumbers int										
  	sumOfNumbers int
	)

  fmt.Scan(&amountOfNumbers)
    
  var arrayOfNumbers []int = make([]int, int(amountOfNumbers))
  
  /*Длина является частью типа массива; он должен оцениваться как
  неотрицательная константа, представляемая значением типа int.
  Используйте срез вместо типа массива и знайте, что вы должны
  использовать целочисленный тип для длины*/                                          
  
  for inputIndex := 0; inputIndex < amountOfNumbers; inputIndex++ {
    fmt.Scan(&arrayOfNumbers[inputIndex])    
  }
    
  for sumIndex := 0; sumIndex < amountOfNumbers; sumIndex++ {
    if (arrayOfNumbers[sumIndex] % 8 == 0) && (arrayOfNumbers[sumIndex] < 100) && (arrayOfNumbers[sumIndex] >= 10) {
      sumOfNumbers = sumOfNumbers + arrayOfNumbers[sumIndex]   
    }
  } 
  
  fmt.Println("Sum of numbers: ", sumOfNumbers) 
}

//Последовательность состоит из натуральных чисел и завершается числом 0.
//Определите количество элементов этой последовательности, которые равны её наибольшему элементу.

func TheNumberOfMaximumNumbersInTheSequence() {											
  var (															 
		number = 1												
		countMaxNumber = 0
  	maxNumber = 0
	) 
  
  for inputIndex :=0; number != 0; inputIndex++ {
    fmt.Scan(&number)

  	if (number == 0 && inputIndex == 0) || number == 0 {
    	fmt.Println(countMaxNumber)
    break
  	}
        
  	if number > maxNumber {
    	maxNumber = number
    	countMaxNumber = 0
  	}
	}
	
  if number == maxNumber {
    countMaxNumber++   
  }
}

//Найдите первое число от 1 до n включительно, кратное c, но НЕ кратное d.

func FirstТumberToAnotherNumberWithTwoMultiples() {													
  var (
		amountOfNumbers int
		firstDivisor int
		secondDivisor int
	)

	fmt.Print("Enter the number of numbers to check: ")
  fmt.Scan(&amountOfNumbers)

	fmt.Print("Enter the first divisor: ")
  fmt.Scan(&firstDivisor)

	fmt.Print("Enter the second divisor")
  fmt.Scan(&secondDivisor)
    
  for numberIndex := 1; numberIndex < amountOfNumbers + 1; numberIndex++ {
    if (numberIndex % firstDivisor == 0) && (numberIndex % secondDivisor != 0) {
      fmt.Println(numberIndex)
      break
    }
  } 
}

//Напишите программу, которая считывает целые числа с консоли по одному числу в строке.
//1. Если число меньше 10, то пропускаем это число; 2. Если число больше 100, то прекращаем 
//считывать числа; 3. в остальных случаях вывести это число обратно на консоль в отдельной строке.

func CheckingGivenConditions() {         
  var number int	 				
													
  for {
    fmt.Print("Entet your number: ")
    fmt.Scan(&number)

    if number > 100 { 
			break
    }

    if number >= 10 {
      fmt.Print(number)
    }
  }  
}

//Вклад в банке составляет x рублей. Ежегодно он увеличивается на p процентов,
//после чего дробная часть копеек отбрасывается. Каждый год сумма вклада
//становится больше. Определите, через сколько лет вклад составит не менее y рублей.

func NumberOfYearsToDeposit() {									
  var (													
		bankDeposit int						
  	annualPercentage int
  	requiredAmount int 	
		numberOfYear int
	)

  fmt.Print("Enter deposit amount: ")
  fmt.Scan(&bankDeposit)

  fmt.Print("Enter annual percentage: ")
  fmt.Scan(&annualPercentage)

  fmt.Print("Enter required amount: ")
  fmt.Scan(&requiredAmount)
  
  for {
    bankDeposit += (bankDeposit % annualPercentage / 100)

		numberOfYear += 1
        
    if bankDeposit >= requiredAmount {
      break
    }
  }
  
  fmt.Println("After ", numberOfYear, "years, you will receive the required amount.")
}

//Даны два числа. Определить цифры, входящие в запись как первого, так и второго числа.

func PrintTotalNumbers() {
  var (
    number1 string 
    number2 string 
  )
  fmt.Print("Enter 2 numbers separated by a space: ") 
  fmt.Scan(&number1, &number2)
  
  for firstSearchIndex:=0 ; firstSearchIndex < len(number1) ; firstSearchIndex++ {
      for secondSearchIndex:=0 ; secondSearchIndex < len(number2) ; secondSearchIndex++ {
          if number1[firstSearchIndex] == number2[secondSearchIndex] {
              fmt.Print(string(number1[firstSearchIndex]), " ")
              break
          }
      }
  }
}

