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
	},
	{
		Title:   "Dino simple",
		CompoID: 124453182,
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
