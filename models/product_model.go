package models

import (
	"database/sql"
	"fmt"
	"golang-database-search/entities"
)

type ProductModel struct {
	Db *sql.DB
}

func (productModel ProductModel) SearchCategory(keyword string) ([]entities.Category, error) {
	fmt.Println("keyword category", keyword)
	rows, err := productModel.Db.Query("SELECT bb_product_category_id, bb_product_category_name FROM bb_product_category WHERE LOWER(bb_product_category_name) like ? AND bb_product_category_status = 1 ORDER BY `bb_product_category_sequence` ASC LIMIT 10", "%"+keyword+"%")
	if err != nil {
		return nil, err
	} else {
		categories := []entities.Category{}
		for rows.Next() {
			var id int64
			var name string
			err2 := rows.Scan(&id, &name)
			if err2 != nil {
				return nil, err2
			} else {
				cat := entities.Category{id, name}
				categories = append(categories, cat)
			}
		}
		return categories, nil
	}
}

func (productModel ProductModel) SearchBrand(keyword string) ([]entities.Brand, error) {
	fmt.Println("keyword", keyword)
	rows, err := productModel.Db.Query("SELECT bb_product_brand_id as id, bb_product_brand_name as name, bb_product_brand_url as photo FROM bb_product_brand WHERE LOWER(bb_product_brand_name) like ? and  bb_product_brand_status = 1 ORDER BY bb_product_brand_name ASC LIMIT 10", "%"+keyword+"%")
	if err != nil {
		return nil, err
	} else {
		brands := []entities.Brand{}
		for rows.Next() {
			var id int64
			var name string
			var photo string
			err2 := rows.Scan(&id, &name, &photo)
			if err2 != nil {
				return nil, err2
			} else {
				brand := entities.Brand{id, name, photo}
				brands = append(brands, brand)
			}
		}
		return brands, nil
	}
}

// test dari awal akhir dan awal akhir
func (productModel ProductModel) StartsWith(keyword string) ([]entities.Product, error) {
	rows, err := productModel.Db.Query("select * from product where name li	ke ?", keyword+"%")
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
