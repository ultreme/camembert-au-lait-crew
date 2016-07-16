package calcrand

import "time"

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
