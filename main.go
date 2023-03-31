package main

import "fmt"

func main() {
	var userChoice string 
	
	fmt.Println("Choose a task: ")
	fmt.Scan(&userChoice)
	
	switch userChoice {
	case "1.3":
	  Tasks13()
	case "1.5":
	  Tasks15()
	case "1.9":
	  Tasks19()
	case "1.10":
	  Tasks110()
	case "1.11":
	  CorrectSquareOfNumber()
	case "1.12":
	  Tasks112()
	case "1.13":
	  Tasks113()
	case "2.1":
		Tasks21()
	case "2.3":
		Tasks23()
	case "2.4":
		FunctionWithStructure()
	case "2.5":
		Tasks25()
	case "2.6":
		ErrorFunction()
	case "2.7":
		Tasks27()
	default:
	  fmt.Println("Introduced a non-existent variant! Please try again.")
	  main()
	}
}
