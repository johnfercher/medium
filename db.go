package main

import (
	"fmt"
)

var memoryDb map[string]*Product

func BuildDb() {
	startProducts := make(map[string]string)
	startProducts["Camisa do Grêmio"] = "clothing"
	startProducts["Capim Dourado"] = "plant"
	startProducts["CD do Atitude 67"] = "music"
	startProducts["Flash 165"] = "boat"
	startProducts["Bandana Dazaranha"] = "clothing"
	startProducts["Motul 5w40"] = "oil"

	memoryDb = make(map[string]*Product)

	i := 0
	for product, productType := range startProducts {
		id := fmt.Sprintf("%d", i)
		memoryDb[id] = &Product{
			ID:       id,
			Name:     product,
			Type:     productType,
			Quantity: 100,
		}
		i++
	}
}
