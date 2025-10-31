package main

import "fmt"

func getData(weight, height *float64) {
	fmt.Print("Enter your weight: ")
	fmt.Scanln(weight)

	fmt.Print("Enter your height: ")
	fmt.Scanln(height)
}

func calculate_bmi(weight, height float64) (bmi float64, category string) {
	bmi = weight / height / height

	switch {
	case bmi < 19:
		category = "underweight"
	case bmi >= 19 && bmi < 25:
		category = "normal weight"
	case bmi >= 25 && bmi < 30:
		category = "pre-obesity"
	case bmi >= 30:
		category = "obesity"
	}

	return
}

func main() {
	var weight, height float64
	getData(&weight, &height)

	bmi, category := calculate_bmi(weight, height)

	fmt.Println("Your BMI:", bmi)
	fmt.Println("Your category:", category)
}
