package repository

import (
	"fmt"
	"os"
	"tidy/src/pb/products"

	"google.golang.org/protobuf/proto"
)

type ProductRepository struct{}

const fileName string = "products.txt"

func (p *ProductRepository) loadData() (products.ProductList, error) {
	productList := products.ProductList{}
	data, err := os.ReadFile(fileName)
	if err != nil {
		return productList, fmt.Errorf("Erro ao load data:", err.Error())
	}
	if err := proto.Unmarshal(data, &productList); err != nil {
		return productList, fmt.Errorf("Erro ao deseralizar:", err.Error())
	}
	return productList, nil
}
func (p *ProductRepository) saveData(productlist products.ProductList) error {
	data, err := proto.Marshal(&productlist)
	if err != nil {
		return err
	}
	errW := os.WriteFile(fileName, data, 0644)
	if errW != nil {
		return errW
	}
	return nil
}

func (p *ProductRepository) Create(product products.Product) (products.Product, error) {
	prodlist, err := p.loadData()
	if err != nil {
		return products.Product{}, err
	}
	product.Id = int32(len(prodlist.Products) + 1)
	prodlist.Products = append(prodlist.Products, &product)
	err = p.saveData(prodlist)
	if err != nil {
		return products.Product{}, err
	}
	return product, nil
}

func (p *ProductRepository) FindAll() (products.ProductList, error) {
	return p.loadData()
}
