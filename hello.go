package main

import "fmt"

var x = "Hello, world!"

func main() {

	//Hello World
	fmt.Println("Hello World")

	//numeros
	fmt.Println("2.5 + 2.5 = ", 2.5+2.5)

	//string
	fmt.Println(len("Hello World"))
	fmt.Println("Hello World"[2])
	fmt.Println("Hello" + "World")

	//boolean
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)

	//inferencia de tipos
	var nome = "Mayara"
	var idade = 27
	var versao = 1.1
	fmt.Println("Olá sr(a).", nome, "sua idade é", idade)
	fmt.Println("Este programa está na versão ", versao)

	//variavel no escopo
	fmt.Println(x)

	//array
	var xArray [5]int
	xArray[4] = 80
	fmt.Println(xArray)

	fmt.Println()

	//For -If - Switch
	for i := 1; i <= 5; i++ {
		switch i {
		case 1:
			fmt.Print("Um")
		case 2:
			fmt.Print("Dois")
		case 3:
			fmt.Print("Tres")
		case 4:
			fmt.Print("Quatro")
		case 5:
			fmt.Print("Cinco")

		}

		if i%2 == 0 {
			fmt.Println(" é par")
		} else {
			fmt.Println(" é impar")
		}
	}

	//Fatia
	fatia := make([]float64, 5)
	fmt.Println(fatia)

	arr := [7]float64{1, 2, 3, 4, 5, 6, 7}
	fatia2 := arr[1:3]
	fmt.Println(fatia2)

	fatia3 := append(fatia2, 4, 5)
	fmt.Println(fatia3)

	fatia4 := make([]float64, 2)
	//copy(fatia4, arr)

	fatia4 = arr[0:1]
	fmt.Println(fatia4)
}
