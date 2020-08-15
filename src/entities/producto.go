package entities

import (
	"fmt"
)

type Product struct {
	NldProduct  int32   `json:"id"`
	CNombProdu  string  `json:"nombre_producto"`
	NPrecioProd float32 `json:"precio_producto"`
	NldCategori int32   `json:"id_categoria"`
}

func (product Product) ToString() string {
	return fmt.Sprintf("NldProduct: %d\n CNombProdu %s\n NPrecioProd: %0.1f\n NldCategori:%d", product.NldProduct, product.CNombProdu, product.NPrecioProd, product.NldCategori)
}
