package main

import "fmt"

// 1.10 Циклы в Go

func Task1101() {											  //Напишите программу, которая выводит квадраты натуральных чисел от 1 до 10.
  for index := 1; index < 11; index++ { //от 1 до 10. Квадрат каждого числа должен выводится в новой строке.
    fmt.Println(index * index)
  } 
}

func Task1102() {												//Написать программу, при выполнении которой с клавиатуры считываются два 
  var summ int = 0											//натуральных числа A и B (каждое не более 100, A < B).  
  var number1 int												//Вывести сумму всех чисел от A до B  включительно.
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

func Task1103() {													//Напишите программу, которая в последовательности чисел находит сумму двузначных чисел,
  var (																		//кратных 8. Программа в первой строке получает на вход число n - количество чисел в
		amountOfNumbers int										//в последовательности, во второй строке - n чисел, входящих в данную последовательность.
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

func Task1104() {											//Последовательность состоит из натуральных чисел и завершается числом 0.
  var (															  //Определите количество элементов этой последовательности, которые равны
		number = 1												//её наибольшему элементу.
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


func Task1105() {											//Найдите первое число от 1 до n включительно, кратное c, но НЕ кратное d.			
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

func Task1106() {         //Напишите программу, которая считывает целые числа с консоли по одному числу в строке.
  var number int	 				//1. Если число меньше 10, то пропускаем это число; 2. Если число больше 100, то прекращаем 
													//считывать числа; 3. в остальных случаях вывести это число обратно на консоль в отдельной строке.
  for {
    fmt.Scan(&number)

    if number > 100 { 
			break
    }

    if number >= 10 {
      fmt.Print(number)
    }
  }  
}

func Task1107() {								//Вклад в банке составляет x рублей. Ежегодно он увеличивается на p процентов,	
  var (													//после чего дробная часть копеек отбрасывается. Каждый год сумма вклада
		bankDeposit int							// становится больше. Определите, через сколько лет вклад составит не менее y рублей.
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

func Task1108() {
  
}

func Task110() {
	var userChoice string

  fmt.Print("Select the task number: ")
  fmt.Scan(&userChoice)

  switch userChoice {
  case "1":
  	Task1101()
  case "2":
    Task1102()
  case "3":
    Task1103()
  case "4":
    Task1104()
  case "5":
    Task1105()
	case "6":
		Task1106()
	case "7":
		Task1107()
	case "8":
		Task1108()
  default:
    fmt.Println("Introduced a non-existent variant! Please try again.")
    Task110() 
  }
}

