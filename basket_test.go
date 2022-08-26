package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/cucumber/godog"
)

type shopping struct{
	shelf *Shelf
	basket *Basket
}

func (sh *shopping) addProduct(productName string, price float64) (err error) {
	return sh.shelf.AddProduct(productName, price)
  }

func (sh *shopping) iAddTheToTheBasket(productName string) error {
	return sh.basket.AddProduct(productName, 1)
}

func (sh *shopping) iShouldHaveProductsInTheBasket(count int) error {
	bLen := len(sh.basket.products)
	if bLen != count {
		return errors.New(fmt.Sprintf("wrong number of items in basket, needed %d but got %d", count, bLen))
	}

	return nil
}

func (sh *shopping) theOverallBasketPriceShouldBe(price float64) error {
	var actualPrice float64 = 0.0

	for name, count := range sh.basket.products {
		actualPrice = actualPrice + sh.shelf.products[name] * float64(count)
	}

	// VAT
	actualPrice = actualPrice * 1.20

	// Delivery
	if actualPrice > 10.0 {
		actualPrice = actualPrice + 2
	} else {
		actualPrice = actualPrice + 3
	}

	if actualPrice != price {
		return errors.New(fmt.Sprintf("wrong price for current basket, needed %f but got %f", price, actualPrice))
	}

	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	sh := &shopping{}

	ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		sh.shelf = NewShelf()
		sh.basket = NewBasket()

		return ctx, nil
	})

	ctx.Step(`^I add the "([^"]*)" to the basket$`, sh.iAddTheToTheBasket)
	ctx.Step(`^I should have (\d+) products in the basket$`, sh.iShouldHaveProductsInTheBasket)
	ctx.Step(`^the overall basket price should be £(\d+)$`, sh.theOverallBasketPriceShouldBe)
	ctx.Step(`^there is a "([a-zA-Z\s]+)", which costs £(\d+)$`, sh.addProduct)
}
