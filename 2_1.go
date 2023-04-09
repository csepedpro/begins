package main

import "fmt"

//2.1 Функции

func Tasks21() {
  var userChoice string

  fmt.Print("Select the task number: ")
  fmt.Scan(&userChoice)

  switch userChoice {
  case "1":
		FunctionF()
  case "2":
    MinimumFromFour()
	case "3":
		VotingFunction()
	case "4":
		FunctionFibonaci()
	case "5":
		FunctionSumInt()
  default:
    fmt.Println("Introduced a non-existent variant! Please try again.")
    Tasks21()
  }
}

//Напишите функцию f(), которая будет принимать 
//строку text и выводить ее (печатать на экране).

func FunctionF(){
	var userText string

	fmt.Println("Enter your text)")
	fmt.Scan(userText)
	
	fmt.Print(userText)
}

//Напишите функцию, находящую наименьшее из четырех '
//введённых в этой же функции чисел.

func MinimumFromFour() int {
  var (
    number1 int
    number2 int
    number3 int
    number4 int
  )

  fmt.Print("Enter your numbers separated by a space: ")
  fmt.Scan(&number1, &number2, &number3, &number4)
    
  minNumber := number1
    
  for searchIndex := 0; searchIndex < 4; searchIndex++ {
    switch {
    case minNumber > number2:
      minNumber = number2
    case minNumber > number3:
      minNumber = number3
    case minNumber > number4:
      minNumber = number4          
    }
  }
  
	fmt.Print("The minimum number is ")
  return minNumber
}

//Напишите "функцию голосования", возвращающую то значение (0 или 1),
//которое среди значений ее аргументов x, y, z встречается чаще.

func VotingFunction() {
	var (
		x	 int
		y int
		z int
)

fmt.Scan(&x, &y, &z)

fmt.Print((x + y + z ) / 2)
}

//Напишите функцию, которая по указанному натуральному n возвращает φn.

func FunctionFibonaci() {
	var (                            
    NumberInTheFibonacciSequence int                                                       
    firstNumber int = 1
    secondNumber int = 1
		result int 
  )

  fmt.Print("Enter your number: ")
  fmt.Scan(&NumberInTheFibonacciSequence)

  for searchIndex := 2; searchIndex < NumberInTheFibonacciSequence; searchIndex++ {
    result := firstNumber + secondNumber
    
    firstNumber = secondNumber

    secondNumber = result
  }
  

	if NumberInTheFibonacciSequence == 1 || NumberInTheFibonacciSequence == 2 {
		fmt.Print("Your number: 1")	
	} else {
		fmt.Print("Your number: ", result)
	}	
}

//Напишите функцию sumInt, принимающую переменное количество аргументов типа int,
//и возвращающую количество полученных функцией аргументов и их сумму. Пакет 
//"fmt" уже импортирован, функция и пакет main объявлены.
  
func FunctionSumInt(numbers ... int)(int, int) {
  sumOfNumbers:= 0

  countOfNumbers := len(numbers)

  for _, number := range numbers {
    sumOfNumbers += number
  }

  return countOfNumbers, sumOfNumbers
}