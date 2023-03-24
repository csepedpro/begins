package main

import "fmt"

// 1.3. Первая программа

func Tasks13() {
  var userChoice string

  fmt.Print("Select the task number: ")
  fmt.Scan(&userChoice)

  switch userChoice {
  case "1":
    PrintILikeGoOnce()
  case "2":
    PrintILikeGoThreeTimes()
  default:
    fmt.Println("Introduced a non-existent variant! Please try again.")
    Tasks13()
  }
}

//Напишите программу, которая выводит "I like Go!"

func PrintILikeGoOnce() { 											
	fmt.Print("I like Go!")
}

//Напишите программу, которая выведет "I like Go!" 3 раза.

func PrintILikeGoThreeTimes() { 											 
	for outputIndex := 0; outputIndex < 3; outputIndex++ { 
		fmt.Println("I like Go!")
	}
}