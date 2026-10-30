package pushswap

import (
	"fmt"
	"strconv"
	"strings"
)

func LireNombres(arguments []string) ([]int, error) {
	nombres := []int{}
	for _, argument := range arguments {
		for _, mot := range strings.Fields(argument) {
			nombre, err := strconv.Atoi(mot)
			if err != nil {
				return nil, fmt.Errorf("pas un entier : %q", mot)
			}
			for _, existant := range nombres {
				if existant == nombre {
					return nil, fmt.Errorf("doublon : %d", nombre)
				}
			}
			nombres = append(nombres, nombre)
		}
	}
	return nombres, nil
}
