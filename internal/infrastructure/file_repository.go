package infrastructure

import (
	"fmt"
	"github.com/macmagic/technical-test-deporvillage/internal/application/config"
	"log"
	"os"
)

type FileRepository struct {
	logFile *os.File
}

func NewFileRepository(appConfig *config.Config) *FileRepository {
	file, err := os.OpenFile(appConfig.SkuLogPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	_, _ = file.WriteString("----SKU SUMMARY----\n")

	if err != nil {
		log.Fatalln("Cannot open the sku log file: ", err.Error())
	}

	return &FileRepository{
		logFile: file,
	}
}

func (r *FileRepository) SaveSku(sku string) {
	_, err := r.logFile.WriteString(fmt.Sprintf("Sku received: %s\n", sku))

	if err != nil {
		log.Println("Cannot write SKU in file: ", err.Error())
	}
}

func (r *FileRepository) CloseFile() {
	err := r.logFile.Close()

	if err != nil {
		log.Fatalln("Cannot close the sku log file: ", err.Error())
	}
}
