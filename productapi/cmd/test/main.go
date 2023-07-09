package main

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"medium/m/v2/internal/chaos"
	"medium/m/v2/internal/product/productdomain/productentities"
)

type State string
type Action func(product *productentities.Product) (*productentities.Product, error)

const (
	MaxActions       = 8
	Init       State = "init"
	Get        State = "get"
	Search     State = "search"
	Create     State = "create"
	Update     State = "update"
	Delete     State = "delete"
)

var client = NewClient()

func main() {
	cycles := 10000

	for i := 0; i < cycles; i++ {
		fmt.Printf("Current %d\n", i)
		product := GenerateProduct()
		state := Init

		for j := 0; j < MaxActions; j++ {
			state = GetNextState(state)

			productDone, err := Execute(state, product)
			if err != nil {
				break
			}

			product = productDone
		}
	}
}

func Execute(state State, product *productentities.Product) (*productentities.Product, error) {
	action := stateAction[state]

	product, err := action(product)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return product, nil
}

func GenerateProduct() *productentities.Product {
	return &productentities.Product{
		Name:     GenerateString(10),
		Type:     GenerateType(),
		Quantity: chaos.RandomInt(5, 100),
	}
}

var types = []string{
	"clothing",
	"music",
	"games",
	"bath",
	"auto",
	"house",
	"electronics",
}

var transitions = map[State][]State{
	Init: {
		Create,
		Get,
		Search,
	},
	Search: {
		Get,
		Update,
		Delete,
	},
	Get: {
		Update,
		Delete,
	},
	Create: {
		Get,
		Update,
	},
	Update: {
		Get,
		Search,
	},
	Delete: {
		Get,
		Search,
	},
}

var stateAction = map[State]Action{
	Search: SearchProduct,
	Get:    GetProduct,
	Create: CreateProduct,
	Update: UpdateProduct,
	Delete: DeleteProduct,
}

func GetNextState(action State) State {
	length := len(transitions[action])
	id := chaos.RandomInt(0, length)
	return transitions[action][id]
}

func SearchProduct(product *productentities.Product) (*productentities.Product, error) {
	productType := GenerateType()
	fmt.Printf("Searching for %s\n", productType)
	products, err := client.Search(productType)
	if err != nil {
		return nil, err
	}

	if len(products) == 0 {
		return nil, errors.New("empty list")
	}

	return products[0], nil
}

func GetProduct(product *productentities.Product) (*productentities.Product, error) {
	fmt.Printf("Get for %s\n", product.ID)
	return client.GetByID(product.ID)
}

func CreateProduct(product *productentities.Product) (*productentities.Product, error) {
	fmt.Printf("Create %s, %s, %d\n", product.Name, product.Type, product.Quantity)
	return client.Create(product)
}

func UpdateProduct(product *productentities.Product) (*productentities.Product, error) {
	product.Name = GenerateString(10)
	product.Type = GenerateType()
	fmt.Printf("Update %s with %s, %s, %d\n", product.ID, product.Name, product.Type, product.Quantity)
	return client.Update(product)
}

func DeleteProduct(product *productentities.Product) (*productentities.Product, error) {
	fmt.Printf("Delete %s\n", product.ID)
	return product, client.Delete(product.ID)
}

func GenerateType() string {
	id := chaos.RandomInt(0, len(types))
	return types[id]
}

func GenerateString(size int) string {
	id, _ := uuid.NewUUID()
	idString := id.String()
	return idString[:size]
}
