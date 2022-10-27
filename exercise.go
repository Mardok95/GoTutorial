// This is a file that i use for exercise about random tasks.

package main

import (
	"fmt"
)

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	switch {
	case balance < 0.0:
		return float32(3.213)
	case balance >= 0.0 && balance < 1000.0:
		return float32(0.5)
	case balance >= 1000.0 && balance < 5000.0:
		return float32(1.621)
	default:
		return float32(2.475)
	}
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	interest := InterestRate(balance)
	interest = float32(balance) * interest / 100
	return float64(interest)
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	interest := Interest(balance)
	result := interest + balance
	return result
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance.
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	years := 0
	for balance <= targetBalance {
		balance = balance + Interest(balance)
		years = years + 1
	}
	return years
}

func main() {
	fmt.Println(InterestRate(1000.0))
	fmt.Println(Interest(200.75))
	fmt.Println(AnnualBalanceUpdate(200.75))

	balance := 200.75
	targetBalance := 214.88
	fmt.Println(YearsBeforeDesiredBalance(balance, targetBalance))

}
