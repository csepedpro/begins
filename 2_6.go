package main

import "fmt"

// 2.6 Обработка ошибок

//Вы должны вызвать функцию divide которая делит два числа, но она 
//возвращает не только результат деления, но и информацию об ошибке. 

func ErrorFunction(){
	var number1, number2 int 

	fmt.Scan (&number1, &number2)

  result, error := divide(number1, number2)

  if error == nil{
    fmt.Print(result)
  }else{
	  fmt.Print("Oшибка!!! Повторите попытку!")
		ErrorFunction()
  }
}

func divide(a int, b int)(int, error) {
	return a/b, nil
}