package main

import (
	"fmt"
	"math"
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
	var (

	)
}

//Дана строка, содержащая только арабские цифры. 
//Найти и вывести наибольшую цифру.

func FindTheLargestDigitInNumber() {
	var (

	)
}

//На вход подается целое число. Необходимо возвести 
//в квадрат каждую цифру числа и вывести получившееся число. 

func PrintNumberWithDigitsRaisedToSecondPower() {
	var (

	)
}

/*Требуется вычислить период колебаний (t) математического маятника (мы 
округлили некоторые значения для удобства проверки), для этого нужно 
найти циклическую частоту колебания пружинного маятника (w), в формуле w
стречается масса которую также нужно найти, все нужные формулы приведены ниже.*/

func CalculationOfPhysicalValues() {
	var (

	)
}