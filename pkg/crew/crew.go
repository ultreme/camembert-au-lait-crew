package calccrew

type Crew struct {
	Name     string                `json:"name"`
	Website  string                `json:"url"`
	Members  map[string]Person     `json:"members"`
	Accounts map[string]WebAccount `json:"accounts"`
}

type Person struct {
	Name string `json:"name"`
}

type WebAccount struct {
	Provider string `json:"provider"`
	Handle   string `json:"handle"`
	URL      string `json:"url"`
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
	CALC.Members = map[string]Person{
		"moul":   Moul,
		"sassou": Sassou,
		"mxs":    MXS,
	}
	CALC.Accounts = map[string]WebAccount{
		"soundcloud": {
			Provider: "SoundCloud",
			Handle:   "camembert-au-lait-crew",
			URL:      "https://soundcloud.com/camembert-au-lait-crew",
		},
		"facebook": {
			Provider: "Facebook",
			Handle:   "camembertaulaitcrew",
			URL:      "https://www.facebook.com/camembertaulaitcrew/",
		},
	}
}
