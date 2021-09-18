package domain

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_RegisterTenRandomSku(t *testing.T) {
	a := require.New(t)

	t.Run("Given a SkuService", func(t *testing.T) {
		repository := &RepositoryMock{
			SaveSkuFunc: func(sku string) {
				return
			},
		}

		service := NewSkuService(repository)
		incomingSkuList := []string{"TEST-2222", "TTT-333", "XSLK-3333", "44444", "POLP-0987", "LKJH-333345", "casa-3333", "BC-3333", "TEST-2222", "5555-ACBD"}

		t.Run(fmt.Sprintf("When we register a incoming SKU with \"%s\" value", incomingSkuList), func(t *testing.T) {
			for _, incomingSku := range incomingSkuList {
				service.RegisterSku(incomingSku)
			}

			t.Run("Then we have a 4 valid SKU, 1 duplicate SKU and 5 invalid SKU", func(t *testing.T) {
				a.Equal(4, service.report.NumberOfElements)
				a.Equal(1, service.report.NumberOfDuplicate)
				a.Equal(5, service.report.NumberOfInvalid)
				a.Equal(4, len(service.report.Skus))
			})
		})
	})
}

func Test_GenerateReport(t *testing.T) {
	a := require.New(t)

	t.Run("Given a SkuService", func(t *testing.T) {
		repository := &RepositoryMock{}

		service := NewSkuService(repository)
		service.report.NumberOfElements = 5
		service.report.NumberOfDuplicate = 5
		service.report.NumberOfInvalid = 4

		t.Run("When we want to get the summary of report", func(t *testing.T) {

			summary := service.GenerateReport()

			t.Run("Then we receive the expected summary with 5 elements, 5 duplicate and 4 invalid", func(t *testing.T) {
				a.Equal("Received 5 unique product skus, 5 duplicates, 4 discard values", summary)
			})
		})
	})
}
