package db

import (
	"fmt"
	"medium/m/v2/internal/domain/entities"
)

var Memory map[string]*entities.Product

func Build() {
	startProducts := make(map[string]string)
	startProducts["Camisa do Grêmio"] = "clothing"
	startProducts["Capim Dourado"] = "plant"
	startProducts["CD do Atitude 67"] = "music"
	startProducts["Flash 165"] = "boat"
	startProducts["Bandana Dazaranha"] = "clothing"
	startProducts["Motul 5w40"] = "oil"

	Memory = make(map[string]*entities.Product)

	i := 0
	for product, productType := range startProducts {
		id := fmt.Sprintf("%d", i)
		Memory[id] = &entities.Product{
			ID:       id,
			Name:     product,
			Type:     productType,
			Quantity: 100,
		}
		i++
	}
}
