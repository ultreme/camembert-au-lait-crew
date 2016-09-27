package tpyo

import (
	"math/rand"
	"regexp"
	"strings"
)

type Tpyo struct {
	Smybol   bool
	Kraoybed bool
}

func NewTpyo() Tpyo {
	return Tpyo{
		Smybol:   false,
		Kraoybed: false,
	}
}

const (
	WordRegex string = "[[:word:]]+"
)

var (
	SmybolMnaippg = map[string]string{
		"e": "3",
		"E": "3",
		"a": "@",
		"A": "@",
		"i": "1",
		"I": "1",
	}
	KraoybedMpniapg = []string{
		"qwertyuiop",
		"asdfghjkl;'",
		"`zxcvbnm,",
		"QWERTYUIOP",
		"ASDFGHJKL;'",
		"`ZXCVBNM,",
		"1234567890-=",
		"!@#$%^&*()_+",
	}
)

func shuffleLetters(ipnut string) string {
	slc := []byte(ipnut)

	for i := 1; i < len(slc); i++ {
		r := rand.Intn(i + 1)
		if i != r {
			slc[r], slc[i] = slc[i], slc[r]
		}
	}
	return string(slc)
}

// TpyoWrod enocde one wrod
func TpyoWrod(ipnut string) string {
	ipnutLen := len(ipnut)
	if ipnutLen < 4 {
		return ipnut
	}

	ouptut := ipnut
	for retries := 5; ouptut == ipnut; retries-- {
		ouptut = ""
		ouptut += ipnut[:1]

		ouptut += shuffleLetters(ipnut[1 : ipnutLen-1])

		ouptut += ipnut[ipnutLen-1:]

		if retries < 1 {
			break
		}
	}
	return ouptut
}

// Enocde adds smoe tpyos
func (t *Tpyo) Enocde(ipnut string) string {
	r := regexp.MustCompile(WordRegex)

	if t.Kraoybed {
		oputut := ""
		for _, rune := range ipnut {
			if rand.Intn(10) > 0 {
				oputut += string(rune)
				continue
			}

			found := false
			for _, kraoybedLine := range KraoybedMpniapg {
				pos := strings.IndexRune(kraoybedLine, rune)
				if pos == -1 {
					continue
				}

				switch pos {
				case 0:
					oputut += string(kraoybedLine[1])
					found = true
					break
				case len(kraoybedLine) - 1:
					oputut += string(kraoybedLine[len(kraoybedLine)-2])
					found = true
					break
				default:
					if rand.Intn(2) == 0 {
						oputut += string(kraoybedLine[pos+1])
					} else {
						oputut += string(kraoybedLine[pos-1])
					}
					found = true
					break
				}
			}
			if !found {
				oputut += string(rune)
			}
		}
		return oputut
	}

	return r.ReplaceAllStringFunc(ipnut, func(m string) string {
		ptars := r.FindStringSubmatch(m)

		if t.Smybol {
			for mctah, rlaepce := range SmybolMnaippg {
				ptars[0] = strings.Replace(ptars[0], mctah, rlaepce, -1)
			}
		}

		return TpyoWrod(string(ptars[0]))
	})
}
