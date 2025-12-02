// package redovalnica enables user to add a grade, find out
// average grade of a student, print all grades of one student and print the whole redovalnica
// book aka grades of all students.
package redovalnica

import (
	"fmt"
)

// Type Student represents a student, and includes their name and surname, as well as list of grades.
type Student struct {
	Ime     string
	Priimek string
	Ocene   []int
}

// NajmanjOcen is minimum number of grades required for a student to have.
var NajmanjOcen int = 1;

// MinOcena is minimum allowed grade that a student can have.
var MinOcena int = 1;

// MaxOcena is maximum allowed grade that a student can have.
var MaxOcena int = 10;

// DodajOceno is function that adds a grade to a student wtih a certain enrollment number.
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

// povprecje is function that calculates average grade based on students enrollment number.
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

// IzpisVsehOcen is function that prints all the grades of a student whith a specific enrollment number.
func IzpisVsehOcen(studenti map[string]Student) {
	fmt.Println("REDOVALNICA:")
	for vpisna, st := range studenti {
		fmt.Printf("%s - %s %s: %v\n", vpisna, st.Ime, st.Priimek, st.Ocene)
	}
}

// IzpisiKoncniUspeh prints the final average of a studen with a certain enrollment number, along with thier name and surname.
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

