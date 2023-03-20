package begins

import "fmt"

// 1.11 Форматированный вывод

func Task1111() {												//На вход подается число типа float64. Вам нужно вывести преобразованное число по правилу:
  var (																	//число возводится в квадрат, затем обрезается дробная часть так что остается 4 знака после 
		number float64											//запятой. Но если число равно или меньше нуля - выводить: "число R не подходит", где R - 
  	result float64											//исходное число с 2 цифрами после запятой и с 2 по ширине. А если число больше чем 10000 -
	)																			//выводить исходное число в экспоненциальном представлении (знак экспоненты в нижнем регистре).

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