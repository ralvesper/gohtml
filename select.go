package main

import (
	"fmt"
	"time"

	"github.com/ralvesper/gohtml"
)

func oMaisRapido(url1, url2, url3 string) string {
	c1 := gohtml.Titulo(url1)
	c2 := gohtml.Titulo(url2)
	c3 := gohtml.Titulo(url3)

	// estrutura de controle específica para concorrência
	select {
	case t1 := <-c1:
		return t1
	case t2 := <-c2:
		return t2
	case t3 := <-c3:
		return t3
	case <-time.After(2000 * time.Millisecond):
		return "Todos perderam!"
		// default:
		// 	return "Sem resposta ainda!"
	}
}

func main() {
	campeao := oMaisRapido(
		"https://www.youtube.com",
		"https://www.amazon.com",
		"https://www.google.com",
	)
	fmt.Println(campeao)
}
