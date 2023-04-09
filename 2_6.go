package main

import "fmt"

// 2.6 Обработка ошибок

//Вы должны вызвать функцию divide которая делит два числа, но она 
//возвращает не только результат деления, но и информацию об ошибке. 

func ErrorFunction(){
	var number1, number2 int 

	fmt.Scan (&number1, &number2)

  result, err := divide(number1, number2)

  if err != nil{ 
		fmt.Print("Oшибка!!! Повторите попытку!")
		return
  }

	fmt.Print(result)
}

func divide(a int, b int)(int, error) {
	return a/b, nil
}