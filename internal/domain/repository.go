package domain

//go:generate moq -out repository_interface_mock.go . Repository
type Repository interface {
	SaveSku(sku string)
}
