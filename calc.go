package calc

type Crew struct {
	Name    string            `json:"name"`
	Website string            `json:"url"`
	Members map[string]Person `json:"members"`
}

type Person struct {
	Name string `json:"name"`
}

var CALC = Crew{
	Name:    "Camembert au lait crew",
	Website: "http://www.camembertaulaitcrew.biz",
}

var Moul = Person{
	Name: "moul",
}

var Sassou = Person{
	Name: "Sassou",
}

var MXS = Person{
	Name: "MXS",
}

func init() {
	Calc.Members = map[string]Peron{
		"moul":   Moul,
		"sassou": Sassou,
		"mxs":    MXS,
	}
}
