package main

import (
	"fmt"
	"math"
)

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

	//Mapa
	mapa := make(map[string]int)
	mapa["chave"] = 10
	fmt.Println(mapa["chave"])
	mapa["casa"] = 5
	fmt.Println(mapa)

	elemento := make(map[string]string)
	elemento["H"] = "Hidrogenio"
	elemento["He"] = "Helio"
	elemento["Li"] = "Litio"
	fmt.Println(elemento["Li"])

	//Estrutura
	fmt.Println(pessoa{"Maria", 54})
	fmt.Println(pessoa{"João", 25})

	r := retangulo{altura: 25, comprimento: 50}
	fmt.Println("Retangulo r comprimento e altura = ", r)

	//Metodo
	fmt.Println("Retangulo r area ", r.area())
	fmt.Println("Retangulo r perimetro ", r.perimetro())

	//Interfaces
	q := quadrado{lado: 25}
	c := circulo{raio: 100}

	medir(q)
	medir(c)

	medir(r)

}

// Estrutura
type pessoa struct {
	nome  string
	idade float64
}

type retangulo struct {
	comprimento, altura float64
}

type quadrado struct {
	lado float64
}

type circulo struct {
	raio float64
}

// Metodo
// variavel tipo - nome metodo - tipo return
func (r retangulo) perimetro() float64 {
	return 2*r.altura + 2*r.comprimento
}

func (r retangulo) area() float64 {
	return r.comprimento * r.altura
}

func (q quadrado) perimetro() float64 {
	return q.lado * 4
}

func (q quadrado) area() float64 {
	return q.lado * q.lado
}

func (c circulo) perimetro() float64 {
	return 2 * math.Pi * c.raio
}

func (c circulo) area() float64 {
	return math.Pi * c.raio * c.raio
}

// Interface
type geometria interface {
	area() float64
	perimetro() float64
}

func medir(g geometria) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perimetro())
}
