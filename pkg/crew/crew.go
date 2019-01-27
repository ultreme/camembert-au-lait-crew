package crew // import "ultre.me/calcbiz/pkg/crew"

var CALC = Crew{
	Name:    "Camembert au lait crew",
	Website: "http://www.camembertaulaitcrew.biz",
	Members: []*Person{
		{
			Key:  "moul",
			Name: "moul",
		},
		{
			Key:  "sassou",
			Name: "Sassou",
		},
		{
			Key:  "mxs",
			Name: "MXS",
		},
	},
	Accounts: []*WebAccount{
		{
			Key:      "soundcloud",
			Provider: "SoundCloud",
			Handle:   "camembert-au-lait-crew",
			URL:      "https://soundcloud.com/camembert-au-lait-crew",
		},
		{
			Key:      "facebook",
			Provider: "Facebook",
			Handle:   "camembertaulaitcrew",
			URL:      "https://www.facebook.com/camembertaulaitcrew/",
		},
	},
	Friends: []*Friend{
		{
			Key:         "sbrk",
			Name:        "sbrk.org",
			Description: "gentil",
			ImageURL:    "./static/img/copaings/sbrk.jpg",
			Links: []*Link{
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
		{
			Key:         "m1ch3l",
			Name:        "m1ch3l",
			Description: "m1ch3l est cool, m1ch3l aime les gommes.",
			ImageURL:    "./static/img/copaings/m1ch3l.jpg",
			Links: []*Link{
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
		{
			Key:         "salutcestcool",
			Name:        "salut c'est cool",
			Description: "<3<3<3<3",
			ImageURL:    "./static/img/copaings/scc.jpg",
			Links: []*Link{
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
		{
			Key:         "garscool",
			Name:        "Gars Cool",
			Description: "je suis cool",
			ImageURL:    "./static/img/copaings/garscool.jpg",
			Links: []*Link{
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
		{
			Key:         "cestmoi",
			Name:        "C'est moi",
			Description: "hihi",
			ImageURL:    "./static/img/copaings/cestmoi.jpg",
			Links: []*Link{
				{
					Name: "Facebook de C'est moi",
					URL:  "http://facebook.com/cestmoi42",
				},
			},
		},
		{
			Key:         "pardonmyfrench",
			Name:        "Pardon My French",
			Description: "je voudrais une croissant",
			ImageURL:    "./static/img/copaings/pardonmyfrench.jpg",
			Links: []*Link{
				{
					Name: "Site officiel",
					URL:  "http://www.pardon-my-french.fr",
				},
			},
		},
		{
			Key:         "furrtek",
			Name:        "Furrtek",
			Description: "Avant on se faisait electrocuter, maintenant on peut se faire electroniquer",
			ImageURL:    "./static/img/copaings/furrtek.png",
			Links: []*Link{
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
		{
			Key:         "tahigo",
			Name:        "Tahigo",
			Description: "Rock festif",
			ImageURL:    "./static/img/copaings/tahigo.jpg",
			Links: []*Link{
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
		{
			Key:         "spreadshirt",
			Name:        "Spreadshirt",
			Description: "Des produits CALC",
			ImageURL:    "./static/img/img-not-found-400.png", // FIXME: make random
			Links: []*Link{
				{
					Name: "Des T-shirts",
					URL:  "http://camembertaulaitcrew.spreadshirt.fr",
				},
			},
		},
		{
			Key:         "estcequecestbientotlapero",
			Name:        "Est-ce que c'est bientôt l'apéro ?",
			Description: "Outil ludique",
			ImageURL:    "./static/img/copaings/apero.jpg",
			Links: []*Link{
				{
					Name: "Outil",
					URL:  "http://estcequecestbientotlapero.fr",
				},
			},
		},
		{
			Key:         "leclubdegym",
			Name:        "Le club de Gym",
			Description: "Collectif d'artistes et de muscles",
			ImageURL:    "./static/img/copaings/club.jpg",
			Links: []*Link{
				{
					Name: "Facebook",
					URL:  "http://facebook.com/legymclub",
				},
			},
		},
		{
			Key:         "leonard",
			Name:        "Léonard Gordon Alain Souchon de la Gomme du Camembert",
			Description: "Waf.",
			ImageURL:    "./static/img/copaings/leonard.jpg",
			Links: []*Link{
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
		{
			Key:         "jesuispasunmonstre",
			Name:        "Je suis pas un monstre !",
			Description: "par Sassou Youpi",
			ImageURL:    "./static/img/copaings/sassou-behance.jpg",
			Links: []*Link{
				{
					Name: "Behance",
					URL:  "htts://www.behance.net/sassouyoupi",
				},
			},
		},
		{
			Key:         "vadim",
			Name:        "Vadim",
			Description: "Pipi",
			ImageURL:    "./static/img/copaings/apokorunta.png",
			Links: []*Link{
				{
					Name: "Apokorunta",
					URL:  "http://apokorunta.free.fr/blog",
				},
			},
		},
	},
}
