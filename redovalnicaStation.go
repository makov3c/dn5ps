package main

import (
	"fmt"
	"os"
	"log"
	"context"
	"github.com/urfave/cli/v3"
	"github.com/makov3c/dn5ps/redovalnica"
)

func main() {
	cmd := &cli.Command{
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:  "stOcen",
				Usage: "najmanjše število ocen potrebnih za pozitivno oceno",
				Value: 3, // default bo najmanj 3 ocene
			},
			&cli.IntFlag{
				Name:  "minOcena",
				Usage: "najmanjša možna ocena",
				Value: 6, // default je 6
			},
			&cli.IntFlag{
				Name:  "maxOcena",
				Usage: "največja možna ocena",
				Value: 10, // default je 10
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			// timeoutSec := cmd.Int("timeout")
			// pollTime := cmd.Int("pollTime")

			redovalnica.NajmanjOcen = cmd.Int("stOcen")
			redovalnica.MinOcena = cmd.Int("minOcena")
			redovalnica.MaxOcena = cmd.Int("maxOcena")
			
			f()
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}

func f() {
	studenti := map[string]redovalnica.Student{
		"63250001": {"Ana", "Novak", []int{10, 9, 8, 9, 4}},
		"63250002": {"Boris", "Kralj", []int{6, 7, 5, 6, 1, 1, 1}},
		"63250003": {"Janez", "Novak", []int{}},
	}
	redovalnica.DodajOceno(studenti, "63250002", 8)
	redovalnica.DodajOceno(studenti, "63250003", 7)
	redovalnica.DodajOceno(studenti, "12345678", 10)
	redovalnica.DodajOceno(studenti, "63250001", 11)

	fmt.Println()
	redovalnica.IzpisVsehOcen(studenti)
	fmt.Println()
	redovalnica.IzpisiKoncniUspeh(studenti)
	fmt.Println()
}