package main

import "fmt"

// 1.11 Форматированный вывод

//На вход подается число типа float64. Вам нужно вывести преобразованное число по правилу:
//число возводится в квадрат, затем обрезается дробная часть так что остается 4 знака после 
//запятой. Но если число равно или меньше нуля - выводить: "число R не подходит", где R - 
//выводить исходное число в экспоненциальном представлении (знак экспоненты в нижнем регистре).

func CorrectSquareOfNumber() {												
  var (																	 
		number float64											
  	result float64											
	)																			

	fmt.Print("Enter your number: ")
  fmt.Scan(&number)
    
  switch {
  case number > 0 && number < 10000:
    result = number * number
	  fmt.Printf("%.4f", result)
   case number <= 0:
    fmt.Printf("%s %.2f %s", "number", number, "doesn't fit")
   default:
    fmt.Printf("%e", number)
  }     
}