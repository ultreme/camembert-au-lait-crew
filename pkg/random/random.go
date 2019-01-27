package random // import "ultre.me/calcbiz/pkg/random"

import (
	"fmt"
	"math/rand"
	"time"
	"unicode"
)

func randomCaps(input string) string {
	output := []rune{}
	for _, c := range input {
		if rand.Intn(2) == 0 {
			output = append(output, unicode.ToUpper(c))
		} else {
			output = append(output, c)
		}
	}
	return string(output)
}

var wotds = []string{
	"une gomme", "une pomme", "phazms", "08 36 65 65 65", "le mot du jour",
	"cool", "trop de la balle", "coucou", "salut les copains",
	"un sandwich", "caribou", "jus", "du pate", "poney", "castor",
	"cancun", "hula hoop", "criterium", "lampadaire", "tabernacle",
	"anticonstitutionnellement", "grut", "i feel good", "cumulonumbus",
	"chewing gum", "jericane", "malaxer", "competant", "moineau", "wesh",
	"platane", "sycomore", "blaireau", "perudo", "azymute",
	"moissoneuse batteuse", "tracteur", "pudding", "1000 pates", "42",
	"centipede",
}

// WOTD returns the word of the day
func WOTD() string {
	return randomCaps(wotds[time.Now().YearDay()%len(wotds)])
}

var alternateLogos = []string{
	"50-bacon-50-biere-50-camembert.png",
	"dj-jean-michel-mc-pierre-gustave.jpg",
	"i-love-chips.jpg",
	"le-passe-c-etait-mieux-avant.jpg",
	"c-est-vachement-cool.jpg",
	"arc-en-le-ciel.jpg",
	"aspiration.jpg",
	"boules.jpg",
	"etoile.jpg",
	"megas-pixels.jpg",
	"petits-pixels.jpg",
	"pochette-camembert.jpg",
	"pouce.jpg",
	"trim.jpg",
	"trop-de-la-balle.jpg",
}

func AlternateLogo() string {
	file := alternateLogos[rand.Intn(len(alternateLogos))]
	return fmt.Sprintf("./static/img/logo-alternate-300/%s", file)
}

func MotDebileQuiSeMange() string {
	mots := []string{
		"beurre", "lait", "yahourt", "pain", "gruyere", "margarine",
		"curcumin", "d'épinard", "banane", "salade", "brioche",
		"sucre", "chips", "laitage", "lukum", "flotte", "chupa-chups",
		"yogourt",
	}
	return mots[rand.Intn(len(mots))]
}

func MotCool() string {
	mots := []string{
		"cool", "sympa", "gentil", "genial", "excellent", "superbe", "super",
		"vraiment tres bien", "bien", "qui en a dans le pantalon", "top",
	}
	return mots[rand.Intn(len(mots))]
}

func MotPasCool() string {
	return "pas tres " + MotCool()
}

func RandomColor(dark bool, light bool, limit int) string {
	if limit == 0 {
		limit = 80
	}

	var r, g, b int
	for good := false; !good; {
		r = rand.Intn(256)
		g = rand.Intn(256)
		b = rand.Intn(256)

		if dark {
			good = r < limit || g < limit || b < limit
		} else if light {
			good = r > 256-limit || g > 256-limit || b > 256-limit
		} else {
			good = true
		}
	}
	return fmt.Sprintf("#%X%X%X", r, g, b)
}
