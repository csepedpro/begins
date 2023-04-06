package main

import (
	"fmt"
	"math"
	"strings"
	"strconv"
)

// 2.7 Решение задач

func Tasks27() {
	var userChoice string

  fmt.Print("Select the task number: ")
  fmt.Scan(&userChoice)

  switch userChoice {
  case "1":
		FindHypotenuseThroughYwoLegs() 
  case "2":
    AddAnAsteriskBetweenLetters()
	case "3":
		FindTheLargestDigitInNumber()
	case "4":
		PrintNumberWithDigitsRaisedToSecondPower()
	case "5":
		CalculationOfPhysicalValues()
  default:
    fmt.Println("Introduced a non-existent variant! Please try again.")
    Tasks27()
  }
}

//На вход подаются a и b - катеты прямоугольного
//треугольника. Нужно найти длину гипотенузы

func FindHypotenuseThroughYwoLegs() {
	var (
		leg1 float64
		leg2 float64
	)

	fmt.Print("Enter two legs separated by a space: ")
	fmt.Scan(&leg1, &leg2)
	
	fmt.Print("The length of the hypotenuse is ", math.Hypot(leg1, leg2))	
}

//Дана строка, содержащая только английские буквы (большие и маленькие).
//Добавить символ ‘*’ (звездочка) между буквами (перед первой буквой и 
//после последней символ '*' добавлять не нужно)

func AddAnAsteriskBetweenLetters() {
	var mainString string
	
	fmt.Print("Enter your row: ")
	fmt.Scan(&mainString)

	resultString := strings.Replace(mainString, "", "*", -1)

	fmt.Print("Result: ", strings.Trim(resultString, string(resultString[0]))) 
}

//Дана строка, содержащая только арабские цифры. 
//Найти и вывести наибольшую цифру.

func FindTheLargestDigitInNumber() {
	var (
		mainString string
		max int = 0
	)

	fmt.Print("Enter your string: ")
	fmt.Scan(&mainString)

	runeString := []rune(mainString)

	for searchIndex := 0; searchIndex < len(runeString); searchIndex++ {
		if runeString[searchIndex] >= '0' && runeString[searchIndex] <= '9' {
			digit, _ := strconv.Atoi(string(runeString[searchIndex]))
				
			if digit > max {
				max = digit
			}	 
		} else {
				fmt.Print("Incorrect input! Please try again!")
				FindTheLargestDigitInNumber()
			}
	}

	fmt.Println("MaxNumber is ", max)
}

//На вход подается целое число. Необходимо возвести 
//в квадрат каждую цифру числа и вывести получившееся число. 

func PrintNumberWithDigitsRaisedToSecondPower() {
	var (
		mainString string
		resultString string
	)

	fmt.Print("Enter your number: ")
	fmt.Scan(&mainString)

	runeString := []rune (mainString)

	for multiplicationIndex := 0; multiplicationIndex < len(runeString); multiplicationIndex++ {
		if runeString[multiplicationIndex] >= '0' && runeString[multiplicationIndex] <= '9' {
			digit, _ := strconv.Atoi(string(runeString[multiplicationIndex]))
			squareOfDigit := digit * digit

			resultString += strconv.Itoa(squareOfDigit)   
		} else {
			fmt.Print("Incorrect input! Please tru again.")
			PrintNumberWithDigitsRaisedToSecondPower()
		}
	}

	fmt.Print(resultString)
}

/*Требуется вычислить период колебаний (t) математического маятника (мы 
округлили некоторые значения для удобства проверки), для этого нужно 
найти циклическую частоту колебания пружинного маятника (w), в формуле w
стречается масса которую также нужно найти, все нужные формулы приведены ниже.*/

var p, v, k float64 
// Глобальные переменные для вычисления формул в данном задании

func CalculationOfPhysicalValues() {

	fmt.Scan(&p, &v)

	M()
	W()
	T()
}

func M() float64 {
	m := p * v
	return m
}

func W() float64 {
	m  := M()
	return math.Sqrt(k / m)
}

func T() float64 {
	w:= W()
    t:= 6/w
	return t
}