package ingredients

import "math/rand"

type IngredientMethod struct {
	singleMale     string
	multipleMale   string
	singleFemale   string
	multipleFemale string
	steps          Steps
	rand           *rand.Rand
	left           Ingredient
}

func (i *IngredientMethod) SetMethod(method Method) { panic("Not implemented") }
func (i *IngredientMethod) GetMethod() Method       { return nil }

func NewIngredientMethod(singleMale, singleFemale, multipleMale, multipleFemale string, steps Steps, rnd *rand.Rand) *IngredientMethod {
	return &IngredientMethod{
		singleMale:     singleMale,
		multipleMale:   multipleMale,
		singleFemale:   singleFemale,
		multipleFemale: multipleFemale,
		steps:          steps,
		rand:           rnd,
	}
}

func (i IngredientMethod) SetLeft(left Ingredient) { i.left = left }
func (i IngredientMethod) GetSteps() Steps         { return i.steps }
func (i IngredientMethod) IsMultiple() bool {
	if i.left != nil {
		return i.left.IsMultiple()
	}
	panic("should not happen")
}
func (i IngredientMethod) GetGender() string {
	if i.left != nil {
		return i.left.GetGender()
	}
	panic("should not happen")
}

func (i IngredientMethod) TitlePart(left Ingredient) string {
	if left == nil {
		return i.singleMale
	}
	gender := left.GetGender()
	isMultiple := left.IsMultiple()
	switch {
	case gender == "male" && !isMultiple:
		return i.singleMale
	case gender == "male" && isMultiple:
		return i.multipleMale
	case gender == "female" && !isMultiple:
		return i.singleFemale
	case gender == "female" && isMultiple:
		return i.multipleFemale
	}
	panic("should not happen")
}

func (i IngredientMethod) Kind() string            { return "main-ingredient-method" }
func (i IngredientMethod) Name() string            { return i.TitlePart(i.left) }
func (i IngredientMethod) NameAndQuantity() string { return i.Name() }

func (i IngredientMethod) ToMap() map[string]interface{} {
	ret := make(map[string]interface{}, 0)
	ret["name"] = i.Name()
	ret["kind"] = i.Kind()
	return ret
}
