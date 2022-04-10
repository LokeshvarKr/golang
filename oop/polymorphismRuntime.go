package main

import "fmt"

type Employee interface {
	CalculateSalary() int
}

type JobGrade1 struct {
	Basic int
	HRA   int
}

type JobGrade2 struct {
	Basic      int
	HRA        int
	Travelling int
	Others     int
}

func (j JobGrade1) CalculateSalary() int {
	return j.Basic + j.HRA + 1000
}

func (j JobGrade2) CalculateSalary() int {
	return j.Basic + j.HRA + j.Travelling + j.Others
}

func CalculateSalaryForIndustry(employees []Employee) {
	total := 0
	for _, e := range employees {
		total += e.CalculateSalary()
	}

	fmt.Println("total Salary Expenses : ", total)
}

func main() {

	employees := []Employee{
		JobGrade1{
			Basic: 10000,
			HRA:   5000,
		},
		JobGrade2{
			Basic:      12000,
			HRA:        5500,
			Travelling: 2500,
			Others:     500,
		},
	}

	CalculateSalaryForIndustry(employees)
}
