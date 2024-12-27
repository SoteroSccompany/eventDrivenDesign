package main

import "fmt"

func main() {
	evento := []string{"test", "test2", "test3", "test4"}
	// evento = evento[2:] //Pula os dois primeiros elementos e mostra os outros
	// evento = evento[:0] //Pula todos os elementos e mostra um array vazio
	// evento = evento[1:] //Pula os dois primeiros elementos e mostra os outros
	evento = append(evento[:2], evento[3:]...) //Aqui remove o 2 evento e mantem os outros

	fmt.Println(evento)
}
