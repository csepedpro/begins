package main

import "fmt"

// 1.3. Первая программа

func Task13() {
  var userChoice string

  fmt.Print("Select the tusk number: ")
  fmt.Scan(&userChoice)

  switch userChoice {
  case "1":
    Task131()
  case "2":
    Task132()
  default:
    fmt.Println("Introduced a non-existent variant! Please try again.")
    Task13()
  }
}

func Task131() { 											//Напишите программу, которая выводит "I like Go!".
	fmt.Print("I like Go!")
}

func Task132() { 											 //Напишите программу, которая выведет "I like Go!" 3 раза.
	for outputIndex := 0; outputIndex < 3; outputIndex++ { 
		fmt.Println("I like Go!")
	}
}