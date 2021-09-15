package io

import "regexp"

type Report struct {
	NumberOfDuplicate int
	NumberOfElements  int
	NumberOfInvalid   int
	Skus              []string
}

type ServiceReport struct {
	ReportData Report
}

func NewServiceReport() *ServiceReport {
	return &ServiceReport{Report{}}
}

func (s *ServiceReport) AddItem(input string) {
	if isAValidSku(input) {

		if !contains(s.ReportData.Skus, input) {
			s.ReportData.Skus = append(s.ReportData.Skus, input)
			s.ReportData.NumberOfElements += 1
		}

		s.ReportData.NumberOfDuplicate += 1
	} else {
		s.ReportData.NumberOfInvalid += 1
	}
}

func isAValidSku(input string) bool {
	skuRegex := regexp.MustCompile(`^[A-Z]{4}-\d{4}$`)
	return skuRegex.MatchString(input)
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
