package productdb

import (
	"fmt"
	"medium/m/v2/internal/product/productdomain/productentities"
)

var Memory map[string]*productentities.Product

func Build() {
	startProducts := make(map[string]string)
	startProducts["Camisa do Grêmio"] = "clothing"
	startProducts["Capim Dourado"] = "plant"
	startProducts["CD do Atitude 67"] = "music"
	startProducts["Flash 165"] = "boat"
	startProducts["Bandana Dazaranha"] = "clothing"
	startProducts["Motul 5w40"] = "oil"

	Memory = make(map[string]*productentities.Product)

	i := 0
	for product, productType := range startProducts {
		id := fmt.Sprintf("%d", i)
		Memory[id] = &productentities.Product{
			ID:       id,
			Name:     product,
			Type:     productType,
			Quantity: 100,
		}
		i++
	}
}
