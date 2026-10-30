package pushswap

import "fmt"

type Piles struct {
	A            []int
	B            []int
	Instructions []string
}

func NouvellesPiles(nombres []int) *Piles {
	copie := append([]int{}, nombres...)
	return &Piles{A: copie}
}

func pousser(source []int, destination []int) ([]int, []int) {
	if len(source) == 0 {
		return source, destination
	}
	premier := source[0]
	source = source[1:]
	destination = append([]int{premier}, destination...)
	return source, destination
}

func echanger(pile []int) {
	if len(pile) >= 2 {
		pile[0], pile[1] = pile[1], pile[0]
	}
}

func rotation(pile []int) []int {
	if len(pile) < 2 {
		return pile
	}
	resultat := append([]int{}, pile[1:]...)
	resultat = append(resultat, pile[0])
	return resultat
}

func rotationInverse(pile []int) []int {
	if len(pile) < 2 {
		return pile
	}
	dernier := pile[len(pile)-1]
	resultat := []int{dernier}
	resultat = append(resultat, pile[:len(pile)-1]...)
	return resultat
}

func (p *Piles) Executer(instruction string) error {
	switch instruction {
	case "pa":
		p.B, p.A = pousser(p.B, p.A)
	case "pb":
		p.A, p.B = pousser(p.A, p.B)
	case "sa":
		echanger(p.A)
	case "sb":
		echanger(p.B)
	case "ss":
		echanger(p.A)
		echanger(p.B)
	case "ra":
		p.A = rotation(p.A)
	case "rb":
		p.B = rotation(p.B)
	case "rr":
		p.A = rotation(p.A)
		p.B = rotation(p.B)
	case "rra":
		p.A = rotationInverse(p.A)
	case "rrb":
		p.B = rotationInverse(p.B)
	case "rrr":
		p.A = rotationInverse(p.A)
		p.B = rotationInverse(p.B)
	default:
		return fmt.Errorf("instruction inconnue : %q", instruction)
	}
	return nil
}

func (p *Piles) Faire(instruction string) {
	p.Executer(instruction)
	p.Instructions = append(p.Instructions, instruction)
}

func EstTriee(liste []int) bool {
	for i := 1; i < len(liste); i++ {
		if liste[i-1] > liste[i] {
			return false
		}
	}
	return true
}

func (p *Piles) ToutEstTrie() bool {
	return len(p.B) == 0 && EstTriee(p.A)
}
