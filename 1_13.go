package begins

import (
	"fmt"
	"strconv"
)

func Task1131() {                               //Дано трехзначное число. Найдите сумму его цифр. 
  var (
		number int
		result int
	)

  fmt.Print("Enter your number: ")
  fmt.Scan(&number)
    
  if number > 1000 && number < 100 {
    fmt.Println("")
  	Task1131()
  }

  result = number % 10 + number / 100 + (number / 10) % 10  
    
  fmt.Println("Result: ", result)
}

func Task1132() {                                   //Дано трехзначное число. Переверните его, а затем выведите. 
  var number, digit1, digit2, digit3 int
    
  fmt.Print("Enter your number: ")
  fmt.Scan(&number)
    
  digit1 = number / 100
  digit2 = (number / 10) % 10
  digit3 = number % 10
    
  fmt.Printf("Inverted number: %d%d%d", digit3, digit2, digit1)
}

func Task1133() {                         //Идёт k-я секунда суток. Определите, сколько целых часов 
  var (                                   //h и целых минут m прошло с начала суток.
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

func Task1134() {                                     //Заданы три числа - a,b,c(a<b<c)a,b,c(a<b<c) - длины сторон треугольника. 
  var (                                               //Нужно проверить, является ли треугольник прямоугольным. Если является, 
    hypotenuse int                                    //вывести "Прямоугольный". Иначе вывести "Непрямоугольный".
    leg1 int
    leg2 int
  )
   
	fmt.Print("Enter the first leg, the second leg and the hypotenuse through a space: ")
  fmt.Scan(&leg1, &leg2, &hypotenuse)
    
  if hypotenuse < leg1 || hypotenuse < leg2 {
    Task1134()   
  }
    
  if leg1 * leg1 + leg2 * leg2 == hypotenuse * hypotenuse {
    fmt.Print("This triangle is right angled.")   
  } else {
    fmt.Print("This triangle is not a right triangle.")
  }
}

func Task1135() {                                     //Даны три натуральных числа a, b, c. Определите, существует 
  var (                                               //ли треугольник с такими сторонами.
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

func Task1136() {                             //Даны два числа. Найти их среднее арифметическое.
  var (

  )
}

func Task1137() {                             //По данным числам, определите количество чисел, которые равны нулю.  
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

func Task1138() {                                     //Найдите количество минимальных элементов в последовательности.
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

func Task1139() {                                     //По данному числу определите его цифровой корень.
var (

)
}

func Task11310() {                                    //Найдите самое большее число на отрезке от a до b, кратное 7 .
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

func Task11311() {                            //По данному числу n закончите фразу "На лугу пасется..." одним из возможных
  var numberOfCows int                        // продолжений: "n коров", "n корова", "n коровы", правильно склоняя слово "корова".
    
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
  }
}

func Task11312() {                            //По данному числу N распечатайте все целые значения степени двойки,
    var (                                     //не превосходящие N, в порядке возрастания.
        number int 
        power int = 1
    )
    
    fmt.Scan(&number)
    
    for power <= number {
        fmt.Print(power, " ")
        power *= 2
    }
}

func numberFibonacci() {            //Дано натуральное число A > 1. Определите, каким по счету числом Фибоначчи 
  var (                             //оно является, то есть выведите такое число n, что φn=A. Если А не является
    fibonacciNumber int              //числом Фибоначчи, выведите число -1.                                           
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

func Task11314() {              //Дано натуральное число N. Выведите его представление в двоичном виде.
  var number int64
  var binaryNumber string

  fmt.Print("Enter your number: ")
  fmt.Scan(&number)
    
  binaryNumber = strconv.FormatInt(number, 2)
    
  fmt.Println("Binary number:", binaryNumber)
}

func Task11315() {              //Из натурального числа удалить заданную цифру.              
var (

)

}

func Task113() {
	var userChoice string

  fmt.Print("Select the task number: ")
  fmt.Scan(&userChoice)

  switch userChoice {
  case "1":
  	Task1131()
  case "2":
    Task1132()
  case "3":
    Task1133()
  case "4":
    Task1134()
  case "5":
    Task1135()
	case "6":
		Task1136()
	case "7":
		Task1137()
	case "8":
		Task1138()
	case "9":
		Task1139()
	case "10":
		Task11310()
	case "11":
		Task11311()
	case "12":
		Task11312()
	case "13":
		Task11313()
	case "14":
		Task11314()
	case "15":
		Task11315()
  default:
    fmt.Println("Introduced a non-existent variant! Please try again.")
    Task110() 
  }
}