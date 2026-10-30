package pushswap

func Trier(nombres []int) []string {
	if len(nombres) < 2 || EstTriee(nombres) {
		return nil
	}
	piles := NouvellesPiles(remplacerParRangs(nombres))
	taille := len(nombres)
	if taille == 2 {
		piles.Faire("sa")
	} else if taille == 3 {
		trierTrois(piles)
	} else if taille <= 5 {
		trierPetitePile(piles)
	} else {
		trierParBits(piles)
	}
	return piles.Instructions
}

func remplacerParRangs(nombres []int) []int {
	rangs := []int{}
	for _, nombre := range nombres {
		rang := 0
		for _, autre := range nombres {
			if autre < nombre {
				rang++
			}
		}
		rangs = append(rangs, rang)
	}
	return rangs
}

func trierTrois(piles *Piles) {
	a := piles.A[0]
	b := piles.A[1]
	c := piles.A[2]
	if a > b && b > c {
		piles.Faire("sa")
		piles.Faire("rra")
	} else if a > b && b < c && a < c {
		piles.Faire("sa")
	} else if a > b && b < c && a > c {
		piles.Faire("ra")
	} else if a < b && b > c && a < c {
		piles.Faire("sa")
		piles.Faire("ra")
	} else if a < b && b > c && a > c {
		piles.Faire("rra")
	}
}

func positionDuPlusPetit(liste []int) int {
	position := 0
	for i := 1; i < len(liste); i++ {
		if liste[i] < liste[position] {
			position = i
		}
	}
	return position
}

func amenerLePlusPetitEnHaut(piles *Piles) {
	position := positionDuPlusPetit(piles.A)
	taille := len(piles.A)
	if position <= taille/2 {
		for i := 0; i < position; i++ {
			piles.Faire("ra")
		}
	} else {
		for i := 0; i < taille-position; i++ {
			piles.Faire("rra")
		}
	}
}

func trierPetitePile(piles *Piles) {
	for len(piles.A) > 3 {
		amenerLePlusPetitEnHaut(piles)
		piles.Faire("pb")
	}
	trierTrois(piles)
	for len(piles.B) > 0 {
		piles.Faire("pa")
	}
}

func nombreDeBits(taille int) int {
	bits := 0
	for (taille-1)>>bits > 0 {
		bits++
	}
	return bits
}

func trierParBits(piles *Piles) {
	for bit := 0; bit < nombreDeBits(len(piles.A)); bit++ {
		taille := len(piles.A)
		for i := 0; i < taille; i++ {
			if (piles.A[0]>>bit)&1 == 0 {
				piles.Faire("pb")
			} else {
				piles.Faire("ra")
			}
		}
		for len(piles.B) > 0 {
			piles.Faire("pa")
		}
	}
}
