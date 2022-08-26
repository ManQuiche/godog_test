package main

func NewBasket() *Basket {
	return &Basket{
		products: make(map[string]int),
	}
}

type Basket struct {
	products map[string]int
}

func (b *Basket) AddProduct(productName string, quantity int) error {
	b.products[productName] = b.products[productName] + quantity

	return nil
}