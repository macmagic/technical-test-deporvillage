package domain

import "regexp"

type SkuServiceInterface interface {
	RegisterSku(incomingSku string)
}

type Report struct {
	NumberOfDuplicate int
	NumberOfElements  int
	NumberOfInvalid   int
	Skus              []string
}

type SkuService struct {
	repository Repository
	report Report
}

func NewSkuService(repository Repository) *SkuService {
	return &SkuService{
		repository: repository,
		report: Report{},
	}
}

func (s *SkuService) RegisterSku(incomingSku string)  {
	if isAValidSku(incomingSku) {
		if s.contains(incomingSku) {
			s.report.NumberOfDuplicate += 1
			return
		}

		s.report.Skus = append(s.report.Skus, incomingSku)
		s.report.NumberOfElements += 1
	} else {
		s.report.NumberOfInvalid += 1
	}
}

func isAValidSku(input string) bool {
	skuRegex := regexp.MustCompile(`^[A-Z]{4}-\d{4}$`)
	return skuRegex.MatchString(input)
}

func (s *SkuService) contains(incomingSku string) bool {
	for _, item := range s.report.Skus {
		if item == incomingSku {
			return true
		}
	}
	return false
}

