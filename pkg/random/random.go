package random // import "ultre.me/calcbiz/pkg/random"

import (
	"fmt"
	"math/rand"
	"time"
)

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
	return wotds[time.Now().YearDay()%len(wotds)]
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
	return fmt.Sprintf("https://camembertaulaitcrew.github.io/assets/logo-alternate-300/%s", file)
}
