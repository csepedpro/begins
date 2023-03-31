package main

import (
	"bufio"
	Console "fmt"
	"strings"
	"unicode"
	"os"
)

// 2.5 Cтроки (string)

func Tasks25() {
	var userChoice string

  Console.Print("Select the task number: ")
  Console.Scan(&userChoice)

  switch userChoice {
  case "1":
		FunctionF()
  case "2":
    MinimumFromFour()
	case "3":
		VotingFunction()
	case "4":
		FunctionFibonaci()
	case "5":
		FunctionSumInt()
  default:
    Console.Println("Introduced a non-existent variant! Please try again.")
    Tasks21()
  }
}

//На вход подается строка. Нужно определить, является ли она правильной или нет. Правильная
//строка начинается с заглавной буквы и заканчивается точкой. Если строка правильная - 
//вывести Right иначе - вывести Wrong.

func DetermineTheCorrectWordOrNot() {
	
	Console.Print("Enter your phrase: ")
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	/* Передается в функцию bufio.NewReader, на основании которого создается
	объект bufio.Reader. Поскольку идет построчное считывание, то каждая 
	строка считывается из потока, пока не будет обнаружен символ перевода строки \n. 
	Подробнее: https://metanit.com/go/tutorial/8.9.php */



	if unicode.IsUpper(text[0]) == true && text[len(text)] == "."{
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

	)

}

//Даются две строки X и S. Нужно найти и вывести первое вхождение
//подстроки S в строке X. Если подстроки S нет в строке X - вывести -1.

func DetermineIfStringIsInAnotherString() {
	var (
		subString string
		mainString string
	)
	Console.Print("Enter main string: ")
	Console.Scan(&mainString)

	Console.Print("Enter main substring: ")
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
	var (
	
	)
}

//Дается строка. Нужно удалить все символы, которые встречаются
//более одного раза и вывести получившуюся строку.

func RemoveCharactersThatOccurMoreThanOnce() {

}

/*Ваша задача сделать проверку подходит ли пароль вводимый пользователем 
под заданные требования. Длина пароля должна быть не менее 5 символов, он
должен содержать только арабские цифры и буквы латинского алфавита. На 
вход подается строка-пароль. Если пароль соответствует требованиям -
вывести "Ok", иначе вывести "Wrong password"*/

func DetermineIfThePasswordIsCorrect() {
	var (

	)
}