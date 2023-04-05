package main

import "fmt"

// 2.4 Структуры (Structure)

/*Вам необходимо реализовать структуру со свойствами-полями On,
Ammo и Power, с типами bool, int, int соответственно. У этой 
структуры должны быть методы: Shoot и RideBike, которые не 
принимают аргументов, но возвращают значение bool.*/

type MainStruct struct {
	On    bool
	Ammo  int
	Power int
}

func (m *MainStruct) Shoot() bool {
	flag := false
	
	if m.On == true && m.Ammo > 0 {
			flag = true
			m.Ammo--
			return flag
	} else {
			return flag       
	}  
}

func (m *MainStruct)  RideBike() bool {
	flag := false
	
	if m.On == true && m.Power > 0{
			flag = true
			m.Power--
			return flag
	} else {
			return flag    
	}
}

func FunctionWithStructure() {
	testStruct := &MainStruct{true, 5,5}
	fmt.Print(testStruct)
}

