package ingredients

import (
	"fmt"
	"math/rand"
	"strings"
)

type MainIngredient struct {
	name     string
	quantity string
	rand     *rand.Rand

	Gender   string
	Multiple bool
}

func NewMainIngredient(name, gender string, multiple bool, rnd *rand.Rand) MainIngredient {
	ingredient := MainIngredient{
		name:     name,
		Gender:   gender,
		Multiple: multiple,
		rand:     rnd,
	}

	var words []string

	switch i := rnd.Intn(3); i {
	case 0, 1:
		var value int
		var unit string
		switch i {
		case 0:
			value = (rnd.Intn(50) + 1) * 10
			if value == 1 {
				unit = "gramme"
			} else {
				unit = "grammes"
			}
			break
		case 1:
			value = rnd.Intn(6) + 2
			if value == 1 {
				unit = "tranche"
			} else {
				unit = "tranches"
			}
			break
		}

		words = append(words, fmt.Sprintf("%d", value), unit)

		if beginsWithVoyel(ingredient.name) {
			words = append(words, "d'")
		} else {
			words = append(words, "de ")
		}
		ingredient.quantity = strings.Join(words, " ")
		break
	case 2:
		options := []string{}

		if ingredient.Gender == "male" && !ingredient.Multiple {
			options = append(options, "un bon gros ")
			options = append(options, "un assez gros ")
			options = append(options, "un plutôt gros ")
			options = append(options, "un relativement gros ")
			options = append(options, "du ")
			options = append(options, "un moyen ")
			options = append(options, "un demi ")
			options = append(options, "un petit ")
			options = append(options, "un gros ")
		}
		if ingredient.Gender == "female" && !ingredient.Multiple {
			options = append(options, "une bonne grosse ")
			options = append(options, "une assez grosse ")
			options = append(options, "une plutôt grosse ")
			options = append(options, "une relativement grosse ")
			options = append(options, "de la ")
			options = append(options, "une moyenne ")
			options = append(options, "une petite ")
			options = append(options, "une grosse ")
		}
		if ingredient.Gender == "male" && ingredient.Multiple {
			options = append(options, "plusieurs gros ")
			options = append(options, "quelques gros ")
			options = append(options, "quelques petites ")
		}
		if ingredient.Gender == "female" && ingredient.Multiple {
			options = append(options, "plusieurs grosses ")
			options = append(options, "quelques grosses ")
			options = append(options, "quelques petites ")
		}
		if ingredient.Multiple {
			options = append(options, "quelques ")
			options = append(options, "plusieurs ")
			options = append(options, "des ")
		}

		beginnings := []string{
			"une quantité suffisante",
			"pas mal",
			"quelques morceaux",
			"un bon paquet",
			"beaucoup",
			"un peu",
			"un tout petit peu",
			"beaucoup",
		}
		for _, beginning := range beginnings {
			if beginsWithVoyel(ingredient.name) {
				options = append(options, fmt.Sprintf("%s d'", beginning))
			} else {
				options = append(options, fmt.Sprintf("%s de ", beginning))
			}
		}

		if len(options) > 0 {
			ingredient.quantity = options[rand.Intn(len(options))]
		}

		break
	}

	return ingredient
}

func (i MainIngredient) nameWithPrefix() string {
	switch {
	case i.Multiple:
		return fmt.Sprintf("les %s", i.name)
	case beginsWithVoyel(i.name):
		return fmt.Sprintf("l'%s", i.name)
	case i.Gender == "male":
		return fmt.Sprintf("le %s", i.name)
	case i.Gender == "female":
		return fmt.Sprintf("la %s", i.name)
	}
	return ""
}

func (i MainIngredient) GetSteps() Steps {
	steps := make(Steps, 0)

	availableStartSteps := Steps{
		Step{
			Instruction: fmt.Sprintf("découpez %s en fines petites tranches", i.nameWithPrefix()),
			Weight:      -100,
		},
	}

	availableFinishSteps := Steps{
		Step{
			Instruction: fmt.Sprintf("déposez %s juste au dessus", i.nameWithPrefix()),
			Weight:      100,
		},
	}

	steps = append(steps, availableStartSteps[i.rand.Intn(len(availableStartSteps))])

	if i.rand.Intn(10) > 1 {
		steps = append(steps, availableFinishSteps[i.rand.Intn(len(availableFinishSteps))])
	}

	return steps
}

func (i MainIngredient) IsMultiple() bool  { return i.Multiple }
func (i MainIngredient) GetGender() string { return i.Gender }

func (i MainIngredient) TitlePart(left Ingredient) string {
	// fixme: get a random possibility not the first one that trigger
	if left == nil {
		return i.name
	}

	switch i.rand.Intn(2) {
	case 0:
		switch {
		case i.Multiple:
			return fmt.Sprintf("aux %s", i.name)
		case beginsWithVoyel(i.name):
			return fmt.Sprintf("à l'%s", i.name)
		case i.Gender == "male":
			return fmt.Sprintf("au %s", i.name)
		case i.Gender == "female":
			return fmt.Sprintf("à la %s", i.name)
		}
	case 1:
		var suffix string
		if beginsWithVoyel(i.name) {
			suffix = "d'"
		} else {
			suffix = "de "
		}

		switch {
		case left.GetGender() == "male" && !left.IsMultiple():
			return fmt.Sprintf("assorti %s%s", suffix, i.name)
		case left.GetGender() == "female" && !left.IsMultiple():
			return fmt.Sprintf("assortie %s%s", suffix, i.name)
		case left.GetGender() == "male" && left.IsMultiple():
			return fmt.Sprintf("assortis %s%s", suffix, i.name)
		case left.GetGender() == "female" && left.IsMultiple():
			return fmt.Sprintf("assorties %s%s", suffix, i.name)
		}
	}
	panic("should not happen")
}

func (i MainIngredient) Kind() string { return "main-ingredient" }
func (i MainIngredient) Name() string { return i.name }
func (i MainIngredient) NameAndQuantity() string {
	return fmt.Sprintf("%s%s", i.quantity, i.name)
}

func (i MainIngredient) ToMap() map[string]interface{} {
	ret := make(map[string]interface{}, 0)
	ret["name"] = i.name
	ret["kind"] = i.Kind()
	ret["name-and-quantity"] = i.NameAndQuantity()
	ret["quantity"] = i.quantity
	ret["is-multiple"] = i.Multiple
	ret["gender"] = i.Gender
	return ret
}
