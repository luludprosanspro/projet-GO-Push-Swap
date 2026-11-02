package main

import (
	"bufio"
	"fmt"
	"os"

	pushswap "push-swap"
)

func erreur() {
	fmt.Fprintln(os.Stderr, "Error")
	os.Exit(1)
}

func main() {
	if len(os.Args) < 2 {
		return
	}
	nombres, err := pushswap.LireNombres(os.Args[1:])
	if err != nil {
		erreur()
	}
	if len(nombres) == 0 {
		return
	}
	piles := pushswap.NouvellesPiles(nombres)
	lecteur := bufio.NewScanner(os.Stdin)
	for lecteur.Scan() {
		err := piles.Executer(lecteur.Text())
		if err != nil {
			erreur()
		}
	}
	if piles.ToutEstTrie() {
		fmt.Println("OK")
	} else {
		fmt.Println("KO")
	}
}
