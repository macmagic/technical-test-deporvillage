package domain

type Repository interface {
	SaveSku(sku string)
}