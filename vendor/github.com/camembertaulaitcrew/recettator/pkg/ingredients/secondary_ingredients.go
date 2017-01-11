package ingredients

import (
	"fmt"
	"math/rand"
)

type SecondaryIngredient struct {
	name          string
	gender        string
	quantity      string
	isMultiple    bool
	isUncountable bool
	isPowder      bool
	isCitrus      bool
	isSpice       bool
	isByPiece     bool
	isSpreadable  bool
	rand          *rand.Rand
}

func (i *SecondaryIngredient) SetMethod(method Method) { panic("Not implemented") }
func (i *SecondaryIngredient) GetMethod() Method       { return nil }

func NewSecondaryIngredient(name string, gender string, isMultiple bool, rnd *rand.Rand) *SecondaryIngredient {
	ingredient := SecondaryIngredient{
		name:       name,
		gender:     gender,
		isMultiple: isMultiple,
		rand:       rnd,
		/*
			isMultiple:
			isUncountable:
			isPowder:
			isCitrus:
			isSpice:
			isByPiece:
			isSpreadable:
		*/
	}
	return &ingredient
}

func (i *SecondaryIngredient) prepare() {
	switch {
	case i.isUncountable:
		switch {
		case i.isMultiple:
			i.quantity = "des "
			break
		case beginsWithVoyel(i.name):
			i.quantity = "de l'"
			break
		case i.gender == "male":
			i.quantity = "du "
			break
		case i.gender == "female":
			i.quantity = "de la "
			break
		}
		break
	case i.isPowder:
		value := (i.rand.Intn(50) + 1) * 10
		switch {
		case value == 1 && !beginsWithVoyel(i.name):
			i.quantity = fmt.Sprintf("%d gramme de ", value)
			break
		case value == 1 && beginsWithVoyel(i.name):
			i.quantity = fmt.Sprintf("%d gramme d'", value)
			break
		case !beginsWithVoyel(i.name):
			i.quantity = fmt.Sprintf("%d grammes de ", value)
			break
		case beginsWithVoyel(i.name):
			i.quantity = fmt.Sprintf("%d grammes d'", value)
			break
		}
		break
	case i.isByPiece:
		if i.isMultiple {
			value := i.rand.Intn(20) + 2
			i.quantity = fmt.Sprintf("%d ", value)
		} else {
			i.quantity = "1 "
		}
		break
	case i.isSpice:
		suffix := "de "
		if beginsWithVoyel(i.name) {
			suffix = "d'"
		}
		options := []string{
			fmt.Sprintf("une poignée %s", suffix),
			fmt.Sprintf("une dosette %s", suffix),
			fmt.Sprintf("un verre %s", suffix),
			fmt.Sprintf("une pincée %s", suffix),
		}
		i.quantity = options[i.rand.Intn(len(options))]
		break
	case i.isSpreadable:
		suffix := "de "
		if beginsWithVoyel(i.name) {
			suffix = "d'"
		}
		options := []string{
			fmt.Sprintf("une noix %s", suffix),
			fmt.Sprintf("un morceau %s", suffix),
			fmt.Sprintf("une dose %s", suffix),
			fmt.Sprintf("une cuillère à café %s", suffix),
			fmt.Sprintf("une cuillère à soupe %s", suffix),
		}
		i.quantity = options[i.rand.Intn(len(options))]
		break
	case i.isCitrus:
		suffix := "de "
		if beginsWithVoyel(i.name) {
			suffix = "d'"
		}
		options := []string{
			fmt.Sprintf("un zeste %s", suffix),
			fmt.Sprintf("un quartier %s", suffix),
			fmt.Sprintf("une pelure %s", suffix),
			fmt.Sprintf("de la pulpe %s", suffix),
		}
		i.quantity = options[i.rand.Intn(len(options))]
		break
	}
}

func (i SecondaryIngredient) GetSteps() Steps {
	steps := make(Steps, 0)

	availableStartSteps := Steps{
		Step{
			Instruction: fmt.Sprintf("réchauffez %s à feu doux", i.nameWithPrefix()),
			Weight:      -100,
		},
		Step{
			Instruction: fmt.Sprintf("placez %s au bain-marie quelques minutes", i.nameWithPrefix()),
			Weight:      -100,
		},
		Step{
			Instruction: fmt.Sprintf("selon votre goût, vous pouvez voiler %s d'un fond de sucre", i.nameWithPrefix()),
			Weight:      -100,
		},
		Step{
			Instruction: fmt.Sprintf("ajoutez %s par dessus", i.nameWithPrefix()),
			Weight:      -100,
		},
		Step{
			Instruction: fmt.Sprintf("faites cuire %s dans un wok", i.nameWithPrefix()),
			Weight:      -100,
		},
		Step{
			Instruction: fmt.Sprintf("faites chauffer %s et penser à vanner pendant le refroidissement", i.nameWithPrefix()),
			Weight:      -100,
		},
	}

	if i.rand.Intn(10) > 1 {
		steps = append(steps, availableStartSteps[i.rand.Intn(len(availableStartSteps))])
	}

	return steps
}

func (i SecondaryIngredient) nameWithPrefix() string {
	switch {
	case i.isMultiple:
		return fmt.Sprintf("les %s", i.name)
	case beginsWithVoyel(i.name):
		return fmt.Sprintf("l'%s", i.name)
	case i.gender == "male":
		return fmt.Sprintf("le %s", i.name)
	case i.gender == "female":
		return fmt.Sprintf("la %s", i.name)
	}
	return ""
}

func (i SecondaryIngredient) Kind() string { return "secondary-ingredient" }
func (i SecondaryIngredient) Name() string { return i.name }
func (i SecondaryIngredient) NameAndQuantity() string {
	if i.quantity == "" {
		i.prepare()
	}
	return fmt.Sprintf("%s%s", i.quantity, i.name)
}
func (i SecondaryIngredient) GetGender() string { return i.gender }
func (i SecondaryIngredient) IsMultiple() bool  { return i.isMultiple }
func (i SecondaryIngredient) TitlePart(left Ingredient) string {
	// FIXME: implement
	if left == nil {
		return i.name
	}

	part := ""

	switch left.Kind() {
	case "main-ingredient", "secondary-ingredient":
		if i.rand.Intn(10) < 5 {
			part += "et "
		}
	}
	switch {
	case i.isMultiple:
		part += "aux "
		break
	case beginsWithVoyel(i.name):
		part += "à l'"
		break
	case i.gender == "male":
		part += "au "
		break
	case i.gender == "female":
		part += "à la "
		break
	}
	part += i.name
	return part
}

func (i SecondaryIngredient) ToMap() map[string]interface{} {
	ret := make(map[string]interface{}, 0)
	ret["name"] = i.name
	ret["kind"] = i.Kind()
	ret["name-and-quantity"] = i.NameAndQuantity()
	ret["quantity"] = i.quantity
	ret["is-multiple"] = i.isMultiple
	ret["gender"] = i.gender
	ret["is-by-piece"] = i.isByPiece
	ret["is-uncountable"] = i.isUncountable
	ret["is-powder"] = i.isPowder
	ret["is-citrus"] = i.isCitrus
	ret["is-spice"] = i.isSpice
	ret["is-spreadable"] = i.isSpreadable
	return ret
}

func (i *SecondaryIngredient) SetIsByPiece() *SecondaryIngredient {
	i.isByPiece = true
	return i
}
func (i *SecondaryIngredient) SetIsSpreadable() *SecondaryIngredient {
	i.isSpreadable = true
	return i
}
func (i *SecondaryIngredient) SetIsPowder() *SecondaryIngredient {
	i.isPowder = true
	return i
}
func (i *SecondaryIngredient) SetIsUncountable() *SecondaryIngredient {
	i.isUncountable = true
	return i
}
func (i *SecondaryIngredient) SetIsSpice() *SecondaryIngredient {
	i.isSpice = true
	return i
}
func (i *SecondaryIngredient) SetIsCitrus() *SecondaryIngredient {
	i.isCitrus = true
	return i
}

//, uncountable, powder, citrus, spice, byPiece, spreadable bool) SecondaryIngredient {
