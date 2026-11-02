package main

import (
	"fmt"
	"os"

	pushswap "push-swap"
)

func main() {
	if len(os.Args) < 2 {
		return
	}
	nombres, err := pushswap.LireNombres(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error")
		os.Exit(1)
	}
	for _, instruction := range pushswap.Trier(nombres) {
		fmt.Println(instruction)
	}
}
