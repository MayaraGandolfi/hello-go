package main

import (
	"fmt"
	"time"
)

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
	}
}

func pingar(c chan string) { // chan palavra usada para canal
	for i := 0; ; i++ {
		c <- "ping" // usado para enviar e receber mensagem pelo canal
	}
}

func imprimir(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)

	}
}

func main() {
	//Goroutines
	go f(0)
	var escreva string
	fmt.Scanln(&escreva)
	fmt.Println(escreva)

	//canais
	var c chan string = make(chan string)

	go pingar(c)
	go imprimir(c)

	var entrada string
	fmt.Scanln(&entrada)

	// Select funciona como um Switch para canais.
	// Select permite que você aguarde operações de vários canais.
	// Combinar goroutines e canais com select é um recurso poderoso do Go.
	// Para nosso exemplo, selecionaremos em dois canais.
	//Cada canal receberá um valor após algum tempo, para simular, por exemplo,
	//o bloqueio de operações RPC em execução em goroutines concorrentes.

	c1 := make(chan string)
	c2 := make(chan string)

	go func() { //Recebemos os valores "um"e depois "dois"conforme o esperado.
		time.Sleep(1 * time.Second) //Observe que o tempo total de execução é de apenas
		//~ 2 segundos,
		//pois o 1 e o 2 segundos são Sleeps executados simultaneamente.
		c1 <- "um"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "dois"
	}()

	for i := 0; i < 2; i++ {
		select { //Usaremos select para aguardar esses
		//dois valores simultaneamente, imprimindo cada um à medida que chegam.
		case msg1 := <-c1:
			fmt.Println("receba", msg1)
		case msg2 := <-c2:
			fmt.Println("receba", msg2)
		}
	}

}
