package infrastructure

import "log"

type FileRepository struct {

}

func NewFileRepository() *FileRepository {
	return &FileRepository{}
}

func (r *FileRepository) SaveSku(sku string) {
	log.Println(sku)
}
