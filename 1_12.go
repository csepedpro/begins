package begins

import "fmt"

// 1.12 Массивы и срезы 

func Task1121() {								//На первом этапе на стандартный ввод подается 10 целых положительных чисел, которые должны
  var (													//быть записаны в порядке ввода в массив из 10 элементов. Тип чисел, входящих в массив,
		firstNumber uint8						//должен соответствовать минимально возможному целому беззнаковому числу. Имя массива который
  	secondNumber uint8					//вы должны сами создать массив. На втором этапе на стандартный ввод подаются еще 3 пары 
  	replacementVariable uint8		//чисел - индексы элементов этого массива, которые требуется поменять местами (если такая
	)															//пара чисел 3 и 7, значит в массиве элемент с индексом 3 нужно поменять местами с элементом,
																//индекс которого 7). Элементы должны быть ввыведены через пробел на стандартный вывод. 
  var workArray [10]uint8
                                     
  for inputIndex := 0; inputIndex < len(workArray); inputIndex++ {
    fmt.Scan(&workArray[inputIndex])    
  }

  for permutationIndex := 0; permutationIndex < 3; permutationIndex++ {
    fmt.Print("Enter a couple of numbers: ")
    fmt.Scan(&firstNumber, &secondNumber)

    replacementVariable = workArray [firstNumber]
    
    workArray [firstNumber] = workArray[secondNumber]

    workArray [secondNumber] = replacementVariable
  }

  for outputIndex := 0; outputIndex < len(workArray); outputIndex++ {
    fmt.Print(workArray[outputIndex])    
  }
}

func Task1122() {                             //Напишите программу, принимающая на вход число N(N≥4)N(N≥4), а затем N
  var (                                         //целых чисел-элементов среза. На вывод нужно подать 4-ый (3-ий по индексу)
    number int  = 0                                       //элемент данного среза.
  )
  
  fmt.Print(number)

}

func Task1123() {                               //На ввод подаются пять целых чисел, которые записываются в массив. Однако 
  array := [5]int {}                            //эта часть программы уже написана. Вам нужно написать фрагмент кода, с помощью
                                                //которого можно найти и вывести максимальное число в этом массиве.
  var (                                        
    number int                                  
    maxNumber int                                  
  ) 
  
  for inputIndex := 0; inputIndex < 5; inputIndex++ {
		fmt.Scan(&number)
		array[inputIndex] = number
	}
    
  for searchIndex := 0;  searchIndex < 5; searchIndex++ {
    if searchIndex == 0 {
      maxNumber = array[searchIndex]
   }
        
    if array[searchIndex] > maxNumber {
      maxNumber = array[searchIndex]    
    }       
  }
      
    fmt.Print(maxNumber)
}

func Task1124() {                             //Дан массив, состоящий из целых чисел. Нумерация элементов начинается с 0.
  var amountOfNumbers int                     //Напишите программу, которая выведет элементы массива, индексы которых четны (0, 2, 4...).

  fmt.Print("Enter the number of elements in the array: ")
  fmt.Scan(&amountOfNumbers)
    
  var arrayOfNumbers []int = make([]int, int(amountOfNumbers))
  
  for inputIndex := 0; inputIndex < len(arrayOfNumbers); inputIndex++ {
    fmt.Scan(&arrayOfNumbers[inputIndex])    
  }
    
  for searchIndex := 0; searchIndex < len(arrayOfNumbers); searchIndex++ {
    if searchIndex == 0 && searchIndex % 2 == 0 {
      fmt.Print(arrayOfNumbers[searchIndex], " ")     
    }
  }
}

func Task1125() {                               //Дана последовательность, состоящая из целых чисел. Напишите программу, которая 
  var (                                         //подсчитывает количество положительных чисел среди элементов последовательности.
    positiveNumbers int
    amountOfNumbers int   
  )
  
  fmt.Print("Enter the number of elements in the array: ")
  fmt.Scan(&amountOfNumbers)
    
  var arrayOfNumbers []int = make([]int, int(amountOfNumbers))
    
  for inputIndex := 0; inputIndex < len(arrayOfNumbers); inputIndex++ {
    fmt.Scan(&arrayOfNumbers[inputIndex])    
  }
    
  for searchIndex := 0; searchIndex < len(arrayOfNumbers); searchIndex++ {
    if arrayOfNumbers[searchIndex] > 0 {
      positiveNumbers++    
    }
  }
    
  fmt.Println("Number of positive numbers: ", positiveNumbers)

}

func Task112() {
	var userChoice string

  fmt.Print("Select the task number: ")
  fmt.Scan(&userChoice)

  switch userChoice {
  case "1":
  	Task1121()
  case "2":
    Task1122()
  case "3":
    Task1123()
  case "4":
    Task1124()
  case "5":
    Task1125()
  default:
    fmt.Println("Introduced a non-existent variant! Please try again.")
    Task112() 
  }
}