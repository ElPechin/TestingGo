package main


import (
	"./mio"
	"fmt"
)

func main() {
	lista := [] int {6, 3, 1, 5, 7, 0, 2, 9, 8, 4}

	err := mio.Acumular(lista)
	if err != nil {
		fmt.Println("Ocurrio un error no esperado")
	} else {
		fmt.Println(lista[9])
	}
}
