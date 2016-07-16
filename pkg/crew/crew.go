package calccrew

type Crew struct {
	Name     string                `json:"name"`
	Website  string                `json:"url"`
	Members  map[string]Person     `json:"members"`
	Accounts map[string]WebAccount `json:"accounts"`
	Friends  map[string]Friend     `json:"friends"`
}

type Person struct {
	Name string `json:"name"`
}

type WebAccount struct {
	Provider string `json:"provider"`
	Handle   string `json:"handle"`
	URL      string `json:"url"`
}

type Link struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Friend struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	LogoURL     string `json:"logo-url"`
	Links       []Link `json:"links"`
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
	CALC.Friends = map[string]Friend{
		"sbrk": {
			Name:        "sbrk.org",
			Description: "gentil",
			Links: []Link{
				{
					Name: "Sbrk",
					URL:  "http://sbrk.org",
				},
				{
					Name: "mxs",
					URL:  "http://mxs.sbrk.org",
				},
			},
		},
	}
}
