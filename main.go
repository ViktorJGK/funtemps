package main

import (
	"flag"
	"fmt"
	"github.com/ViktorJGK/funtemps/conv"
)

// Definerer flag-variablene i hoved-"scope"
var fahr float64
var out string
var funfacts string

// Bruker init (som anbefalt i dokumentasjonen) for å sikre at flagvariablene
// er initialisert.
func init() {

	/*
	   Her er eksempler på hvordan man implementerer parsing av flagg.
	   For eksempel, kommando
	       funtemps -F 0 -out C
	   skal returnere output: 0°F er -17.78°C
	*/

	flag.Float64Var(&fahr, "F", 0.0, "temperatur i grader fahrenheit")
	flag.Float64Var(&fahr, "C", 0.0, "temperatur i grader celsius")
	flag.Float64Var(&fahr, "K", 0.0, "temperatur i grader kelvin")
	//
	flag.StringVar(&out, "out", "C", "beregne temperatur i C - celsius, F -farhenheit, K - Kelvin")
	//	flag.StringVar(&funfacts, "funfacts", "sun", "\"fun-facts\") om sun - Solen, luna - Månen og terra - Jorden")
	// Du må selv definere flag-variabelen for -t flagget, som bestemmer
	// hvilken temperaturskala skal brukes når funfacts skal vises
}

func main() {
	flag.Parse()
	switch out {
	case "C":
		result := conv.FahrenheitToCelsius(fahr)
		fmt.Printf("%v°F er %v°C\n", fahr, result)
	case "K":
		result := conv.FahrenheitToKelvin(fahr)
		fmt.Printf("%v°F er %vK\n", fahr, result)
	default:
		fmt.Printf("%v°F\n", fahr)
	}

	switch out {
	case "F":
		result := conv.CelsiusToFahrenheit(fahr)
		fmt.Printf("%v°C er %vF\n", fahr, result)
	case "K":
		result := conv.CelsiusToKelvin(fahr)
		fmt.Printf("%v°C er %vK\n", fahr, result)
	default:
		fmt.Printf("%v°C\n", fahr)
	}

	switch out {
	case "C":
		result := conv.KelvinToCelsius(fahr)
		fmt.Printf("%v°K er %vC\n", fahr, result)
	case "F":
		result := conv.KelvinToFahrenheit(fahr)
		fmt.Printf("%v°K er %vF\n", fahr, result)
	default:
		fmt.Printf("%v°K\n", fahr)

	}

	/**
	    Her må logikken for flaggene og kall til funksjoner fra conv og funfacts
	    pakkene implementeres.
	    Det er anbefalt å sette opp en tabell med alle mulige kombinasjoner
	    av flagg. flag-pakken har funksjoner som man kan bruke for å teste
	    hvor mange flagg og argumenter er spesifisert på kommandolinje.
	        fmt.Println("len(flag.Args())", len(flag.Args()))
			    fmt.Println("flag.NFlag()", flag.NFlag())
	    Enkelte kombinasjoner skal ikke være gyldige og da må kontrollstrukturer
	    brukes for å utelukke ugyldige kombinasjoner:
	    -F, -C, -K kan ikke brukes samtidig
	    disse tre kan brukes med -out, men ikke med -funfacts
	    -funfacts kan brukes kun med -t
	    ...
	    Jobb deg gjennom alle tilfellene. Vær obs på at det er en del sjekk
	    implementert i flag-pakken og at den vil skrive ut "Usage" med
	    beskrivelsene av flagg-variablene, som angitt i parameter fire til
	    funksjonene Float64Var og StringVar
	*/

	// Her er noen eksempler du kan bruke i den manuelle testingen
	// fmt.Println(fahr, out, funfacts)

	//fmt.Println("len(flag.Args())", len(flag.Args()))
	//fmt.Println("flag.NFlag()", flag.NFlag())

	//	fmt.Println(isFlagPassed("out"))

	// Eksempel på enkel logikk
	if out == "C" && isFlagPassed("F") {
		conv.FahrenheitToCelsius(fahr)

		// Kalle opp funksjonen FahrenheitToCelsius(fahr), som da
		// skal returnere °C
		fmt.Println("0°F er -17.78°C")
	}
}

// Funksjonen sjekker om flagget er spesifisert på kommandolinje
// Du trenger ikke å bruke den, men den kan hjelpe med logikken
func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}
