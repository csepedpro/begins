package begins

import "fmt"

func main() {
	var userChoice string = ""
	
	fmt.Println("Choose a task: ")
	fmt.Scan(&userChoice)
	
	switch userChoice {
	case "1.3":
	  Task13()
	case "1.5":
	  Task15()
	case "1.9":
	  Task19()
	case "1.10":
	  Task110()
	case "1.11":
	  Task1111()
	case "1.12":
	  Task112()
	case "1.13":
	  Task113()
	default:
	  fmt.Println("Introduced a non-existent variant! Please try again.")
	  main()
	}
}
