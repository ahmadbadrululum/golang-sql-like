package main

import (
	"fmt"
	"golang-database-search/config"
	"golang-database-search/models"

	"github.com/gosimple/slug"
)

func main() {
	fmt.Println("Start with lap string")
	CallBrandMethod()

	// fmt.Println("End with top 1 string")
	// CallEndsWithMethod()

	// fmt.Println("Contains bile string")
	// CallContainsMethod()
}

func CallBrandMethod() {
	db, err := config.GetMySQLDB()
	if err != nil {
		fmt.Println(err)
	} else {
		productModel := models.ProductModel{
			Db: db,
		}
		brands, err2 := productModel.SearchBrand("voucher")
		if err2 != nil {
			fmt.Println(err2)
		} else {
			for _, brand := range brands {
				fmt.Println("Id:", brand.Id)
				fmt.Println("Name:", brand.Name)
				fmt.Println("Slug :", slug.Make(brand.Name))
				fmt.Println("Photo", brand.Photo)
				fmt.Println("----------------------------")
			}
		}
	}
}

func CallStartWithMethod() {
	db, err := config.GetMySQLDB()
	if err != nil {
		fmt.Println(err)
	} else {
		productModel := models.ProductModel{
			Db: db,
		}
		products, err2 := productModel.StartsWith("bile")
		if err2 != nil {
			fmt.Println(err2)
		} else {
			for _, product := range products {
				fmt.Println("Id:", product.Id)
				fmt.Println("Name:", product.Name)
				fmt.Println("Price:", product.Price)
				fmt.Println("Quantity:", product.Quantity)
				fmt.Println("Status:", product.Status)
				fmt.Println("----------------------------")
			}
		}
	}
}

func CallEndsWithMethod() {
	db, err := config.GetMySQLDB()
	if err != nil {
		fmt.Println(err)
	} else {
		productModel := models.ProductModel{
			Db: db,
		}
		products, err2 := productModel.EndsWith("top 1")
		if err2 != nil {
			fmt.Println(err2)
		} else {
			for _, product := range products {
				fmt.Println("Id:", product.Id)
				fmt.Println("Name:", product.Name)
				fmt.Println("Price:", product.Price)
				fmt.Println("Quantity:", product.Quantity)
				fmt.Println("Status:", product.Status)
				fmt.Println("----------------------------")
			}
		}
	}
}

func CallContainsMethod() {
	db, err := config.GetMySQLDB()
	if err != nil {
		fmt.Println(err)
	} else {
		productModel := models.ProductModel{
			Db: db,
		}
		products, err2 := productModel.Contains("bile")
		if err2 != nil {
			fmt.Println(err2)
		} else {
			for _, product := range products {
				fmt.Println("Id:", product.Id)
				fmt.Println("Name:", product.Name)
				fmt.Println("Price:", product.Price)
				fmt.Println("Quantity:", product.Quantity)
				fmt.Println("Status:", product.Status)
				fmt.Println("----------------------------")
			}
		}
	}
}
