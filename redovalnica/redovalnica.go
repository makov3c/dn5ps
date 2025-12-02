package redovalnica

import (
	"fmt"
)

type Student struct {
	Ime     string
	Priimek string
	Ocene   []int
}

var NajmanjOcen int = 1;
var MinOcena int = 1;
var MaxOcena int = 10;

func DodajOceno(studenti map[string]Student, vpisnaStevilka string, ocena int) {
	st, obstaja := studenti[vpisnaStevilka]
	if !obstaja {
		fmt.Println("Napaka: študent z vpisno številko", vpisnaStevilka, "ne obstaja.")
		return
	}

	if ocena < MinOcena || ocena > MaxOcena {
		fmt.Println("Napaka: ocena mora biti med MinOcena in MaxOcena.")
		return
	}

	st.Ocene = append(st.Ocene, ocena)
	studenti[vpisnaStevilka] = st
}

func povprecje(studenti map[string]Student, vpisnaStevilka string) float64 {
	st, obstaja := studenti[vpisnaStevilka]
	if !obstaja {
		return -1.0
	}

	if len(st.Ocene) < NajmanjOcen {
		fmt.Println("Napaka: za računanje povprečja študent potrebuje vsaj eno oceno")
		return 0
	}

	var vsota int
	for _, ocena := range st.Ocene {
		vsota += ocena
	}
	return float64(vsota) / float64(len(st.Ocene))
}

func IzpisVsehOcen(studenti map[string]Student) {
	fmt.Println("REDOVALNICA:")
	for vpisna, st := range studenti {
		fmt.Printf("%s - %s %s: %v\n", vpisna, st.Ime, st.Priimek, st.Ocene)
	}
}
func IzpisiKoncniUspeh(studenti map[string]Student) {
	for vpisna, student := range studenti {
		p := povprecje(studenti, vpisna)
		fmt.Printf("%s %s: povprečna ocena %.1f -> ", student.Ime, student.Priimek, p)

		if p >= 9 {
			fmt.Println("Odličen študent!")
		} else if p >= 6 {
			fmt.Println("Povprečen študent")
		} else {
			fmt.Println("Neuspešen študent")
		}
	}
}

