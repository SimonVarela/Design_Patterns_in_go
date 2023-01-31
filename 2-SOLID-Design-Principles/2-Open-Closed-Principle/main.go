package main

import "fmt"

//OCP
//open for extension, closed for modification
//Specification pattern

type Color int

const (
	red Color = iota
	green
	blue
)

type Size int

const (
	small Size = iota
	medium
	large
)

type Product struct {
	name  string
	color Color
	size  Size
}

type Filter struct {
	//
}

// Breaking the OCP

func (f *Filter) FilterByColor(
	products []Product,
	color Color,
) []*Product {
	result := make([]*Product, 0)

	for i, v := range products {
		if v.color == color {
			result = append(result, &products[i])
		}
	}
	return result
}

func (f *Filter) FilterBySize(
	products []Product,
	size Size,
) []*Product {
	result := make([]*Product, 0)

	for i, v := range products {
		if v.size == size {
			result = append(result, &products[i])
		}
	}
	return result
}

func (f *Filter) FilterBySizeAndColor(
	products []Product,
	color Color,
	size Size,
) []*Product {
	result := make([]*Product, 0)

	for i, v := range products {
		if v.color == color && v.size == size {
			result = append(result, &products[i])
		}
	}
	return result
}

// *******************************

type Specification interface {
	IsSatisfied(p *Product) bool
}

type ColorSpecification struct {
	color Color
}

func (c ColorSpecification) IsSatisfied(p *Product) bool {
	return p.color == c.color
}

type SizeSpecification struct {
	size Size
}

func (s SizeSpecification) IsSatisfied(p *Product) bool {
	return p.size == s.size
}

type AndSpecification struct {
	firts, second Specification
}

func (a AndSpecification) IsSatisfied(p *Product) bool {
	return a.firts.IsSatisfied(p) && a.second.IsSatisfied(p)
}

type BetterFilter struct{}

func (f *BetterFilter) Filter(
	products []Product,
	spec Specification,
) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if spec.IsSatisfied(&v) {
			result = append(result, &products[i])
		}
	}

	return result
}

func main() {
	apple := Product{"Apple", green, small}
	tree := Product{"Tree", green, large}
	house := Product{"House", blue, large}

	products := []Product{apple, tree, house}
	fmt.Printf("Green products (old): \n")

	f := Filter{}
	for _, v := range f.FilterByColor(products, green) {
		fmt.Printf(" - %s is green\n", v.name)
	}

	fmt.Printf("Green products (new): \n")
	greenSpec := ColorSpecification{green}
	bf := BetterFilter{}
	for _, v := range bf.Filter(products, greenSpec) {
		fmt.Printf(" - %s is green\n", v.name)
	}

	largeSpec := SizeSpecification{large}
	lgSpec := AndSpecification{greenSpec, largeSpec}
	fmt.Printf("Large and green products: \n")
	for _, v := range bf.Filter(products, lgSpec) {
		fmt.Printf(" - %s is green and large\n", v.name)
	}
}
