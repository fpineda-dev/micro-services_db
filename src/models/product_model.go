package models

import (
	"database/sql"
	"fmt"

	"main.go/src/entities"
)

type ProductModel struct {
	Db *sql.DB
}

func (productModel ProductModel) FindAll() (product []entities.Product, err error) {
	rows, err := productModel.Db.Query("select * from coProducto")
	if err != nil {
		return nil, err
	} else {
		var products []entities.Product
		for rows.Next() {
			var ldProduct int32
			var NombProdu string
			var PrecioProd float32 // eslint
			var ldCategori int32
			err2 := rows.Scan(&ldProduct, &NombProdu, &PrecioProd, &ldCategori)
			if err2 != nil {
				return nil, err2
			} else {
				product := entities.Product{
					NldProduct:  ldProduct,
					CNombProdu:  NombProdu,
					NPrecioProd: PrecioProd,
					NldCategori: ldCategori,
				}
				products = append(products, product)
			}

		}

		return products, nil
	}
}

func (productModel ProductModel) Search(keyword string) (product []entities.Product, err error) {
	rows, err := productModel.Db.Query("select * from coProducto where cNombProdu like'%" + keyword + "%'")
	if err != nil {
		return nil, err
	} else {
		var products []entities.Product
		for rows.Next() {
			var ldProduct int32
			var NombProdu string
			var PrecioProd float32 // eslint
			var ldCategori int32
			err2 := rows.Scan(&ldProduct, &NombProdu, &PrecioProd, &ldCategori)
			if err2 != nil {
				return nil, err2
			} else {
				product := entities.Product{
					NldProduct:  ldProduct,
					CNombProdu:  NombProdu,
					NPrecioProd: PrecioProd,
					NldCategori: ldCategori,
				}
				products = append(products, product)
			}

		}

		return products, nil
	}
}

func (productModel ProductModel) SearchPrices(min, max float64) (product []entities.Product, err error) {
	smin := fmt.Sprintf("%f", min)
	smax := fmt.Sprintf("%f", max)
	rows, err := productModel.Db.Query("select * from coProducto where nPrecioProd between" + " " + smin + " " + "and" + " " + smax)
	if err != nil {
		return nil, err
	} else {
		var products []entities.Product
		for rows.Next() {
			var ldProduct int32
			var NombProdu string
			var PrecioProd float32 // eslint
			var ldCategori int32
			err2 := rows.Scan(&ldProduct, &NombProdu, &PrecioProd, &ldCategori)
			if err2 != nil {
				return nil, err2
			} else {
				product := entities.Product{
					NldProduct:  ldProduct,
					CNombProdu:  NombProdu,
					NPrecioProd: PrecioProd,
					NldCategori: ldCategori,
				}
				products = append(products, product)
			}

		}

		return products, nil
	}
}

func (productModel ProductModel) Create(product *entities.Product) (err error) {
	money := int64(product.NPrecioProd)

	tsql := fmt.Sprintf("insert into coProducto(cNombProdu, nPrecioProd, nIdCategori) VALUES('%s', %d, %d);",
		product.CNombProdu, money, product.NldCategori)

	result, err := productModel.Db.Exec(tsql)
	if err != nil {
		fmt.Println("Error inserting new row: " + err.Error())
		return err
	}

	ldProduct := int64(product.NldProduct)

	fmt.Println(tsql)
	fmt.Println(ldProduct)
	//fmt.Println(money)
	if err != nil {
		return err
	} else {

		ldProduct, _ = result.LastInsertId()
		return nil
	}
}

func (productModel ProductModel) Update(product *entities.Product) (int64, error) {
	money := int64(product.NPrecioProd)

	tsql := fmt.Sprintf("update coProducto set cNombProdu = '%s', nPrecioProd = %d, nIdCategori = %d where nIdProduct = %d",
		product.CNombProdu, money, product.NldCategori, product.NldProduct)

	result, err := productModel.Db.Exec(tsql)
	if err != nil {
		fmt.Println("Error updated new row: " + err.Error())
		return 0, err
	}

	ldProduct := int64(product.NldProduct)

	fmt.Println(tsql)
	fmt.Println(ldProduct)
	//fmt.Println(money)
	if err != nil {
		return 0, err
	} else {
		return result.RowsAffected()
	}
}

func (productModel ProductModel) Delete(id int64) (int64, error) {

	tsql := fmt.Sprintf("delete coProducto where nIdProduct = %d", id)

	//fmt.Println(tsql)

	result, err := productModel.Db.Exec(tsql)

	if err != nil {
		return 0, err
	} else {
		return result.RowsAffected()
	}
}
