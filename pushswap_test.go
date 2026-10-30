package pushswap

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestLireNombres(t *testing.T) {
	nombres, err := LireNombres([]string{"2 1 3", "6", "-5"})
	if err != nil {
		t.Fatal(err)
	}
	attendu := []int{2, 1, 3, 6, -5}
	if !reflect.DeepEqual(nombres, attendu) {
		t.Fatalf("obtenu %v, attendu %v", nombres, attendu)
	}
}

func TestLireNombresErreurs(t *testing.T) {
	_, err := LireNombres([]string{"0 one 2 3"})
	if err == nil {
		t.Fatal("un mot qui n'est pas un entier doit donner une erreur")
	}
	_, err = LireNombres([]string{"1 2 1"})
	if err == nil {
		t.Fatal("un doublon doit donner une erreur")
	}
}

func verifierPiles(t *testing.T, piles *Piles, a []int, b []int) {
	t.Helper()
	if len(piles.A) != len(a) || (len(a) > 0 && !reflect.DeepEqual(piles.A, a)) {
		t.Fatalf("pile a = %v, attendue %v", piles.A, a)
	}
	if len(piles.B) != len(b) || (len(b) > 0 && !reflect.DeepEqual(piles.B, b)) {
		t.Fatalf("pile b = %v, attendue %v", piles.B, b)
	}
}

func executerTout(t *testing.T, piles *Piles, instructions ...string) {
	t.Helper()
	for _, instruction := range instructions {
		if err := piles.Executer(instruction); err != nil {
			t.Fatal(err)
		}
	}
}

func TestExempleDuSujet(t *testing.T) {
	piles := NouvellesPiles([]int{2, 1, 3, 6, 8, 5})

	executerTout(t, piles, "sa")
	verifierPiles(t, piles, []int{1, 2, 3, 6, 8, 5}, nil)

	executerTout(t, piles, "pb", "pb", "pb")
	verifierPiles(t, piles, []int{6, 8, 5}, []int{3, 2, 1})

	executerTout(t, piles, "rb")
	verifierPiles(t, piles, []int{6, 8, 5}, []int{2, 1, 3})

	executerTout(t, piles, "rra", "rrb")
	verifierPiles(t, piles, []int{5, 6, 8}, []int{3, 2, 1})

	executerTout(t, piles, "pa", "pa", "pa")
	verifierPiles(t, piles, []int{1, 2, 3, 5, 6, 8}, nil)

	if !piles.ToutEstTrie() {
		t.Fatal("la pile a devrait etre triee et la pile b vide")
	}
}

func TestInstructionsDoubles(t *testing.T) {
	piles := NouvellesPiles([]int{1, 2, 3})
	piles.B = []int{4, 5, 6}

	executerTout(t, piles, "ss")
	verifierPiles(t, piles, []int{2, 1, 3}, []int{5, 4, 6})

	executerTout(t, piles, "rr")
	verifierPiles(t, piles, []int{1, 3, 2}, []int{4, 6, 5})

	executerTout(t, piles, "rrr")
	verifierPiles(t, piles, []int{2, 1, 3}, []int{5, 4, 6})
}

func TestPilesVides(t *testing.T) {
	piles := NouvellesPiles(nil)
	executerTout(t, piles, "pa", "pb", "sa", "sb", "ss", "ra", "rb", "rr", "rra", "rrb", "rrr")
	verifierPiles(t, piles, nil, nil)
}

func TestInstructionInconnue(t *testing.T) {
	piles := NouvellesPiles([]int{1, 2})
	for _, instruction := range []string{"", "xx", "SA", "sa ", "pc"} {
		if piles.Executer(instruction) == nil {
			t.Fatalf("%q devrait etre refusee", instruction)
		}
	}
}

func verifierTri(t *testing.T, nombres []int) []string {
	t.Helper()
	copie := append([]int{}, nombres...)
	instructions := Trier(nombres)
	if !reflect.DeepEqual(nombres, copie) {
		t.Fatalf("Trier a modifie la liste de depart %v", copie)
	}
	piles := NouvellesPiles(nombres)
	for _, instruction := range instructions {
		if err := piles.Executer(instruction); err != nil {
			t.Fatal(err)
		}
	}
	if !piles.ToutEstTrie() {
		t.Fatalf("liste %v mal triee, resultat a=%v b=%v", copie, piles.A, piles.B)
	}
	return instructions
}

func toutesLesPermutations(taille int) [][]int {
	resultat := [][]int{}
	var construire func(courant []int, utilise []bool)
	construire = func(courant []int, utilise []bool) {
		if len(courant) == taille {
			resultat = append(resultat, append([]int{}, courant...))
			return
		}
		for i := 0; i < taille; i++ {
			if utilise[i] {
				continue
			}
			utilise[i] = true
			construire(append(courant, i*3-4), utilise)
			utilise[i] = false
		}
	}
	construire([]int{}, make([]bool, taille))
	return resultat
}

func TestToutesLesPermutationsPetites(t *testing.T) {
	maximum := map[int]int{2: 1, 3: 2, 4: 8, 5: 12}
	for taille := 0; taille <= 7; taille++ {
		for _, liste := range toutesLesPermutations(taille) {
			instructions := verifierTri(t, liste)
			limite, existe := maximum[taille]
			if existe && len(instructions) > limite {
				t.Fatalf("liste %v : %d instructions (max %d)", liste, len(instructions), limite)
			}
		}
	}
}

func TestDejaTriee(t *testing.T) {
	if len(Trier([]int{1, 2, 3, 4, 5, 6, 7, 8})) != 0 {
		t.Fatal("une liste deja triee ne doit donner aucune instruction")
	}
	if len(Trier(nil)) != 0 {
		t.Fatal("une liste vide ne doit donner aucune instruction")
	}
}

func TestListesAleatoires(t *testing.T) {
	hasard := rand.New(rand.NewSource(42))
	cas := []struct {
		taille int
		limite int
	}{
		{8, 50}, {20, 200}, {100, 1300}, {500, 7500},
	}
	for _, c := range cas {
		for essai := 0; essai < 10; essai++ {
			liste := hasard.Perm(c.taille)
			for i := range liste {
				liste[i] -= c.taille / 2
			}
			instructions := verifierTri(t, liste)
			if len(instructions) > c.limite {
				t.Fatalf("%d nombres : %d instructions (limite %d)", c.taille, len(instructions), c.limite)
			}
		}
	}
}

func TestValeursExtremes(t *testing.T) {
	verifierTri(t, []int{9223372036854775807, -9223372036854775808, 0, 5, -5, 42})
}
