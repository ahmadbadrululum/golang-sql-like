package models

import (
	"database/sql"
	"golang-database-search/entities"
)

type ProductModel struct {
	Db *sql.DB
}

func (productModel ProductModel) StartsWith(keyword string) ([]entities.Product, error) {
	rows, err := productModel.Db.Query("select * from product where name like ?", keyword+"%")
	if err != nil {
		return nil, err
	} else {
		products := []entities.Product{}
		for rows.Next() {
			var id int64
			var name string
			var price float32
			var quantity int
			var status bool
			err2 := rows.Scan(&id, &name, &price, &quantity, &status)
			if err2 != nil {
				return nil, err2
			} else {
				product := entities.Product{id, name, price, quantity, status}
				products = append(products, product)
			}
		}
		return products, nil
	}
}

func (productModel ProductModel) EndsWith(keyword string) ([]entities.Product, error) {
	rows, err := productModel.Db.Query("select * from product where name like ?", "%"+keyword)
	if err != nil {
		return nil, err
	} else {
		products := []entities.Product{}
		for rows.Next() {
			var id int64
			var name string
			var price float32
			var quantity int
			var status bool
			err2 := rows.Scan(&id, &name, &price, &quantity, &status)
			if err2 != nil {
				return nil, err2
			} else {
				product := entities.Product{id, name, price, quantity, status}
				products = append(products, product)
			}
		}
		return products, nil
	}
}

func (productModel ProductModel) Contains(keyword string) ([]entities.Product, error) {
	rows, err := productModel.Db.Query("select * from product where name like ?", "%"+keyword+"%")
	if err != nil {
		return nil, err
	} else {
		products := []entities.Product{}
		for rows.Next() {
			var id int64
			var name string
			var price float32
			var quantity int
			var status bool
			err2 := rows.Scan(&id, &name, &price, &quantity, &status)
			if err2 != nil {
				return nil, err2
			} else {
				product := entities.Product{id, name, price, quantity, status}
				products = append(products, product)
			}
		}
		return products, nil
	}
}
