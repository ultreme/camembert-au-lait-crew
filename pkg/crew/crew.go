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
		"salutcestcool": {
			Name:        "salut c'est cool",
			Description: "<3<3<3<3",
			ImageURL:    "https://camembertaulaitcrew.github.io/assets/copaings/scc.jpg",
			Links: []Link{
				{
					Name: "Le site de les saluts",
					URL:  "http://salutcestcool.com",
				},
				{
					Name: "Facebook",
					URL:  "http://facebook.com/salutcestcool",
				},
			},
		},
		"garscool": {
			Name:        "Gars Cool",
			Description: "je suis cool",
			ImageURL:    "https://camembertaulaitcrew.github.io/assets/copaings/garscool.jpg",
			Links: []Link{
				{
					Name: "Facebook",
					URL:  "http://facebook.com/garscool",
				},
				{
					Name: "Twitter",
					URL:  "https://twitter.com/garscool",
				},
				{
					Name: "Tumblr",
					URL:  "http://garscool.tumblr.com",
				},
			},
		},
		"cestmoi": {
			Name:        "C'est moi",
			Description: "hihi",
			ImageURL:    "https://camembertaulaitcrew.github.io/assets/copaings/cestmoi.jpg",
			Links: []Link{
				{
					Name: "Facebook de C'est moi",
					URL:  "http://facebook.com/cestmoi42",
				},
			},
		},
		"pardonmyfrench": {
			Name:        "Pardon My French",
			Description: "je voudrais une croissant",
			ImageURL:    "https://camembertaulaitcrew.github.io/assets/copaings/pardonmyfrench.jpg",
			Links: []Link{
				{
					Name: "Site officiel",
					URL:  "http://www.pardon-my-french.fr",
				},
			},
		},
		"furrtek": {
			Name:        "Furrtek",
			Description: "Avant on se faisait electrocuter, maintenant on peut se faire electroniquer",
			ImageURL:    "https://camembertaulaitcrew.github.io/assets/copaings/furrtek.png",
			Links: []Link{
				{
					Name: "Site officiel",
					URL:  "http://furrtek.free.fr",
				},
				{
					Name: "Youtube",
					URL:  "https://www.youtube.com/user/furrtek/videos",
				},
			},
		},
		"tahigo": {
			Name:        "Tahigo",
			Description: "Rock festif",
			ImageURL:    "https://camembertaulaitcrew.github.io/assets/copaings/tahigo.jpg",
			Links: []Link{
				{
					Name: "Facebook",
					URL:  "http://facebook.com/tahigo",
				},
				{
					Name: "Soundcloud",
					URL:  "http://soundcloud.com/tahigo",
				},
			},
		},
		"spreadshirt": {
			Name:        "Spreadshirt",
			Description: "Des produits CALC",
			ImageURL:    "https://camembertaulaitcrew.github.io/assets/copaings/spreadshirt.jpg", // FIXME: make random
			Links: []Link{
				{
					Name: "Des T-shirts",
					URL:  "http://camembertaulaitcrew.spreadshirt.fr",
				},
			},
		},
		"estcequecestbientotlapero": {
			Name:        "Est-ce que c'est bientôt l'apéro ?",
			Description: "Outil ludique",
			ImageURL:    "https://camembertaulaitcrew.github.io/assets/copaings/apero.jpg",
			Links: []Link{
				{
					Name: "Outil",
					URL:  "http://estcequecestbientotlapero.fr",
				},
			},
		},
		"leclubdegym": {
			Name:        "Le club de Gym",
			Description: "Collectif d'artistes et de muscles",
			ImageURL:    "https://camembertaulaitcrew.github.io/assets/copaings/club.jpg",
			Links: []Link{
				{
					Name: "Facebook",
					URL:  "http://facebook.com/legymclub",
				},
			},
		},
		"leonard": {
			Name:        "Léonard Gordon Alain Souchon de la Gomme du Camembert",
			Description: "Waf.",
			ImageURL:    "https://camembertaulaitcrew.github.io/assets/copaings/leonard.jpg",
			Links: []Link{
				{
					Name: "Facebook",
					URL:  "http://facebook.com/leonard.gomme",
				},
				{
					Name: "Tumblr",
					URL:  "http://leonard-camembert.tumblr.com",
				},
			},
		},
		"jesuispasunmonstre": {
			Name:        "Je suis pas un monstre !",
			Description: "par Sassou Youpi",
			ImageURL:    "https://camembertaulaitcrew.github.io/assets/copaings/sassou-behance.jpg",
			Links: []Link{
				{
					Name: "Behance",
					URL:  "htts://www.behance.net/sassouyoupi",
				},
			},
		},
		"vadim": {
			Name:        "Vadim",
			Description: "Pipi",
			ImageURL:    "https://camembertaulaitcrew.github.io/assets/copaings/apokorounta.png",
			Links: []Link{
				{
					Name: "Apokorunta",
					URL:  "http://apokorunta.free.fr/blog",
				},
			},
		},
	}
}
