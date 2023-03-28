package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {

	//strings
	fmt.Println(strings.Contains("computador", "dor"))  //se contem a menor string na maior
	fmt.Println(strings.Count("computador", "o"))       //quantas vezes a string menor é usada na maior
	fmt.Println(strings.HasPrefix("computador", "com")) // se a string maior comeca com a string menor

	//io
	if _, err := io.WriteString(os.Stdout, "Hello World"); err != nil {
		log.Fatal(err)
	}

	//Criando arquivo
	message := []byte("Hello, Gophers!")
	err := ioutil.WriteFile("Teste", message, 0644)
	if err != nil {
		log.Fatal(err)
	}

	//Bytes
	fmt.Printf("%s", bytes.Title([]byte("Hello, Go!")))

	//os
	f, err := os.OpenFile("AulaGo.txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}

	//path/filepath
	/*filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		return nil
	}) */

	//error
	//arquivo error.go

	//

}
