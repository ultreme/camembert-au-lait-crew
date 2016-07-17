package calcspreadshirt

import (
	"fmt"
	"math/rand"
)

type Product struct {
	Title    string `json:"title"`
	CompoID  uint64 `json:"compo-id"`
	Type     string `json:"type"`
	URL      string `json:"url"`
	ImageURL string `json:"image-url"`
}

var products = []Product{
	{
		Title:   "Cours de sport",
		CompoID: 124452996,
	}, {
		Title:   "Dino simple",
		CompoID: 124453182,
	}, {
		Title:   "Dino rainbow",
		CompoID: 124452969,
	}, {
		Title:   "Calc degrade",
		CompoID: 124452956,
	}, {
		Title:   "Calc pixel",
		CompoID: 124452945,
	}, {
		Title:   "J'veux faire des trucs",
		CompoID: 124452911,
	}, {
		Title:   "Gars cool approved",
		CompoID: 124452840,
	}, {
		Title:   "Gars cool fait du rap",
		CompoID: 124403201,
	}, {
		Title:   "Un bon plan",
		CompoID: 124403188,
	}, {
		Title:   "Calc white-on-black",
		CompoID: 124399349,
	}, {
		Title:   "Trop de la balle",
		CompoID: 124399356,
	}, {
		Title:   "DJ Jean-Luc Puissant",
		CompoID: 124399370,
	}, {
		Title:   "C'est la fiesta, wouhouuu",
		CompoID: 124399385,
	}, {
		Title:   "A la campagne, il y a un banc",
		CompoID: 124399407,
	}, {
		Title:   "50% bacon, 50% biere, 50% camembert",
		CompoID: 124399002,
	}, {
		Title:   "Des trucs",
		CompoID: 124399027,
	}, {
		Title:   "Pochette camembert",
		CompoID: 124399084,
	}, {
		Title:   "Une gomme",
		CompoID: 124399103,
	}, {
		Title:   "Une vraie gomme",
		CompoID: 124399138,
	}, {
		Title:   "TPYO",
		CompoID: 133595868,
	}, {
		Title:   "Millipede",
		CompoID: 133595825,
	}, {
		Title:   "Sapin",
		CompoID: 133595722,
	}, {
		Title:   ":ok_hand:",
		CompoID: 133595728,
	},
}

func init() {
	for idx, product := range products {
		if product.Type == "" {
			products[idx].Type = "t-shirt"
		}
		products[idx].URL = "http://camembertaulaitcrew.spreadshirt.fr/"
	}
}

func (p *Product) getRandomImage(width, height uint64) string {
	appearanceID := rand.Intn(4) + 1
	baseURL := "http://image.spreadshirt.net/image-server/v1/products/%d/views/1,width=%d,height=%d,appearanceId=%d,rotateX=0,rotateY=-20,rotateZ=0.png"
	return fmt.Sprintf(baseURL, p.CompoID, width, height, appearanceID)
}

func GetAllProducts(width, height uint64) []Product {
	// FIXME: shuffle
	for idx := range products {
		products[idx].prepare(width, height)
	}
	return products
}

func (p *Product) prepare(width, height uint64) {
	p.ImageURL = p.getRandomImage(width, height)
}

func GetRandomProduct(width, height uint64) Product {
	product := products[rand.Intn(len(products))]
	product.prepare(width, height)
	return product
}

func getProductDetail(compoID, width, height uint64) (*Product, error) {
	for _, product := range products {
		if product.CompoID == compoID {
			product.prepare(width, height)
			return &product, nil
		}
	}
	return nil, fmt.Errorf("No such compo %d", compoID)
}
