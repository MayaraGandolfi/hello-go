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

}
