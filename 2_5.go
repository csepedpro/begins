package main

import (
	"bufio"
	Console "fmt"
	"os"
	"strings"
	"unicode"
)

// 2.5 Cтроки (string)

func Tasks25() {
	var userChoice string

	Console.Print("Select the task number: ")
	Console.Scan(&userChoice)

	switch userChoice {
	case "1":
		DetermineTheCorrectWordOrNot()	
	case "2":
		DeterminePalindromeOrNot()
	case "3":
		DetermineIfStringIsInAnotherString()
	case "4":
		RemoveEvenCharactersInString() 
	case "5":
		DetermineIfThePasswordIsCorrect()
	default:
		Console.Println("Introduced a non-existent variant! Please try again.")
		Tasks25()
	}
}

//На вход подается строка. Нужно определить, является ли она правильной или нет. Правильная
//строка начинается с заглавной буквы и заканчивается точкой. Если строка правильная -
//вывести Right иначе - вывести Wrong.

func DetermineTheCorrectWordOrNot() {

	Console.Print("Enter your phrase: ")
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	byteText := []rune(text)

	/* Передается в функцию bufio.NewReader, на основании которого создается
	объект bufio.Reader. Поскольку идет построчное считывание, то каждая
	строка считывается из потока, пока не будет обнаружен символ перевода строки \n.
	Подробнее: https://metanit.com/go/tutorial/8.9.php */

	if unicode.IsUpper(byteText[0]) == true && byteText[len(byteText)-1] == '.' {
		Console.Print("Right")
	} else {
		Console.Print("Wrong")
	}
}

//На вход подается строка. Нужно определить, является ли она
//палиндромом. Если строка является палиндромом - вывести Нет
//Все входные строчки в нижнем регистре.

func DeterminePalindromeOrNot() {
	var (
		mainString string
		comparisonVariable string
	)	
	
	Console.Print("Enter your word: ")
	Console.Scan(&mainString)

	runeString := []rune(mainString)

	for comprasionIndex := len(runeString) - 1;  comprasionIndex >= 0; comprasionIndex-- {
		comparisonVariable += string(runeString[comprasionIndex])
	}

	if mainString == comparisonVariable {
		Console.Println("Palindrome")
	} else {
		Console.Println("No")
	}
}

//Даются две строки X и S. Нужно найти и вывести первое вхождение
//подстроки S в строке X. Если подстроки S нет в строке X - вывести -1.

func DetermineIfStringIsInAnotherString() {
	var (
		subString  string
		mainString string
	)
	Console.Print("Enter main string: ")
	Console.Scan(&mainString)

	Console.Print("Enter substring: ")
	Console.Scan(&subString)

	//Ниже можно было не писать конструкцию с If, так как при использовании
	//метода Index в хорошем случае выводится индекс из строки, при котором
	//встречается вторая подстрока, а в плохом -1, как и надо было в
	//условии задачи, но я написал, чтобы запомнить об этом методе.

	if strings.Contains(mainString, subString) == true {
		Console.Println(strings.Index(mainString, subString))
	} else {
		Console.Print("1")
	}

	//strings.Contains(mainString, subString) == true - показывает содержится ли
	//подстрока в данной строке .
	//Console.Println(strings.Index(mainString, subString)) - выводит позицию,
	//с которой начинается первое совпадение.
}

//На вход дается строка, из нее нужно сделать другую строку,
//оставив только нечетные символы (считая с нуля).

func RemoveEvenCharactersInString() {
	var mainString string
	var resultString string

	Console.Print("Enter your string: ")
	Console.Scan(&mainString)

	runeMainString := []rune(mainString)

	for searchIndex := 0; searchIndex < len(runeMainString); searchIndex++ {
		if searchIndex%2 == 0 {
			resultString += string(runeMainString[searchIndex])
		}
	}

	Console.Print("Result: ", resultString)
}

//Дается строка. Нужно удалить все символы, которые встречаются
//более одного раза и вывести получившуюся строку.

func RemoveCharactersThatOccurMoreThanOnce() {
	var (
		mainString   string
		resultString string
	)

	Console.Print("Enter your string ")
	Console.Scan(&mainString)

	runeString := []rune(mainString)

	for searchIndex := 0; searchIndex < len(runeString); searchIndex++ {
		if strings.Count(mainString, string(runeString[searchIndex])) == 1 {
			resultString += string(runeString[searchIndex])
		}
	}

	Console.Print(resultString)
}

/*Ваша задача сделать проверку подходит ли пароль вводимый пользователем
под заданные требования. Длина пароля должна быть не менее 5 символов, он
должен содержать только арабские цифры и буквы латинского алфавита. На
вход подается строка-пароль. Если пароль соответствует требованиям -
вывести "Ok", иначе вывести "Wrong password"*/

func DetermineIfThePasswordIsCorrect() {
	var (
		passwordString string
		right bool = true
		alphabet string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	)

	Console.Print("Enter your password: ")
	Console.Scan(&passwordString)
		
	passwordRune := []rune (passwordString)

	if len(passwordRune) < 5 {
		right = false
	}

	for checkIndex := 0; checkIndex < len(passwordRune); checkIndex++ {
		if right == false {
			break
		}
		
		if !strings.Contains(alphabet, string(passwordRune[checkIndex])) {
			right = false
			break
		}    
		
		if unicode.IsSpace(passwordRune[checkIndex]) {
			right = false
			break
		}
		
		if !unicode.IsLetter(passwordRune[checkIndex]) && !unicode.IsDigit(passwordRune[checkIndex]) {
			right = false
			break
		}
}

		if right {
			Console.Print("Ok")
		} else {
			Console.Print("Wrong password")    
}    
}
