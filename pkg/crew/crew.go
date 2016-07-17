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
	ImageURL    string `json:"image-url"`
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
			ImageURL:    "https://camembertaulaitcrew.github.io/assets/copaings/sbrk.jpg",
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
		"m1ch3l": {
			Name:        "m1ch3l",
			Description: "m1ch3l est cool, m1ch3l aime les gommes.",
			ImageURL:    "https://camembertaulaitcrew.github.io/assets/copaings/m1ch3l.jpg",
			Links: []Link{
				{
					Name: "Le site de m1ch3l",
					URL:  "http://m1ch3l.biz/",
				},
				{
					Name: "Radio m1ch3l",
					URL:  "http://radio.m1ch3l.biz",
				},
			},
		},
	}
}
