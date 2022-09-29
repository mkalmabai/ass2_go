package main

import "fmt"

type Shop interface {
	getPrice() int
}
type product struct {
}

func (i *product) getPrice() int {
	return 400
}

type productTp struct {
	prod Shop
}

func (c *productTp) getPrice() int {
	prodprice := c.prod.getPrice()
	return prodprice + 100
}

type walptop struct {
	prod Shop
}

func (w *walptop) getPrice() int {
	prodprice := w.prod.getPrice()
	return prodprice + 500
}

func main() {
	prod := &product{}

	prodwithwal := &productTp{
		prod: prod,
	}

	prodwithpr := &walptop{
		prod: prodwithwal,
	}
	fmt.Println("Price of ice prod with chocolate and walnut topping is: ", prodwithpr.getPrice(), "Tenge")
}
