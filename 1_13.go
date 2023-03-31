package main

import (
	"fmt"
	"strconv"
)

//Дано трехзначное число. Найдите сумму его цифр.

func Tasks113() {
	var userChoice string

  fmt.Print("Select the task number: ")
  fmt.Scan(&userChoice)

  switch userChoice {
  case "1":
  	PrintSumOfThreeDigitsOfNumber()
  case "2":
    PrintReversedThreeDigitNumber()  
  case "3":
    PrintNumberOfHoursAndMinutes()
  case "4":
    RightTriangleOrNot()
  case "5":
    TriangleExistsOrNot()  
	case "6":
    ArithmeticMeanOfTwoNumbers()
	case "7":
    CountTheNumberOfZeros()  
	case "8":
    NumberOfMinimumNumbers()  
	case "9":
    DetermineDigitalRoot()
	case "10":
    FindTheLargestMultipleOfSeven()                                                             
	case "11":
		TypeTheWordInTheCorrectDeclension()
	case "12":
		PowerOfTwoToNumberValues()
	case "13":
		FibonacciNumberOrNot()
	case "14":
		PrintNumberInBinary()
	case "15":
		DeleteGivenNumber()
  default:
    fmt.Println("Introduced a non-existent variant! Please try again.")
    Tasks113() 
  }
}

func PrintSumOfThreeDigitsOfNumber() {                                     
  var (
		number int
		result int
	)

  fmt.Print("Enter your number: ")
  fmt.Scan(&number)
    
  if number > 1000 && number < 100 {
    fmt.Println("Number doesn't fit. Please try again.")
  	PrintSumOfThreeDigitsOfNumber()
  }

  result = number % 10 + number / 100 + (number / 10) % 10  
    
  fmt.Println("Result: ", result)
}

 //Дано трехзначное число. Переверните его, а затем выведите. 

func PrintReversedThreeDigitNumber() {                                    
  var number, digit1, digit2, digit3 int
    
  fmt.Print("Enter your number: ")
  fmt.Scan(&number)
    
  digit1 = number / 100
  digit2 = (number / 10) % 10
  digit3 = number % 10
    
  fmt.Printf("Inverted number: %d%d%d", digit3, digit2, digit1)
}

//Идёт k-я секунда суток. Определите, сколько целых часов h и целых минут m прошло с начала суток.

func PrintNumberOfHoursAndMinutes() {                                   
  var (                                                                 
    dailySecond  int
    numberOfHours int
    numberOfMinutes int
  )

  fmt.Print("Enter daily second: ")
  fmt.Scan(&dailySecond)

  numberOfHours = dailySecond / 3600

  numberOfMinutes = (dailySecond - numberOfHours * 3600) / 60

  fmt.Print("It is ", numberOfHours, " hours", numberOfMinutes, " minutes")

}

//Заданы три числа - a,b,c(a<b<c)a,b,c(a<b<c) - длины сторон треугольника. Нужно проверить, 
//является ли треугольник прямоугольным. Если является, вывести "Прямоугольный". Иначе вывести "Непрямоугольный".

func RightTriangleOrNot() {                                              
  var (                                                                 
    hypotenuse int                                                      
    leg1 int
    leg2 int
  )
   
	fmt.Print("Enter the first leg, the second leg and the hypotenuse through a space: ")
  fmt.Scan(&leg1, &leg2, &hypotenuse)
    
  if hypotenuse < leg1 || hypotenuse < leg2 {
    RightTriangleOrNot()   
  }
    
  if leg1 * leg1 + leg2 * leg2 == hypotenuse * hypotenuse {
    fmt.Print("This triangle is right angled.")   
  } else {
    fmt.Print("This triangle is not a right triangle.")
  }
}

 //Даны три натуральных числа a, b, c. Определите, существует ли треугольник с таким сторонами.

func TriangleExistsOrNot() {                                           
  var (                                                                
    side1 int
    side2 int
    side3 int
  )
    
  fmt.Scan(&side1, &side2, &side3)
    
  if side1 < side2 + side3 && side3 < side2 + side1 && side2 < side3 + side1 {
    fmt.Print("This triangle exists.")   
  } else {
    fmt.Print("This triangle does not exist.")
  }
}

//Даны два числа. Найти их среднее арифметическое.

func ArithmeticMeanOfTwoNumbers() {                                                       
  var (
    number1 float64
    number2 float64
  )

  fmt.Print("Enter 2 numbers separated by a space: ")
  fmt.Scan(&number1, &number2)

  result := (number1 + number2) / 2.0

  //Число 2.0 используется для того, чтобы произвести деление
  //с плавающей точкой и получить результат также с плавающей точкой. 

  fmt.Println(result)
}

//По данным числам, определите количество чисел, которые равны нулю. 

func CountTheNumberOfZeros() {                                                        
  var amountOfNumber int
  var number int
  var numberOfZeros int

  fmt.Print("Enter amount of number: ")
  fmt.Scan(&amountOfNumber)
  
  for inputIndex := 0; inputIndex < amountOfNumber; inputIndex++ {
    fmt.Print("Enter number: ")
    fmt.Scan(&number)
    
    if number == 0 {
      numberOfZeros ++
    }
  }
  
  fmt.Print("Number of zeros: ", numberOfZeros)
}

//Найдите количество минимальных элементов в последовательности.

func NumberOfMinimumNumbers() {                                                       
  var (
    number int
    amountOfNumbers int
    quantity int
    minNumber int 
  )
    
  fmt.Scan(&amountOfNumbers)
    
  for inputIndex := 0; inputIndex < amountOfNumbers; inputIndex++ {
    fmt.Scan(&number)
    
    if inputIndex == 0 {
      minNumber = number
    }
        
    if number == minNumber {
      quantity++   
    }
    
    if number < minNumber {
      minNumber = number
      quantity = 1
    }
  }
    
  fmt.Print(quantity)  
}

//По данному числу определите его цифровой корень.

func DetermineDigitalRoot() {                                                         
  var number int
   
  fmt.Print("Enter your number: ")
  fmt.Scan(&number)
  
  result := (number - 1) % 9 + 1

  //https://codeforces.com/blog/entry/18286 - Формула для цифрового корня

  fmt.Print(result)
}

//Найдите самое большее число на отрезке от a до b, кратное 7.

func FindTheLargestMultipleOfSeven() {                                                        
  var (
    firstBorder int
    secondBorder int
    maxNumber int
    searchIndex int
    counter int = 0 
  )
    
  fmt.Scan(&firstBorder, &secondBorder)
  
  maxNumber = firstBorder - 1
  searchIndex = firstBorder
  
  for searchIndex < secondBorder + 1 {
    if searchIndex % 7 == 0 && searchIndex > maxNumber {
      maxNumber = searchIndex
      counter += 1  
    }

    searchIndex++
  }
  
  if counter == 0 {
      fmt.Print("There is no such number!")
  } else {
      fmt.Print("Max Number: ", maxNumber)
  }
}

//По данному числу n закончите фразу "На лугу пасется..." одним из возможных
// продолжений: "n коров", "n корова", "n коровы", правильно склоняя слово "корова".

func TypeTheWordInTheCorrectDeclension() {                                                          
  var numberOfCows int                                                      
    
  fmt.Scan(&numberOfCows)
    
  switch {
  case numberOfCows == 1:
    fmt.Print(numberOfCows, " korova")
  case (numberOfCows > 1 && numberOfCows < 5):
    fmt.Print(numberOfCows, " korovy")
  case numberOfCows >= 5 && numberOfCows <= 20:
    fmt.Print(numberOfCows, " korov")
  case numberOfCows > 20 && numberOfCows < 100 && numberOfCows % 10 == 1:
    fmt.Print(numberOfCows, " korova")
  case numberOfCows > 21 && numberOfCows < 100 && numberOfCows % 10 > 1 && numberOfCows % 10 < 5:
    fmt.Print(numberOfCows, " korovy")
  case numberOfCows > 24 && numberOfCows < 100:
    fmt.Print(numberOfCows, " korov")
  default:
    fmt.Print("Invalid input. Try again!")
    TypeTheWordInTheCorrectDeclension()
  }
}

//По данному числу N распечатайте все целые значения степени двойки,
//не превосходящие N, в порядке возрастания.

func PowerOfTwoToNumberValues() {                                                                  
    var (                                     
        number int 
        power int = 1
    )
    
    fmt.Scan(&number)
    
    for power <= number {
        fmt.Print(power, " ")
        power *= 2
    }
}

//Дано натуральное число A > 1. Определите, каким по счету числом Фибоначчи
//оно является, то есть выведите такое число n, что φn=A. Если А не является
//числом Фибоначчи, выведите число -1.

func FibonacciNumberOrNot() {          
  var (                            
    fibonacciNumber int                                                       
    firstNumber int = 1
    secondNumber int = 1
    flag bool = true
  )

  fmt.Print("Enter your number: ")
  fmt.Scan(&fibonacciNumber)

  for searchIndex := 1; searchIndex < fibonacciNumber; searchIndex++ {
    result := firstNumber + secondNumber
    
    firstNumber = secondNumber

    secondNumber = result

    if result == fibonacciNumber {
      fmt.Print(searchIndex + 1)
      flag = false
      break
    }  
  }
  
  if flag == false {
    fmt.Print("-1")  
  }  
}

//Дано натуральное число N. Выведите его представление в двоичном виде.

func PrintNumberInBinary() {              
  var number int64
  var binaryNumber string

  fmt.Print("Enter your number: ")
  fmt.Scan(&number)
    
  binaryNumber = strconv.FormatInt(number, 2)
    
  fmt.Println("Binary number:", binaryNumber)
}

//Из натурального числа удалить заданную цифру.

func DeleteGivenNumber() {                            
  var (
    digit string
    number string
    result string
  )

  fmt.Print("Enter your number: ")
  fmt.Scan(&number)

  fmt.Print("Enter the number to be deleted: ")
  fmt.Scan(&digit)

  for searchIndex := range number {
    if string(number[searchIndex]) != digit {
        result += (string(number[searchIndex]))
    }
  }
  
  /*Выражение "number[searchIndex]" обращается к символу строки с 
  индексом "searchIndex" и возвращает его в виде байта. Функция 
  "string()" преобразует этот байт обратно в строку.*/
  
  fmt.Print("Result: ", result)
}
